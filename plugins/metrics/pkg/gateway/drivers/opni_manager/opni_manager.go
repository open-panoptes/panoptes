package drivers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cisco-open/k8s-objectmatcher/patch"
	opnicorev1beta1 "github.com/rancher/opni/apis/core/v1beta1"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	"github.com/rancher/opni/pkg/plugins/driverutil"
	"github.com/rancher/opni/pkg/util/k8sutil"
	"github.com/rancher/opni/plugins/metrics/apis/cortexops"
	"github.com/rancher/opni/plugins/metrics/pkg/gateway/drivers"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type OpniManager struct {
	cortexops.UnsafeCortexOpsServer
	OpniManagerClusterDriverOptions
}

type OpniManagerClusterDriverOptions struct {
	K8sClient         client.Client        `option:"k8sClient"`
	MonitoringCluster types.NamespacedName `option:"monitoringCluster"`
	GatewayRef        types.NamespacedName `option:"gatewayRef"`
}

func NewOpniManagerClusterDriver(options OpniManagerClusterDriverOptions) (*OpniManager, error) {
	if options.K8sClient == nil {
		s := scheme.Scheme
		opnicorev1beta1.AddToScheme(s)
		c, err := k8sutil.NewK8sClient(k8sutil.ClientOptions{
			Scheme: s,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create kubernetes client: %w", err)
		}
		options.K8sClient = c
	}
	return &OpniManager{
		OpniManagerClusterDriverOptions: options,
	}, nil
}

func (k *OpniManager) newMonitoringCluster() *opnicorev1beta1.MonitoringCluster {
	return &opnicorev1beta1.MonitoringCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k.MonitoringCluster.Name,
			Namespace: k.MonitoringCluster.Namespace,
		},
	}
}

func (k *OpniManager) GetClusterConfiguration(ctx context.Context, _ *emptypb.Empty) (*cortexops.ClusterConfiguration, error) {
	mc := k.newMonitoringCluster()
	err := k.K8sClient.Get(ctx, client.ObjectKeyFromObject(mc), mc)
	if err != nil {
		return nil, err
	}
	storage := mc.Spec.Cortex.Storage.DeepCopy()
	storage.RedactSecrets()
	apiWorkloads := &cortexops.CortexWorkloadsSpec{}
	mergeK8sWorkloadsToApi(apiWorkloads, &mc.Spec.Cortex.Workloads)
	return &cortexops.ClusterConfiguration{
		Mode: cortexops.DeploymentMode(cortexops.DeploymentMode_value[string(mc.Spec.Cortex.DeploymentMode)]),
		Grafana: &cortexops.GrafanaConfig{
			Enabled:  mc.Spec.Grafana.Enabled,
			Hostname: mc.Spec.Grafana.Hostname,
		},
		Workloads: apiWorkloads,
		Cortex: &cortexops.CortexConfig{
			Storage:       storage,
			Limits:        mc.Spec.Cortex.Limits,
			RuntimeConfig: mc.Spec.Cortex.RuntimeConfig,
			Compactor:     mc.Spec.Cortex.CompactorConfig,
			Querier:       mc.Spec.Cortex.QuerierConfig,
			LogLevel:      mc.Spec.Cortex.LogLevel,
		},
	}, nil
}

