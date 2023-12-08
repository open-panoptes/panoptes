// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v1.0.0
// source: github.com/open-panoptes/opni/plugins/logging/apis/alerting/alerting.proto

package alerting

import (
	v1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
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

type Monitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonitorId   string `protobuf:"bytes,1,opt,name=monitorId,proto3" json:"monitorId,omitempty"`
	MonitorType string `protobuf:"bytes,2,opt,name=monitorType,proto3" json:"monitorType,omitempty"`
	Spec        []byte `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Monitor) Reset() {
	*x = Monitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitor) ProtoMessage() {}

func (x *Monitor) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitor.ProtoReflect.Descriptor instead.
func (*Monitor) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescGZIP(), []int{0}
}

func (x *Monitor) GetMonitorId() string {
	if x != nil {
		return x.MonitorId
	}
	return ""
}

func (x *Monitor) GetMonitorType() string {
	if x != nil {
		return x.MonitorType
	}
	return ""
}

func (x *Monitor) GetSpec() []byte {
	if x != nil {
		return x.Spec
	}
	return nil
}

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Spec      []byte `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescGZIP(), []int{1}
}

func (x *Channel) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Channel) GetSpec() []byte {
	if x != nil {
		return x.Spec
	}
	return nil
}

type ChannelList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []byte `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
}

func (x *ChannelList) Reset() {
	*x = ChannelList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelList) ProtoMessage() {}

func (x *ChannelList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelList.ProtoReflect.Descriptor instead.
func (*ChannelList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelList) GetList() []byte {
	if x != nil {
		return x.List
	}
	return nil
}

type ListAlertsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts []byte `protobuf:"bytes,1,opt,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *ListAlertsResponse) Reset() {
	*x = ListAlertsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlertsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlertsResponse) ProtoMessage() {}

func (x *ListAlertsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlertsResponse.ProtoReflect.Descriptor instead.
func (*ListAlertsResponse) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescGZIP(), []int{3}
}

func (x *ListAlertsResponse) GetAlerts() []byte {
	if x != nil {
		return x.Alerts
	}
	return nil
}

type AcknowledgeAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonitorId string   `protobuf:"bytes,1,opt,name=monitorId,proto3" json:"monitorId,omitempty"`
	AlertIds  []string `protobuf:"bytes,2,rep,name=alertIds,proto3" json:"alertIds,omitempty"`
}

func (x *AcknowledgeAlertRequest) Reset() {
	*x = AcknowledgeAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcknowledgeAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcknowledgeAlertRequest) ProtoMessage() {}

func (x *AcknowledgeAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcknowledgeAlertRequest.ProtoReflect.Descriptor instead.
func (*AcknowledgeAlertRequest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescGZIP(), []int{4}
}

func (x *AcknowledgeAlertRequest) GetMonitorId() string {
	if x != nil {
		return x.MonitorId
	}
	return ""
}

func (x *AcknowledgeAlertRequest) GetAlertIds() []string {
	if x != nil {
		return x.AlertIds
	}
	return nil
}

var File_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x3b, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x53, 0x0a, 0x17, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x64, 0x73, 0x32, 0xec, 0x02, 0x0a, 0x11, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x57, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x12, 0x19, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22,
	0x08, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x19, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5c, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x19, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x1a, 0x0d, 0x2f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x32, 0xf3, 0x03, 0x0a, 0x16, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x61, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x61, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x66, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x1a, 0x12, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x53, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x2a, 0x0d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0xe5, 0x01, 0x0a, 0x0f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x12, 0x75, 0x0a, 0x10, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x12, 0x29, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a,
	0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x6b, 0x6e,
	0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70,
	0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescData = file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDesc
)

func file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescData)
	})
	return file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDescData
}

var file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_goTypes = []interface{}{
	(*Monitor)(nil),                 // 0: alerting.logging.Monitor
	(*Channel)(nil),                 // 1: alerting.logging.Channel
	(*ChannelList)(nil),             // 2: alerting.logging.ChannelList
	(*ListAlertsResponse)(nil),      // 3: alerting.logging.ListAlertsResponse
	(*AcknowledgeAlertRequest)(nil), // 4: alerting.logging.AcknowledgeAlertRequest
	(*v1.Reference)(nil),            // 5: core.Reference
	(*emptypb.Empty)(nil),           // 6: google.protobuf.Empty
}
var file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_depIdxs = []int32{
	0,  // 0: alerting.logging.MonitorManagement.CreateMonitor:input_type -> alerting.logging.Monitor
	5,  // 1: alerting.logging.MonitorManagement.GetMonitor:input_type -> core.Reference
	0,  // 2: alerting.logging.MonitorManagement.UpdateMonitor:input_type -> alerting.logging.Monitor
	5,  // 3: alerting.logging.MonitorManagement.DeleteMonitor:input_type -> core.Reference
	1,  // 4: alerting.logging.NotificationManagement.CreateNotification:input_type -> alerting.logging.Channel
	5,  // 5: alerting.logging.NotificationManagement.GetNotification:input_type -> core.Reference
	6,  // 6: alerting.logging.NotificationManagement.ListNotifications:input_type -> google.protobuf.Empty
	1,  // 7: alerting.logging.NotificationManagement.UpdateNotification:input_type -> alerting.logging.Channel
	5,  // 8: alerting.logging.NotificationManagement.DeleteDestination:input_type -> core.Reference
	6,  // 9: alerting.logging.AlertManagement.ListAlerts:input_type -> google.protobuf.Empty
	4,  // 10: alerting.logging.AlertManagement.AcknowledgeAlert:input_type -> alerting.logging.AcknowledgeAlertRequest
	6,  // 11: alerting.logging.MonitorManagement.CreateMonitor:output_type -> google.protobuf.Empty
	0,  // 12: alerting.logging.MonitorManagement.GetMonitor:output_type -> alerting.logging.Monitor
	6,  // 13: alerting.logging.MonitorManagement.UpdateMonitor:output_type -> google.protobuf.Empty
	6,  // 14: alerting.logging.MonitorManagement.DeleteMonitor:output_type -> google.protobuf.Empty
	6,  // 15: alerting.logging.NotificationManagement.CreateNotification:output_type -> google.protobuf.Empty
	6,  // 16: alerting.logging.NotificationManagement.GetNotification:output_type -> google.protobuf.Empty
	2,  // 17: alerting.logging.NotificationManagement.ListNotifications:output_type -> alerting.logging.ChannelList
	6,  // 18: alerting.logging.NotificationManagement.UpdateNotification:output_type -> google.protobuf.Empty
	6,  // 19: alerting.logging.NotificationManagement.DeleteDestination:output_type -> google.protobuf.Empty
	3,  // 20: alerting.logging.AlertManagement.ListAlerts:output_type -> alerting.logging.ListAlertsResponse
	6,  // 21: alerting.logging.AlertManagement.AcknowledgeAlert:output_type -> google.protobuf.Empty
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_init() }
func file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_init() {
	if File_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monitor); i {
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
		file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelList); i {
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
		file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlertsResponse); i {
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
		file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcknowledgeAlertRequest); i {
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
			RawDescriptor: file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto = out.File
	file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_rawDesc = nil
	file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_goTypes = nil
	file_github_com_rancher_opni_plugins_logging_apis_alerting_alerting_proto_depIdxs = nil
}
