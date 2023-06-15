// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v1.0.0
// source: github.com/rancher/opni/plugins/aiops/apis/modeltraining/modeltraining.proto

package modeltraining

import (
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GPUInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity    string `protobuf:"bytes,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Allocatable string `protobuf:"bytes,3,opt,name=allocatable,proto3" json:"allocatable,omitempty"`
}

func (x *GPUInfo) Reset() {
	*x = GPUInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPUInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPUInfo) ProtoMessage() {}

func (x *GPUInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPUInfo.ProtoReflect.Descriptor instead.
func (*GPUInfo) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP(), []int{0}
}

func (x *GPUInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GPUInfo) GetCapacity() string {
	if x != nil {
		return x.Capacity
	}
	return ""
}

func (x *GPUInfo) GetAllocatable() string {
	if x != nil {
		return x.Allocatable
	}
	return ""
}

type GPUInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GPUInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GPUInfoList) Reset() {
	*x = GPUInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPUInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPUInfoList) ProtoMessage() {}

func (x *GPUInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPUInfoList.ProtoReflect.Descriptor instead.
func (*GPUInfoList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP(), []int{1}
}

func (x *GPUInfoList) GetItems() []*GPUInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type ModelStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string                   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Statistics *ModelTrainingStatistics `protobuf:"bytes,2,opt,name=statistics,proto3" json:"statistics,omitempty"`
}

func (x *ModelStatus) Reset() {
	*x = ModelStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelStatus) ProtoMessage() {}

func (x *ModelStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelStatus.ProtoReflect.Descriptor instead.
func (*ModelStatus) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP(), []int{2}
}

func (x *ModelStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ModelStatus) GetStatistics() *ModelTrainingStatistics {
	if x != nil {
		return x.Statistics
	}
	return nil
}

type ModelTrainingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ModelTrainingResponse) Reset() {
	*x = ModelTrainingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelTrainingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelTrainingResponse) ProtoMessage() {}

func (x *ModelTrainingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelTrainingResponse.ProtoReflect.Descriptor instead.
func (*ModelTrainingResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP(), []int{3}
}

func (x *ModelTrainingResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type ModelTrainingParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId  string `protobuf:"bytes,1,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	Namespace  string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Deployment string `protobuf:"bytes,3,opt,name=deployment,proto3" json:"deployment,omitempty"`
}

func (x *ModelTrainingParameters) Reset() {
	*x = ModelTrainingParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelTrainingParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelTrainingParameters) ProtoMessage() {}

func (x *ModelTrainingParameters) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelTrainingParameters.ProtoReflect.Descriptor instead.
func (*ModelTrainingParameters) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP(), []int{4}
}

func (x *ModelTrainingParameters) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ModelTrainingParameters) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ModelTrainingParameters) GetDeployment() string {
	if x != nil {
		return x.Deployment
	}
	return ""
}

type ModelTrainingParametersList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ModelTrainingParameters `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ModelTrainingParametersList) Reset() {
	*x = ModelTrainingParametersList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelTrainingParametersList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelTrainingParametersList) ProtoMessage() {}

func (x *ModelTrainingParametersList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelTrainingParametersList.ProtoReflect.Descriptor instead.
func (*ModelTrainingParametersList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP(), []int{5}
}

func (x *ModelTrainingParametersList) GetItems() []*ModelTrainingParameters {
	if x != nil {
		return x.Items
	}
	return nil
}

type ModelTrainingStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeElapsed         int64   `protobuf:"varint,1,opt,name=timeElapsed,proto3" json:"timeElapsed,omitempty"`
	PercentageCompleted int64   `protobuf:"varint,2,opt,name=percentageCompleted,proto3" json:"percentageCompleted,omitempty"`
	RemainingTime       int64   `protobuf:"varint,3,opt,name=remainingTime,proto3" json:"remainingTime,omitempty"`
	CurrentEpoch        int64   `protobuf:"varint,4,opt,name=currentEpoch,proto3" json:"currentEpoch,omitempty"`
	ModelAccuracy       float64 `protobuf:"fixed64,5,opt,name=modelAccuracy,proto3" json:"modelAccuracy,omitempty"`
	Stage               string  `protobuf:"bytes,6,opt,name=stage,proto3" json:"stage,omitempty"`
}

func (x *ModelTrainingStatistics) Reset() {
	*x = ModelTrainingStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelTrainingStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelTrainingStatistics) ProtoMessage() {}

func (x *ModelTrainingStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelTrainingStatistics.ProtoReflect.Descriptor instead.
func (*ModelTrainingStatistics) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP(), []int{6}
}

func (x *ModelTrainingStatistics) GetTimeElapsed() int64 {
	if x != nil {
		return x.TimeElapsed
	}
	return 0
}

func (x *ModelTrainingStatistics) GetPercentageCompleted() int64 {
	if x != nil {
		return x.PercentageCompleted
	}
	return 0
}

func (x *ModelTrainingStatistics) GetRemainingTime() int64 {
	if x != nil {
		return x.RemainingTime
	}
	return 0
}

func (x *ModelTrainingStatistics) GetCurrentEpoch() int64 {
	if x != nil {
		return x.CurrentEpoch
	}
	return 0
}

func (x *ModelTrainingStatistics) GetModelAccuracy() float64 {
	if x != nil {
		return x.ModelAccuracy
	}
	return 0
}

func (x *ModelTrainingStatistics) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

type WorkloadAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId  string `protobuf:"bytes,1,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	Namespace  string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Deployment string `protobuf:"bytes,3,opt,name=deployment,proto3" json:"deployment,omitempty"`
	LogCount   int64  `protobuf:"varint,4,opt,name=logCount,proto3" json:"logCount,omitempty"`
}