func (k *OpniManager) ConfigureCluster(ctx context.Context, conf *cortexops.ClusterConfiguration) (*emptypb.Empty, error) {
	cluster := k.newMonitoringCluster()

	objectKey := client.ObjectKeyFromObject(cluster)
	err := k.K8sClient.Get(ctx, objectKey, cluster)
	exists := true
	if err != nil {
		if k8serrors.IsNotFound(err) {
			exists = false
		} else {
			return nil, fmt.Errorf("failed to get monitoring cluster: %w", err)
		}
	}

	// look up the gateway so we can set it as an owner reference
	gateway := &opnicorev1beta1.Gateway{}
	err = k.K8sClient.Get(ctx, k.GatewayRef, gateway)
	if err != nil {
		return nil, fmt.Errorf("failed to get gateway: %w", err)
	}
	defaultGrafanaHostname := "grafana." + gateway.Spec.Hostname

	if conf.Grafana == nil {
		conf.Grafana = &cortexops.GrafanaConfig{
			Enabled: true,
		}
	}
	if conf.Grafana.Hostname == "" {
		conf.Grafana.Hostname = defaultGrafanaHostname
	}

	mutator := func(cluster *opnicorev1beta1.MonitoringCluster) error {
		if err := conf.GetCortex().GetStorage().UnredactSecrets(cluster.Spec.Cortex.Storage); err != nil {
			return err
		}
		cluster.Spec.Cortex.Enabled = true
		cluster.Spec.Cortex.Storage = conf.GetCortex().GetStorage()
		if cluster.Spec.Cortex.Storage.Filesystem != nil &&
			cluster.Spec.Cortex.Storage.Filesystem.Directory == "" {
			cluster.Spec.Cortex.Storage.Filesystem.Directory = "/data"
		}
		cluster.Spec.Grafana.Enabled = conf.Grafana.Enabled
		cluster.Spec.Grafana.Hostname = conf.Grafana.Hostname
		cluster.Spec.Gateway = v1.LocalObjectReference{
			Name: k.GatewayRef.Name,
		}
		cluster.Spec.Cortex.DeploymentMode = opnicorev1beta1.DeploymentMode(cortexops.DeploymentMode_name[int32(conf.GetMode())])
		if conf.GetWorkloads() != nil {
			if err := mergeApiWorkloadsToK8s(&cluster.Spec.Cortex.Workloads, conf.GetWorkloads()); err != nil {
				return err
			}
		}
		if conf.GetCortex().GetLimits() != nil {
			cluster.Spec.Cortex.Limits = conf.Cortex.Limits
		}
		if conf.GetCortex().GetCompactor() != nil {
			cluster.Spec.Cortex.CompactorConfig = conf.Cortex.Compactor
		}
		if conf.GetCortex().GetQuerier() != nil {
			cluster.Spec.Cortex.QuerierConfig = conf.Cortex.Querier
		}
		if conf.GetCortex().GetRuntimeConfig() != nil {
			cluster.Spec.Cortex.RuntimeConfig = conf.Cortex.RuntimeConfig
		}

		controllerutil.SetOwnerReference(gateway, cluster, k.K8sClient.Scheme())
		return nil
	}

	if exists {
		err := retry.OnError(retry.DefaultBackoff, k8serrors.IsConflict, func() error {
			existing := k.newMonitoringCluster()
			err := k.K8sClient.Get(ctx, objectKey, existing)
			if err != nil {
				return err
			}
			clone := existing.DeepCopy()
			if err := mutator(clone); err != nil {
				return err
			}
			cmp, err := patch.DefaultPatchMaker.Calculate(existing, clone,
				patch.IgnoreStatusFields(),
				patch.IgnoreVolumeClaimTemplateTypeMetaAndStatus(),
				patch.IgnorePDBSelector(),
			)
			if err == nil {
				if cmp.IsEmpty() {
					return status.Error(codes.FailedPrecondition, "no changes to apply")
				}
			}

			return k.K8sClient.Update(ctx, clone)
		})
		if err != nil {
			return nil, fmt.Errorf("failed to update monitoring cluster: %w", err)
		}
	} else {
		if err := mutator(cluster); err != nil {
			return nil, err
		}
		err := k.K8sClient.Create(ctx, cluster)
		if err != nil {
			return nil, fmt.Errorf("failed to create monitoring cluster: %w", err)
		}
	}

	return &emptypb.Empty{}, nil
}

func mergeApiWorkloadsToK8s(k8sWorkloads *opnicorev1beta1.CortexWorkloadsSpec, apiWorkloads *cortexops.CortexWorkloadsSpec) error {
	for _, w := range []lo.Tuple2[**opnicorev1beta1.CortexWorkloadSpec, *cortexops.CortexWorkloadSpec]{
		lo.T2(&k8sWorkloads.Distributor, apiWorkloads.Distributor),
		lo.T2(&k8sWorkloads.Ingester, apiWorkloads.Ingester),
		lo.T2(&k8sWorkloads.Compactor, apiWorkloads.Compactor),
		lo.T2(&k8sWorkloads.StoreGateway, apiWorkloads.StoreGateway),
		lo.T2(&k8sWorkloads.Ruler, apiWorkloads.Ruler),
		lo.T2(&k8sWorkloads.QueryFrontend, apiWorkloads.QueryFrontend),
		lo.T2(&k8sWorkloads.Querier, apiWorkloads.Querier),
		lo.T2(&k8sWorkloads.Purger, apiWorkloads.Purger),
	} {
		pDest, src := w.Unpack()
		if src == nil {
			continue
		}
		if *pDest == nil {
			*pDest = &opnicorev1beta1.CortexWorkloadSpec{}
		}
		dest := *pDest

		dest.Replicas = src.Replicas
		dest.ExtraArgs = src.ExtraArgs

		jsonData, err := src.JsonData.MarshalJSON()
		if err != nil {
			return err
		}
		if err := json.Unmarshal(jsonData, dest); err != nil {
			return err
		}
	}
	return nil
}

