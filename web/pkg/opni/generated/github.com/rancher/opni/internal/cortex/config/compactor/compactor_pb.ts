// Code generated by internal/codegen. You may edit parts of this file.
// Field numbers and custom options will be preserved for matching field names.
// All other modifications will be lost.

// @generated by protoc-gen-es v1.3.0 with parameter "target=ts,import_extension=none,ts_nocheck=false"
// @generated from file github.com/rancher/opni/internal/cortex/config/compactor/compactor.proto (package compactor, syntax proto3)
/* eslint-disable */

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message compactor.Config
 */
export class Config extends Message<Config> {
  /**
   * List of compaction time ranges.
   *
   * @generated from field: repeated google.protobuf.Duration block_ranges = 1;
   */
  blockRanges: Duration[] = [];

  /**
   * Number of Go routines to use when syncing block index and chunks files from the long term storage.
   *
   * @generated from field: optional int32 block_sync_concurrency = 2;
   */
  blockSyncConcurrency?: number;

  /**
   * Number of Go routines to use when syncing block meta files from the long term storage.
   *
   * @generated from field: optional int32 meta_sync_concurrency = 3;
   */
  metaSyncConcurrency?: number;

  /**
   * Minimum age of fresh (non-compacted) blocks before they are being processed. Malformed blocks older than the maximum of consistency-delay and 48h0m0s will be removed.
   *
   * @generated from field: google.protobuf.Duration consistency_delay = 4;
   */
  consistencyDelay?: Duration;

  /**
   * The frequency at which the compaction runs
   *
   * @generated from field: google.protobuf.Duration compaction_interval = 5;
   */
  compactionInterval?: Duration;

  /**
   * How many times to retry a failed compaction within a single compaction run.
   *
   * @generated from field: optional int32 compaction_retries = 6;
   */
  compactionRetries?: number;

  /**
   * Max number of concurrent compactions running.
   *
   * @generated from field: optional int32 compaction_concurrency = 7;
   */
  compactionConcurrency?: number;

  /**
   * How frequently compactor should run blocks cleanup and maintenance, as well as update the bucket index.
   *
   * @generated from field: google.protobuf.Duration cleanup_interval = 8;
   */
  cleanupInterval?: Duration;

  /**
   * Max number of tenants for which blocks cleanup and maintenance should run concurrently.
   *
   * @generated from field: optional int32 cleanup_concurrency = 9;
   */
  cleanupConcurrency?: number;

  /**
   * Time before a block marked for deletion is deleted from bucket. If not 0, blocks will be marked for deletion and compactor component will permanently delete blocks marked for deletion from the bucket. If 0, blocks will be deleted straight away. Note that deleting blocks immediately can cause query failures.
   *
   * @generated from field: google.protobuf.Duration deletion_delay = 10;
   */
  deletionDelay?: Duration;

  /**
   * For tenants marked for deletion, this is time between deleting of last block, and doing final cleanup (marker files, debug files) of the tenant.
   *
   * @generated from field: google.protobuf.Duration tenant_cleanup_delay = 11;
   */
  tenantCleanupDelay?: Duration;

  /**
   * When enabled, mark blocks containing index with out-of-order chunks for no compact instead of halting the compaction.
   *
   * @generated from field: optional bool skip_blocks_with_out_of_order_chunks_enabled = 12;
   */
  skipBlocksWithOutOfOrderChunksEnabled?: boolean;

  /**
   * Number of goroutines to use when fetching/uploading block files from object storage.
   *
   * @generated from field: optional int32 block_files_concurrency = 13;
   */
  blockFilesConcurrency?: number;

  /**
   * Number of goroutines to use when fetching blocks from object storage when compacting.
   *
   * @generated from field: optional int32 blocks_fetch_concurrency = 14;
   */
  blocksFetchConcurrency?: number;

  /**
   * When enabled, at compactor startup the bucket will be scanned and all found deletion marks inside the block location will be copied to the markers global location too. This option can (and should) be safely disabled as soon as the compactor has successfully run at least once.
   *
   * @generated from field: optional bool block_deletion_marks_migration_enabled = 15;
   */
  blockDeletionMarksMigrationEnabled?: boolean;

  /**
   * Comma separated list of tenants that can be compacted. If specified, only these tenants will be compacted by compactor, otherwise all tenants can be compacted. Subject to sharding.
   *
   * @generated from field: repeated string enabled_tenants = 16;
   */
  enabledTenants: string[] = [];

  /**
   * Comma separated list of tenants that cannot be compacted by this compactor. If specified, and compactor would normally pick given tenant for compaction (via -compactor.enabled-tenants or sharding), it will be ignored instead.
   *
   * @generated from field: repeated string disabled_tenants = 17;
   */
  disabledTenants: string[] = [];

  /**
   * How long block visit marker file should be considered as expired and able to be picked up by compactor again.
   *
   * @generated from field: google.protobuf.Duration block_visit_marker_timeout = 18;
   */
  blockVisitMarkerTimeout?: Duration;

  /**
   * How frequently block visit marker file should be updated duration compaction.
   *
   * @generated from field: google.protobuf.Duration block_visit_marker_file_update_interval = 19;
   */
  blockVisitMarkerFileUpdateInterval?: Duration;

  /**
   * When enabled, index verification will ignore out of order label names.
   *
   * @generated from field: optional bool accept_malformed_index = 20;
   */
  acceptMalformedIndex?: boolean;

  constructor(data?: PartialMessage<Config>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "compactor.Config";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "block_ranges", kind: "message", T: Duration, repeated: true },
    { no: 2, name: "block_sync_concurrency", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 3, name: "meta_sync_concurrency", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 4, name: "consistency_delay", kind: "message", T: Duration },
    { no: 5, name: "compaction_interval", kind: "message", T: Duration },
    { no: 6, name: "compaction_retries", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 7, name: "compaction_concurrency", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 8, name: "cleanup_interval", kind: "message", T: Duration },
    { no: 9, name: "cleanup_concurrency", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 10, name: "deletion_delay", kind: "message", T: Duration },
    { no: 11, name: "tenant_cleanup_delay", kind: "message", T: Duration },
    { no: 12, name: "skip_blocks_with_out_of_order_chunks_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 13, name: "block_files_concurrency", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 14, name: "blocks_fetch_concurrency", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 15, name: "block_deletion_marks_migration_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 16, name: "enabled_tenants", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 17, name: "disabled_tenants", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 18, name: "block_visit_marker_timeout", kind: "message", T: Duration },
    { no: 19, name: "block_visit_marker_file_update_interval", kind: "message", T: Duration },
    { no: 20, name: "accept_malformed_index", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Config {
    return new Config().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Config {
    return new Config().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Config {
    return new Config().fromJsonString(jsonString, options);
  }

  static equals(a: Config | PlainMessage<Config> | undefined, b: Config | PlainMessage<Config> | undefined): boolean {
    return proto3.util.equals(Config, a, b);
  }
}

