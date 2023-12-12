// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v1.0.0
// source: github.com/rancher/opni/pkg/apis/alerting/v1/alerting.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OpniSeverity int32

const (
	OpniSeverity_Info     OpniSeverity = 0
	OpniSeverity_Warning  OpniSeverity = 1
	OpniSeverity_Error    OpniSeverity = 2
	OpniSeverity_Critical OpniSeverity = 3
)

// Enum value maps for OpniSeverity.
var (
	OpniSeverity_name = map[int32]string{
		0: "Info",
		1: "Warning",
		2: "Error",
		3: "Critical",
	}
	OpniSeverity_value = map[string]int32{
		"Info":     0,
		"Warning":  1,
		"Error":    2,
		"Critical": 3,
	}
)

func (x OpniSeverity) Enum() *OpniSeverity {
	p := new(OpniSeverity)
	*p = x
	return p
}

func (x OpniSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpniSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_enumTypes[0].Descriptor()
}

func (OpniSeverity) Type() protoreflect.EnumType {
	return &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_enumTypes[0]
}

func (x OpniSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpniSeverity.Descriptor instead.
func (OpniSeverity) EnumDescriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescGZIP(), []int{0}
}

type GoldenSignal int32

const (
	GoldenSignal_Custom     GoldenSignal = 0
	GoldenSignal_Errors     GoldenSignal = 1
	GoldenSignal_Saturation GoldenSignal = 2
	GoldenSignal_Traffic    GoldenSignal = 3
	GoldenSignal_Latency    GoldenSignal = 4
)

// Enum value maps for GoldenSignal.
var (
	GoldenSignal_name = map[int32]string{
		0: "Custom",
		1: "Errors",
		2: "Saturation",
		3: "Traffic",
		4: "Latency",
	}
	GoldenSignal_value = map[string]int32{
		"Custom":     0,
		"Errors":     1,
		"Saturation": 2,
		"Traffic":    3,
		"Latency":    4,
	}
)

func (x GoldenSignal) Enum() *GoldenSignal {
	p := new(GoldenSignal)
	*p = x
	return p
}

func (x GoldenSignal) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoldenSignal) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_enumTypes[1].Descriptor()
}

func (GoldenSignal) Type() protoreflect.EnumType {
	return &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_enumTypes[1]
}

func (x GoldenSignal) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GoldenSignal.Descriptor instead.
func (GoldenSignal) EnumDescriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescGZIP(), []int{1}
}

type ConditionReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// empty string represents the default group
	GroupId string `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *ConditionReference) Reset() {
	*x = ConditionReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConditionReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionReference) ProtoMessage() {}

func (x *ConditionReference) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionReference.ProtoReflect.Descriptor instead.
func (*ConditionReference) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescGZIP(), []int{0}
}

func (x *ConditionReference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConditionReference) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type ConditionReferenceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ConditionReference `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ConditionReferenceList) Reset() {
	*x = ConditionReferenceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConditionReferenceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionReferenceList) ProtoMessage() {}

func (x *ConditionReferenceList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionReferenceList.ProtoReflect.Descriptor instead.
func (*ConditionReferenceList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescGZIP(), []int{1}
}

func (x *ConditionReferenceList) GetItems() []*ConditionReference {
	if x != nil {
		return x.Items
	}
	return nil
}

type CachedState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Healthy   bool                   `protobuf:"varint,1,opt,name=healthy,proto3" json:"healthy,omitempty"`
	Firing    bool                   `protobuf:"varint,2,opt,name=firing,proto3" json:"firing,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Metadata  map[string]string      `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CachedState) Reset() {
	*x = CachedState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CachedState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CachedState) ProtoMessage() {}

func (x *CachedState) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CachedState.ProtoReflect.Descriptor instead.
func (*CachedState) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescGZIP(), []int{2}
}

func (x *CachedState) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *CachedState) GetFiring() bool {
	if x != nil {
		return x.Firing
	}
	return false
}

func (x *CachedState) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CachedState) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type IncidentIntervals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Interval `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *IncidentIntervals) Reset() {
	*x = IncidentIntervals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncidentIntervals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentIntervals) ProtoMessage() {}

func (x *IncidentIntervals) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentIntervals.ProtoReflect.Descriptor instead.
func (*IncidentIntervals) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescGZIP(), []int{3}
}

func (x *IncidentIntervals) GetItems() []*Interval {
	if x != nil {
		return x.Items
	}
	return nil
}

type Interval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start        *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End          *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	Fingerprints []string               `protobuf:"bytes,3,rep,name=fingerprints,proto3" json:"fingerprints,omitempty"`
}

func (x *Interval) Reset() {
	*x = Interval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interval) ProtoMessage() {}

func (x *Interval) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interval.ProtoReflect.Descriptor instead.
func (*Interval) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescGZIP(), []int{4}
}

func (x *Interval) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Interval) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Interval) GetFingerprints() []string {
	if x != nil {
		return x.Fingerprints
	}
	return nil
}

var File_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x12, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x16, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xf7, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x66, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x3f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x3d, 0x0a, 0x11, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x8e, 0x01, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x30,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x73, 0x2a, 0x3e, 0x0a, 0x0c, 0x4f, 0x70, 0x6e, 0x69, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x10, 0x03, 0x2a, 0x50, 0x0a, 0x0c, 0x47, 0x6f, 0x6c, 0x64, 0x65, 0x6e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x61,
	0x74, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x10, 0x04, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescData = file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_goTypes = []interface{}{
	(OpniSeverity)(0),              // 0: alerting.OpniSeverity
	(GoldenSignal)(0),              // 1: alerting.GoldenSignal
	(*ConditionReference)(nil),     // 2: alerting.ConditionReference
	(*ConditionReferenceList)(nil), // 3: alerting.ConditionReferenceList
	(*CachedState)(nil),            // 4: alerting.CachedState
	(*IncidentIntervals)(nil),      // 5: alerting.IncidentIntervals
	(*Interval)(nil),               // 6: alerting.Interval
	nil,                            // 7: alerting.CachedState.MetadataEntry
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
}
var file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_depIdxs = []int32{
	2, // 0: alerting.ConditionReferenceList.items:type_name -> alerting.ConditionReference
	8, // 1: alerting.CachedState.timestamp:type_name -> google.protobuf.Timestamp
	7, // 2: alerting.CachedState.metadata:type_name -> alerting.CachedState.MetadataEntry
	6, // 3: alerting.IncidentIntervals.items:type_name -> alerting.Interval
	8, // 4: alerting.Interval.start:type_name -> google.protobuf.Timestamp
	8, // 5: alerting.Interval.end:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_init() }
func file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_init() {
	if File_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionReference); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionReferenceList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CachedState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncidentIntervals); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interval); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_depIdxs,
		EnumInfos:         file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_enumTypes,
		MessageInfos:      file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto = out.File
	file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_apis_alerting_v1_alerting_proto_depIdxs = nil
}