func mergeK8sWorkloadsToApi(apiWorkloads *cortexops.CortexWorkloadsSpec, k8sWorkloads *opnicorev1beta1.CortexWorkloadsSpec) error {
	for _, w := range []lo.Tuple2[**cortexops.CortexWorkloadSpec, *opnicorev1beta1.CortexWorkloadSpec]{
		lo.T2(&apiWorkloads.Distributor, k8sWorkloads.Distributor),
		lo.T2(&apiWorkloads.Ingester, k8sWorkloads.Ingester),
		lo.T2(&apiWorkloads.Compactor, k8sWorkloads.Compactor),
		lo.T2(&apiWorkloads.StoreGateway, k8sWorkloads.StoreGateway),
		lo.T2(&apiWorkloads.Ruler, k8sWorkloads.Ruler),
		lo.T2(&apiWorkloads.QueryFrontend, k8sWorkloads.QueryFrontend),
		lo.T2(&apiWorkloads.Querier, k8sWorkloads.Querier),
		lo.T2(&apiWorkloads.Purger, k8sWorkloads.Purger),
	} {
		pDest, src := w.Unpack()
		if src == nil {
			continue
		}
		if *pDest == nil {
			*pDest = &cortexops.CortexWorkloadSpec{}
		}
		dest := *pDest

		dest.Replicas = src.Replicas
		dest.ExtraArgs = src.ExtraArgs

		jsonData, err := json.Marshal(src)
		if err != nil {
			return err
		}
		if dest.JsonData == nil {
			dest.JsonData = &structpb.Struct{Fields: map[string]*structpb.Value{}}
		}
		if err := dest.JsonData.UnmarshalJSON(jsonData); err != nil {
			return err
		}
	}
	return nil
}

func (k *OpniManager) GetClusterStatus(ctx context.Context, _ *emptypb.Empty) (*cortexops.InstallStatus, error) {
	metadata := map[string]string{}
	var state cortexops.InstallState
	var version string

	cluster := k.newMonitoringCluster()
	err := k.K8sClient.Get(ctx, k.MonitoringCluster, cluster)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			state = cortexops.InstallState_NotInstalled
		} else {
			return nil, fmt.Errorf("failed to get monitoring cluster: %w", err)
		}
	} else {
		status := cluster.Status.Cortex
		if err != nil {
			return nil, err
		}
		version = status.Version
		if cluster.GetDeletionTimestamp() != nil {
			state = cortexops.InstallState_Uninstalling
		} else {
			if status.WorkloadsReady {
				state = cortexops.InstallState_Installed
			} else {
				state = cortexops.InstallState_Updating
				metadata["Conditions"] = strings.Join(status.Conditions, "; ")
			}
		}
	}

	return &cortexops.InstallStatus{
		State:   state,
		Version: version,
		Metadata: lo.Assign(metadata, map[string]string{
			"Driver": "opni-manager",
		}),
	}, nil
}

func (k *OpniManager) UninstallCluster(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	cluster := k.newMonitoringCluster()
	err := k.K8sClient.Get(ctx, k.MonitoringCluster, cluster)
	if err != nil {
		return nil, fmt.Errorf("failed to uninstall monitoring cluster: %w", err)
	}

	err = k.K8sClient.Delete(ctx, cluster)
	if err != nil {
		return nil, fmt.Errorf("failed to uninstall monitoring cluster: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (k *OpniManager) ShouldDisableNode(_ *corev1.Reference) error {
	stat, err := k.GetClusterStatus(context.TODO(), &emptypb.Empty{})
	if err != nil {
		// can't determine cluster status, so don't disable the node
		return nil
	}
	switch stat.State {
	case cortexops.InstallState_NotInstalled, cortexops.InstallState_Uninstalling:
		return status.Error(codes.Unavailable, fmt.Sprintf("Cortex cluster is not installed"))
	case cortexops.InstallState_Updating, cortexops.InstallState_Installed:
		return nil
	case cortexops.InstallState_Unknown:
		fallthrough
	default:
		// can't determine cluster status, so don't disable the node
		return nil
	}
}

func init() {
	drivers.ClusterDrivers.Register("opni-manager", func(_ context.Context, opts ...driverutil.Option) (drivers.ClusterDriver, error) {
		options := OpniManagerClusterDriverOptions{
			MonitoringCluster: types.NamespacedName{
				Namespace: os.Getenv("POD_NAMESPACE"),
				Name:      "opni",
			},
			GatewayRef: types.NamespacedName{
				Namespace: os.Getenv("POD_NAMESPACE"),
				Name:      os.Getenv("GATEWAY_NAME"),
			},
		}
		if err := driverutil.ApplyOptions(&options, opts...); err != nil {
			return nil, err
		}

		return NewOpniManagerClusterDriver(options)
	})
}
