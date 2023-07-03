// Code generated by internal/codegen. DO NOT EDIT.

// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/internal/cortex/config/querier/querier.proto

package querier

import (
	flagutil "github.com/rancher/opni/pkg/util/flagutil"
	pflag "github.com/spf13/pflag"
	proto "google.golang.org/protobuf/proto"
	strings "strings"
	time "time"
)

func (in *Config) DeepCopyInto(out *Config) {
	proto.Merge(in, out)
}

func (in *Config) DeepCopy() *Config {
	return proto.Clone(in).(*Config)
}

func (in *Config) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Config", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Int32Var(&in.MaxConcurrent, strings.Join(append(prefix, "max-concurrent"), "."), 20, "The maximum number of concurrent queries.")
	fs.Var(flagutil.DurationpbValue(2*time.Minute, &in.Timeout), strings.Join(append(prefix, "timeout"), "."), "The timeout for a query.")
	fs.BoolVar(&in.Iterators, strings.Join(append(prefix, "iterators"), "."), false, "Use iterators to execute query, as opposed to fully materialising the series in memory.")
	fs.BoolVar(&in.BatchIterators, strings.Join(append(prefix, "batch-iterators"), "."), true, "Use batch iterators to execute query, as opposed to fully materialising the series in memory.  Takes precedent over the -querier.iterators flag.")
	fs.BoolVar(&in.IngesterStreaming, strings.Join(append(prefix, "ingester-streaming"), "."), true, "Use streaming RPCs to query ingester.")
	fs.BoolVar(&in.IngesterMetadataStreaming, strings.Join(append(prefix, "ingester-metadata-streaming"), "."), false, "Use streaming RPCs for metadata APIs from ingester.")
	fs.Int32Var(&in.MaxSamples, strings.Join(append(prefix, "max-samples"), "."), 50000000, "Maximum number of samples a single query can load into memory.")
	fs.Var(flagutil.DurationpbValue(0, &in.QueryIngestersWithin), strings.Join(append(prefix, "query-ingesters-within"), "."), "Maximum lookback beyond which queries are not sent to ingester. 0 means all queries are sent to ingester.")
	fs.BoolVar(&in.QueryStoreForLabelsEnabled, strings.Join(append(prefix, "query-store-for-labels-enabled"), "."), false, "Deprecated (Querying long-term store for labels will be always enabled in the future.): Query long-term store for series, label values and label names APIs.")
	fs.BoolVar(&in.PerStepStatsEnabled, strings.Join(append(prefix, "per-step-stats-enabled"), "."), false, "Enable returning samples stats per steps in query response.")
	fs.Var(flagutil.DurationpbValue(0, &in.QueryStoreAfter), strings.Join(append(prefix, "query-store-after"), "."), "The time after which a metric should be queried from storage and not just ingesters. 0 means all queries are sent to store. When running the blocks storage, if this option is enabled, the time range of the query sent to the store will be manipulated to ensure the query end is not more recent than 'now - query-store-after'.")
	fs.Var(flagutil.DurationpbValue(10*time.Minute, &in.MaxQueryIntoFuture), strings.Join(append(prefix, "max-query-into-future"), "."), "Maximum duration into the future you can query. 0 to disable.")
	fs.Var(flagutil.DurationpbValue(1*time.Minute, &in.DefaultEvaluationInterval), strings.Join(append(prefix, "default-evaluation-interval"), "."), "The default evaluation interval or step size for subqueries.")
	fs.Var(flagutil.DurationpbValue(5*time.Minute, &in.LookbackDelta), strings.Join(append(prefix, "lookback-delta"), "."), "Time since the last sample after which a time series is considered stale and ignored by expression evaluations.")
	fs.Var(flagutil.DurationpbValue(0, &in.ShuffleShardingIngestersLookbackPeriod), strings.Join(append(prefix, "shuffle-sharding-ingesters-lookback-period"), "."), "When distributor's sharding strategy is shuffle-sharding and this setting is > 0, queriers fetch in-memory series from the minimum set of required ingesters, selecting only ingesters which may have received series since 'now - lookback period'. The lookback period should be greater or equal than the configured 'query store after' and 'query ingesters within'. If this setting is 0, queriers always query all ingesters (ingesters shuffle sharding on read path is disabled).")
	fs.BoolVar(&in.ThanosEngine, strings.Join(append(prefix, "thanos-engine"), "."), false, "Experimental. Use Thanos promql engine https://github.com/thanos-io/promql-engine rather than the Prometheus promql engine.")
	return fs
}
