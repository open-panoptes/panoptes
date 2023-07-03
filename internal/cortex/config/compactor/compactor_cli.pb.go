// Code generated by internal/codegen. DO NOT EDIT.

// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/internal/cortex/config/compactor/compactor.proto

package compactor

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
	fs.Var(flagutil.DurationpbSliceValue([]time.Duration{2 * time.Hour, 12 * time.Hour, 24 * time.Hour}, &in.BlockRanges), strings.Join(append(prefix, "block-ranges"), "."), "List of compaction time ranges.")
	fs.Int32Var(&in.BlockSyncConcurrency, strings.Join(append(prefix, "block-sync-concurrency"), "."), 20, "Number of Go routines to use when syncing block index and chunks files from the long term storage.")
	fs.Int32Var(&in.MetaSyncConcurrency, strings.Join(append(prefix, "meta-sync-concurrency"), "."), 20, "Number of Go routines to use when syncing block meta files from the long term storage.")
	fs.Var(flagutil.DurationpbValue(0, &in.ConsistencyDelay), strings.Join(append(prefix, "consistency-delay"), "."), "Minimum age of fresh (non-compacted) blocks before they are being processed. Malformed blocks older than the maximum of consistency-delay and 48h0m0s will be removed.")
	fs.Var(flagutil.DurationpbValue(1*time.Hour, &in.CompactionInterval), strings.Join(append(prefix, "compaction-interval"), "."), "The frequency at which the compaction runs")
	fs.Int32Var(&in.CompactionRetries, strings.Join(append(prefix, "compaction-retries"), "."), 3, "How many times to retry a failed compaction within a single compaction run.")
	fs.Int32Var(&in.CompactionConcurrency, strings.Join(append(prefix, "compaction-concurrency"), "."), 1, "Max number of concurrent compactions running.")
	fs.Var(flagutil.DurationpbValue(15*time.Minute, &in.CleanupInterval), strings.Join(append(prefix, "cleanup-interval"), "."), "How frequently compactor should run blocks cleanup and maintenance, as well as update the bucket index.")
	fs.Int32Var(&in.CleanupConcurrency, strings.Join(append(prefix, "cleanup-concurrency"), "."), 20, "Max number of tenants for which blocks cleanup and maintenance should run concurrently.")
	fs.Var(flagutil.DurationpbValue(12*time.Hour, &in.DeletionDelay), strings.Join(append(prefix, "deletion-delay"), "."), "Time before a block marked for deletion is deleted from bucket. If not 0, blocks will be marked for deletion and compactor component will permanently delete blocks marked for deletion from the bucket. If 0, blocks will be deleted straight away. Note that deleting blocks immediately can cause query failures.")
	fs.Var(flagutil.DurationpbValue(6*time.Hour, &in.TenantCleanupDelay), strings.Join(append(prefix, "tenant-cleanup-delay"), "."), "For tenants marked for deletion, this is time between deleting of last block, and doing final cleanup (marker files, debug files) of the tenant.")
	fs.BoolVar(&in.SkipBlocksWithOutOfOrderChunksEnabled, strings.Join(append(prefix, "skip-blocks-with-out-of-order-chunks-enabled"), "."), false, "When enabled, mark blocks containing index with out-of-order chunks for no compact instead of halting the compaction.")
	fs.Int32Var(&in.BlockFilesConcurrency, strings.Join(append(prefix, "block-files-concurrency"), "."), 10, "Number of goroutines to use when fetching/uploading block files from object storage.")
	fs.Int32Var(&in.BlocksFetchConcurrency, strings.Join(append(prefix, "blocks-fetch-concurrency"), "."), 3, "Number of goroutines to use when fetching blocks from object storage when compacting.")
	fs.BoolVar(&in.BlockDeletionMarksMigrationEnabled, strings.Join(append(prefix, "block-deletion-marks-migration-enabled"), "."), false, "When enabled, at compactor startup the bucket will be scanned and all found deletion marks inside the block location will be copied to the markers global location too. This option can (and should) be safely disabled as soon as the compactor has successfully run at least once.")
	fs.StringSliceVar(&in.EnabledTenants, strings.Join(append(prefix, "enabled-tenants"), "."), nil, "Comma separated list of tenants that can be compacted. If specified, only these tenants will be compacted by compactor, otherwise all tenants can be compacted. Subject to sharding.")
	fs.StringSliceVar(&in.DisabledTenants, strings.Join(append(prefix, "disabled-tenants"), "."), nil, "Comma separated list of tenants that cannot be compacted by this compactor. If specified, and compactor would normally pick given tenant for compaction (via -compactor.enabled-tenants or sharding), it will be ignored instead.")
	fs.Var(flagutil.DurationpbValue(5*time.Minute, &in.BlockVisitMarkerTimeout), strings.Join(append(prefix, "block-visit-marker-timeout"), "."), "How long block visit marker file should be considered as expired and able to be picked up by compactor again.")
	fs.Var(flagutil.DurationpbValue(1*time.Minute, &in.BlockVisitMarkerFileUpdateInterval), strings.Join(append(prefix, "block-visit-marker-file-update-interval"), "."), "How frequently block visit marker file should be updated duration compaction.")
	fs.BoolVar(&in.AcceptMalformedIndex, strings.Join(append(prefix, "accept-malformed-index"), "."), false, "When enabled, index verification will ignore out of order label names.")
	return fs
}