func (x *WorkloadAggregation) Reset() {
	*x = WorkloadAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadAggregation) ProtoMessage() {}

func (x *WorkloadAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadAggregation.ProtoReflect.Descriptor instead.
func (*WorkloadAggregation) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP(), []int{7}
}

func (x *WorkloadAggregation) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *WorkloadAggregation) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WorkloadAggregation) GetDeployment() string {
	if x != nil {
		return x.Deployment
	}
	return ""
}

func (x *WorkloadAggregation) GetLogCount() int64 {
	if x != nil {
		return x.LogCount
	}
	return 0
}

type WorkloadAggregationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*WorkloadAggregation `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *WorkloadAggregationList) Reset() {
	*x = WorkloadAggregationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadAggregationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadAggregationList) ProtoMessage() {}

func (x *WorkloadAggregationList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadAggregationList.ProtoReflect.Descriptor instead.
func (*WorkloadAggregationList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP(), []int{8}
}

func (x *WorkloadAggregationList) GetItems() []*WorkloadAggregation {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x61, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a,
	0x07, 0x47, 0x50, 0x55, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x3b, 0x0a, 0x0b, 0x47, 0x50,
	0x55, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x50, 0x55, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x6d, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x46,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0x33, 0x0a, 0x15, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x0a, 0x17, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x1b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0xf3, 0x01, 0x0a, 0x17, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x69, 0x6d, 0x65, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x30, 0x0a,
	0x13, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x67,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xa9, 0x05, 0x0a, 0x0d, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x77, 0x0a, 0x0a,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2a, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x6e, 0x0a, 0x16, 0x50, 0x75, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x1a, 0x15,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x79, 0x0a, 0x1a, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x1a, 0x26, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x84, 0x01,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x50, 0x0a, 0x07, 0x47, 0x50, 0x55, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x50, 0x55, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x67, 0x70,
	0x75, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e,
	0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x61, 0x69, 0x6f, 0x70, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescData = file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDesc
)

func file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescData)
	})
	return file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDescData
}

var file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_goTypes = []interface{}{
	(*GPUInfo)(nil),                     // 0: modeltraining.GPUInfo
	(*GPUInfoList)(nil),                 // 1: modeltraining.GPUInfoList
	(*ModelStatus)(nil),                 // 2: modeltraining.ModelStatus
	(*ModelTrainingResponse)(nil),       // 3: modeltraining.ModelTrainingResponse
	(*ModelTrainingParameters)(nil),     // 4: modeltraining.ModelTrainingParameters
	(*ModelTrainingParametersList)(nil), // 5: modeltraining.ModelTrainingParametersList
	(*ModelTrainingStatistics)(nil),     // 6: modeltraining.ModelTrainingStatistics
	(*WorkloadAggregation)(nil),         // 7: modeltraining.WorkloadAggregation
	(*WorkloadAggregationList)(nil),     // 8: modeltraining.WorkloadAggregationList
	(*v1.Reference)(nil),                // 9: core.Reference
	(*emptypb.Empty)(nil),               // 10: google.protobuf.Empty
}
var file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_depIdxs = []int32{
	0,  // 0: modeltraining.GPUInfoList.items:type_name -> modeltraining.GPUInfo
	6,  // 1: modeltraining.ModelStatus.statistics:type_name -> modeltraining.ModelTrainingStatistics
	4,  // 2: modeltraining.ModelTrainingParametersList.items:type_name -> modeltraining.ModelTrainingParameters
	7,  // 3: modeltraining.WorkloadAggregationList.items:type_name -> modeltraining.WorkloadAggregation
	5,  // 4: modeltraining.ModelTraining.TrainModel:input_type -> modeltraining.ModelTrainingParametersList
	2,  // 5: modeltraining.ModelTraining.PutModelTrainingStatus:input_type -> modeltraining.ModelStatus
	9,  // 6: modeltraining.ModelTraining.ClusterWorkloadAggregation:input_type -> core.Reference
	10, // 7: modeltraining.ModelTraining.GetModelStatus:input_type -> google.protobuf.Empty
	10, // 8: modeltraining.ModelTraining.GetModelTrainingParameters:input_type -> google.protobuf.Empty
	10, // 9: modeltraining.ModelTraining.GPUInfo:input_type -> google.protobuf.Empty
	3,  // 10: modeltraining.ModelTraining.TrainModel:output_type -> modeltraining.ModelTrainingResponse
	10, // 11: modeltraining.ModelTraining.PutModelTrainingStatus:output_type -> google.protobuf.Empty
	8,  // 12: modeltraining.ModelTraining.ClusterWorkloadAggregation:output_type -> modeltraining.WorkloadAggregationList
	2,  // 13: modeltraining.ModelTraining.GetModelStatus:output_type -> modeltraining.ModelStatus
	5,  // 14: modeltraining.ModelTraining.GetModelTrainingParameters:output_type -> modeltraining.ModelTrainingParametersList
	1,  // 15: modeltraining.ModelTraining.GPUInfo:output_type -> modeltraining.GPUInfoList
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_init() }
func file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_init() {
	if File_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GPUInfo); i {
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
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GPUInfoList); i {
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
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelStatus); i {
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
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelTrainingResponse); i {
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
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelTrainingParameters); i {
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
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelTrainingParametersList); i {
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
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelTrainingStatistics); i {
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
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadAggregation); i {
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
		file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadAggregationList); i {
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
			RawDescriptor: file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto = out.File
	file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_rawDesc = nil
	file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_goTypes = nil
	file_github_com_rancher_opni_plugins_aiops_apis_modeltraining_modeltraining_proto_depIdxs = nil
}
