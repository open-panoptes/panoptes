package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/prometheus/prompb"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	"github.com/rancher/opni/pkg/clients"
	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/task"
	"github.com/rancher/opni/plugins/metrics/pkg/apis/remoteread"
	"github.com/rancher/opni/plugins/metrics/pkg/apis/remotewrite"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"sync"
	"time"
)

// todo: this should probably be more sophisticated than this to handle read size limits
var TimeDeltaMillis = time.Minute.Milliseconds()

// todo: import prometheus LabelMatcher into plugins/metrics/pkg/apis/remoteread.proto to remove this
func toLabelMatchers(rrLabelMatchers []*remoteread.LabelMatcher) []*prompb.LabelMatcher {
	pbLabelMatchers := make([]*prompb.LabelMatcher, 0, len(rrLabelMatchers))

	for _, matcher := range rrLabelMatchers {
		var matchType prompb.LabelMatcher_Type

		switch matcher.Type {
		case remoteread.LabelMatcher_Equal:
			matchType = prompb.LabelMatcher_EQ
		case remoteread.LabelMatcher_NotEqual:
			matchType = prompb.LabelMatcher_NEQ
		case remoteread.LabelMatcher_RegexEqual:
			matchType = prompb.LabelMatcher_RE
		case remoteread.LabelMatcher_NotRegexEqual:
			matchType = prompb.LabelMatcher_NRE
		default:
			// todo: log something
		}

		pbLabelMatchers = append(pbLabelMatchers, &prompb.LabelMatcher{
			Type:  matchType,
			Name:  matcher.Name,
			Value: matcher.Value,
		})
	}

	return pbLabelMatchers
}

func dereferenceResultTimeseries(in []*prompb.TimeSeries) []prompb.TimeSeries {
	dereferenced := make([]prompb.TimeSeries, 0, len(in))

	for _, ref := range in {
		dereferenced = append(dereferenced, *ref)
	}

	return dereferenced
}

func getMessageFromTaskLogs(logs []*corev1.LogEntry) string {
	for i := len(logs) - 1; i >= 0; i-- {
		log := logs[i]
		if log.Level == int32(zapcore.ErrorLevel) {
			return log.Msg
		}
	}

	return ""
}

type TargetRunMetadata struct {
	Target *remoteread.Target
	Query  *remoteread.Query
}

type targetStore struct {
	innerMu sync.RWMutex
	inner   map[string]*corev1.TaskStatus
}

func (store *targetStore) Put(ctx context.Context, key string, value *corev1.TaskStatus) error {
	store.innerMu.Lock()
	defer store.innerMu.Unlock()
	store.inner[key] = value
	return nil
}

func (store *targetStore) Get(ctx context.Context, key string) (*corev1.TaskStatus, error) {
	store.innerMu.RLock()
	defer store.innerMu.RUnlock()

	status, found := store.inner[key]
	if !found {
		return nil, storage.ErrNotFound
	}

	return status, nil
}

func (store *targetStore) Delete(ctx context.Context, key string) error {
	store.innerMu.Lock()
	defer store.innerMu.Unlock()
	delete(store.inner, key)
	return nil
}

func (store *targetStore) ListKeys(ctx context.Context, prefix string) ([]string, error) {
	store.innerMu.RLock()
	defer store.innerMu.RUnlock()
	keys := make([]string, 0, len(store.inner))
	for key := range store.inner {
		if strings.HasPrefix(key, prefix) {
			keys = append(keys, key)
		}
	}
	return keys, nil
}

type taskRunner struct {
	remoteWriteClient clients.Locker[remotewrite.RemoteWriteClient]

	remoteReaderMu sync.RWMutex
	remoteReader   RemoteReader
}

func newTaskRunner() *taskRunner {
	return &taskRunner{
		//remoteReader: NewRemoteReader(&http.Client{}),
	}
}

func (runner *taskRunner) SetRemoteWriteClient(client clients.Locker[remotewrite.RemoteWriteClient]) {
	runner.remoteWriteClient = client
}

func (runner *taskRunner) SetRemoteReaderClient(client RemoteReader) {
	runner.remoteReaderMu.Lock()
	defer runner.remoteReaderMu.Unlock()

	runner.remoteReader = client
}

func (runner *taskRunner) OnTaskPending(ctx context.Context, activeTask task.ActiveTask) error {
	return nil
}

