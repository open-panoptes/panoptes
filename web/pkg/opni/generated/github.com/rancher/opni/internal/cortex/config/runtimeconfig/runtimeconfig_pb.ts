// Code generated by internal/codegen. DO NOT EDIT.

// @generated by protoc-gen-es v1.3.0 with parameter "target=ts,import_extension=none,ts_nocheck=false"
// @generated from file github.com/open-panoptes/opni/internal/cortex/config/runtimeconfig/runtimeconfig.proto (package runtimeconfig, syntax proto3)
/* eslint-disable */

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Limits } from "../validation/limits_pb";

/**
 * @generated from message runtimeconfig.IngesterInstanceLimits
 */
export class IngesterInstanceLimits extends Message<IngesterInstanceLimits> {
  /**
   * @generated from field: optional double max_ingestion_rate = 1;
   */
  maxIngestionRate?: number;

  /**
   * @generated from field: optional int64 max_tenants = 2;
   */
  maxTenants?: bigint;

  /**
   * @generated from field: optional int64 max_series = 3;
   */
  maxSeries?: bigint;

  /**
   * @generated from field: optional int64 max_inflight_push_requests = 4;
   */
  maxInflightPushRequests?: bigint;

  constructor(data?: PartialMessage<IngesterInstanceLimits>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "runtimeconfig.IngesterInstanceLimits";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "max_ingestion_rate", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 2, name: "max_tenants", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 3, name: "max_series", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 4, name: "max_inflight_push_requests", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngesterInstanceLimits {
    return new IngesterInstanceLimits().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngesterInstanceLimits {
    return new IngesterInstanceLimits().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngesterInstanceLimits {
    return new IngesterInstanceLimits().fromJsonString(jsonString, options);
  }

  static equals(a: IngesterInstanceLimits | PlainMessage<IngesterInstanceLimits> | undefined, b: IngesterInstanceLimits | PlainMessage<IngesterInstanceLimits> | undefined): boolean {
    return proto3.util.equals(IngesterInstanceLimits, a, b);
  }
}

/**
 * @generated from message runtimeconfig.KvMultiRuntimeConfig
 */
export class KvMultiRuntimeConfig extends Message<KvMultiRuntimeConfig> {
  /**
   * @generated from field: optional string primary = 1;
   */
  primary?: string;

  /**
   * @generated from field: optional bool mirror_enabled = 2;
   */
  mirrorEnabled?: boolean;

  constructor(data?: PartialMessage<KvMultiRuntimeConfig>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "runtimeconfig.KvMultiRuntimeConfig";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "primary", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "mirror_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KvMultiRuntimeConfig {
    return new KvMultiRuntimeConfig().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KvMultiRuntimeConfig {
    return new KvMultiRuntimeConfig().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KvMultiRuntimeConfig {
    return new KvMultiRuntimeConfig().fromJsonString(jsonString, options);
  }

  static equals(a: KvMultiRuntimeConfig | PlainMessage<KvMultiRuntimeConfig> | undefined, b: KvMultiRuntimeConfig | PlainMessage<KvMultiRuntimeConfig> | undefined): boolean {
    return proto3.util.equals(KvMultiRuntimeConfig, a, b);
  }
}

/**
 * @generated from message runtimeconfig.RuntimeConfigValues
 */
export class RuntimeConfigValues extends Message<RuntimeConfigValues> {
  /**
   * @generated from field: map<string, validation.Limits> overrides = 1;
   */
  overrides: { [key: string]: Limits } = {};

  /**
   * @generated from field: runtimeconfig.KvMultiRuntimeConfig multi_kv_config = 2;
   */
  multiKvConfig?: KvMultiRuntimeConfig;

  /**
   * @generated from field: optional bool ingester_stream_chunks_when_using_blocks = 3;
   */
  ingesterStreamChunksWhenUsingBlocks?: boolean;

  /**
   * @generated from field: runtimeconfig.IngesterInstanceLimits ingester_limits = 4;
   */
  ingesterLimits?: IngesterInstanceLimits;

  constructor(data?: PartialMessage<RuntimeConfigValues>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "runtimeconfig.RuntimeConfigValues";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "overrides", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Limits} },
    { no: 2, name: "multi_kv_config", kind: "message", T: KvMultiRuntimeConfig },
    { no: 3, name: "ingester_stream_chunks_when_using_blocks", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 4, name: "ingester_limits", kind: "message", T: IngesterInstanceLimits },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RuntimeConfigValues {
    return new RuntimeConfigValues().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RuntimeConfigValues {
    return new RuntimeConfigValues().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RuntimeConfigValues {
    return new RuntimeConfigValues().fromJsonString(jsonString, options);
  }

  static equals(a: RuntimeConfigValues | PlainMessage<RuntimeConfigValues> | undefined, b: RuntimeConfigValues | PlainMessage<RuntimeConfigValues> | undefined): boolean {
    return proto3.util.equals(RuntimeConfigValues, a, b);
  }
}