func (runner *taskRunner) OnTaskRunning(ctx context.Context, activeTask task.ActiveTask) error {
	run := &TargetRunMetadata{}
	activeTask.LoadTaskMetadata(run)

	labelMatchers := toLabelMatchers(run.Query.Matchers)

	importEnd := run.Query.EndTimestamp.AsTime().UnixMilli()
	nextStart := run.Query.StartTimestamp.AsTime().UnixMilli()
	nextEnd := nextStart

	progress := &corev1.Progress{
		Current: 0,
		Total:   uint64(importEnd - nextStart),
	}

	activeTask.SetProgress(progress)

	for nextStart < importEnd {
		nextStart = nextEnd
		nextEnd = nextStart + TimeDeltaMillis

		if nextStart >= importEnd {
			break
		}

		if nextEnd >= importEnd {
			nextEnd = importEnd
		}

		readRequest := &prompb.ReadRequest{
			Queries: []*prompb.Query{
				{
					StartTimestampMs: nextStart,
					EndTimestampMs:   nextEnd,
					Matchers:         labelMatchers,
				},
			},
		}

		runner.remoteReaderMu.Lock()
		readResponse, err := runner.remoteReader.Read(context.Background(), run.Target.Spec.Endpoint, readRequest)
		runner.remoteReaderMu.Unlock()

		if err != nil {
			activeTask.AddLogEntry(zapcore.ErrorLevel, fmt.Sprintf("failed to read from target endpoint: %s", err))
			return fmt.Errorf("failed to read from target endpoint: %w", err)
		}

		for _, result := range readResponse.Results {
			if len(result.Timeseries) == 0 {
				continue
			}

			writeRequest := prompb.WriteRequest{
				Timeseries: dereferenceResultTimeseries(result.Timeseries),
			}

			uncompressed, err := proto.Marshal(&writeRequest)
			if err != nil {
				activeTask.AddLogEntry(zapcore.ErrorLevel, fmt.Sprintf("failed to uncompress data from target endpoint: %s", err))
				return fmt.Errorf("failed to uncompress data from target endpoint: %w", err)
			}

			compressed := snappy.Encode(nil, uncompressed)

			payload := &remotewrite.Payload{
				Contents: compressed,
			}

			runner.remoteWriteClient.Use(func(remoteWriteClient remotewrite.RemoteWriteClient) {
				if _, err := remoteWriteClient.Push(context.Background(), payload); err != nil {
					activeTask.AddLogEntry(zapcore.ErrorLevel, fmt.Sprintf("failed to push to remote write: %s", err))
					return
				}

				activeTask.AddLogEntry(zapcore.DebugLevel, fmt.Sprintf("pushed %d bytes to remote write", len(payload.Contents)))
			})

			progress.Current = uint64(nextEnd)
			activeTask.SetProgress(progress)
		}
	}

	return nil
}

func (runner *taskRunner) OnTaskCompleted(ctx context.Context, activeTask task.ActiveTask, state task.State, args ...any) {
}

type TargetRunner interface {
	Start(target *remoteread.Target, query *remoteread.Query) error

	Stop(name string) error

	GetStatus(name string) (*remoteread.TargetStatus, error)

	SetRemoteWriteClient(client clients.Locker[remotewrite.RemoteWriteClient])

	SetRemoteReaderClient(client RemoteReader)
}

type taskingTargetRunner struct {
	logger *zap.SugaredLogger

	runnerMu sync.RWMutex
	runner   *taskRunner

	controller *task.Controller

	remoteWriteClient clients.Locker[remotewrite.RemoteWriteClient]
}

func NewTargetRunner(logger *zap.SugaredLogger) TargetRunner {
	store := &targetStore{
		inner: make(map[string]*corev1.TaskStatus),
	}

	runner := newTaskRunner()

	controller, err := task.NewController(context.Background(), "Target-runner", store, runner)
	if err != nil {
		panic(fmt.Sprintf("bug: failed to create Target task corntoller: %s", err))
	}

	return &taskingTargetRunner{
		logger:     logger,
		runner:     runner,
		controller: controller,
	}
}

func (runner *taskingTargetRunner) Start(target *remoteread.Target, query *remoteread.Query) error {
	if status, err := runner.controller.TaskStatus(target.Meta.Name); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return fmt.Errorf("error checking for Target status: %s", err)
		}
	} else {
		switch status.State {
		case corev1.TaskState_Running, corev1.TaskState_Pending:
			return fmt.Errorf("target is already running")
		case corev1.TaskState_Unknown:
			return fmt.Errorf("target state is unknown")
		}
	}

	if err := runner.controller.LaunchTask(target.Meta.Name, task.WithMetadata(TargetRunMetadata{
		Target: target,
		Query:  query,
	})); err != nil {
		return fmt.Errorf("could not run Target: %w", err)
	}

	runner.logger.Infof("started Target '%s'", target.Meta.Name)

	return nil
}

func (runner *taskingTargetRunner) Stop(name string) error {
	if status, err := runner.controller.TaskStatus(name); err != nil {
		return fmt.Errorf("Target not found")
	} else {
		switch status.State {
		case corev1.TaskState_Canceled, corev1.TaskState_Completed, corev1.TaskState_Failed:
			return fmt.Errorf("Target is not running")
		}
	}

	runner.controller.CancelTask(name)

	runner.logger.Infof("stopped Target '%s'", name)

	return nil
}

func (runner *taskingTargetRunner) GetStatus(name string) (*remoteread.TargetStatus, error) {
	taskStatus, err := runner.controller.TaskStatus(name)
	if err != nil {
		return nil, fmt.Errorf("could not get Target status: %w", err)
	}

	taskMetadata := TargetRunMetadata{}
	if err := json.Unmarshal([]byte(taskStatus.Metadata), &taskMetadata); err != nil {
		return nil, fmt.Errorf("could not parse Target metedata: %w", err)
	}

	statusProgress := &remoteread.TargetProgress{
		StartTimestamp:    taskMetadata.Query.StartTimestamp,
		LastReadTimestamp: nil,
		EndTimestamp:      taskMetadata.Query.EndTimestamp,
	}

	if taskStatus.Progress == nil {
		statusProgress.LastReadTimestamp = statusProgress.StartTimestamp
	}

	return &remoteread.TargetStatus{
		Progress: statusProgress,
		Message:  getMessageFromTaskLogs(taskStatus.Logs),
		State:    taskStatus.State,
	}, nil
}

func (runner *taskingTargetRunner) SetRemoteWriteClient(client clients.Locker[remotewrite.RemoteWriteClient]) {
	runner.runner.SetRemoteWriteClient(client)
}

func (runner *taskingTargetRunner) SetRemoteReaderClient(client RemoteReader) {
	runner.runnerMu.Lock()
	defer runner.runnerMu.Unlock()

	runner.runner.SetRemoteReaderClient(client)
}
