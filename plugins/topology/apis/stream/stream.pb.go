// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v1.0.0
// source: github.com/open-panoptes/opni/plugins/topology/apis/stream/stream.proto

package stream

import (
	v1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
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

type GraphRepr int32

const (
	GraphRepr_None         GraphRepr = 0
	GraphRepr_KubectlGraph GraphRepr = 1
)

// Enum value maps for GraphRepr.
var (
	GraphRepr_name = map[int32]string{
		0: "None",
		1: "KubectlGraph",
	}
	GraphRepr_value = map[string]int32{
		"None":         0,
		"KubectlGraph": 1,
	}
)

func (x GraphRepr) Enum() *GraphRepr {
	p := new(GraphRepr)
	*p = x
	return p
}

func (x GraphRepr) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GraphRepr) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_enumTypes[0].Descriptor()
}

func (GraphRepr) Type() protoreflect.EnumType {
	return &file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_enumTypes[0]
}

func (x GraphRepr) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GraphRepr.Descriptor instead.
func (GraphRepr) EnumDescriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescGZIP(), []int{0}
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Graph *TopologyGraph `protobuf:"bytes,1,opt,name=graph,proto3" json:"graph,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetGraph() *TopologyGraph {
	if x != nil {
		return x.Graph
	}
	return nil
}

// FIXME: copied from orchestrator.proto due to duplicate symbol bug
type TopologyGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId *v1.Reference `protobuf:"bytes,1,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	Data      []byte        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Repr      GraphRepr     `protobuf:"varint,3,opt,name=repr,proto3,enum=stream.topology.GraphRepr" json:"repr,omitempty"`
}

func (x *TopologyGraph) Reset() {
	*x = TopologyGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopologyGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopologyGraph) ProtoMessage() {}

func (x *TopologyGraph) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopologyGraph.ProtoReflect.Descriptor instead.
func (*TopologyGraph) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescGZIP(), []int{1}
}

func (x *TopologyGraph) GetClusterId() *v1.Reference {
	if x != nil {
		return x.ClusterId
	}
	return nil
}

func (x *TopologyGraph) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TopologyGraph) GetRepr() GraphRepr {
	if x != nil {
		return x.Repr
	}
	return GraphRepr_None
}

var File_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x74, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x34, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x22, 0x82, 0x01, 0x0a, 0x0d, 0x54, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x04,
	0x72, 0x65, 0x70, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x52, 0x65, 0x70, 0x72, 0x52, 0x04, 0x72, 0x65, 0x70, 0x72, 0x2a, 0x27, 0x0a, 0x09,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x70, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e,
	0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x74, 0x6c, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x10, 0x01, 0x32, 0x8c, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x38, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68,
	0x12, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x40, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x74, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescData = file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDesc
)

func file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescData)
	})
	return file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDescData
}

var file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_goTypes = []interface{}{
	(GraphRepr)(0),        // 0: stream.topology.GraphRepr
	(*Payload)(nil),       // 1: stream.topology.Payload
	(*TopologyGraph)(nil), // 2: stream.topology.TopologyGraph
	(*v1.Reference)(nil),  // 3: core.Reference
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_depIdxs = []int32{
	2, // 0: stream.topology.Payload.graph:type_name -> stream.topology.TopologyGraph
	3, // 1: stream.topology.TopologyGraph.clusterId:type_name -> core.Reference
	0, // 2: stream.topology.TopologyGraph.repr:type_name -> stream.topology.GraphRepr
	1, // 3: stream.topology.RemoteTopology.Push:input_type -> stream.topology.Payload
	1, // 4: stream.topology.RemoteTopology.SyncTopology:input_type -> stream.topology.Payload
	4, // 5: stream.topology.RemoteTopology.Push:output_type -> google.protobuf.Empty
	4, // 6: stream.topology.RemoteTopology.SyncTopology:output_type -> google.protobuf.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_init() }
func file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_init() {
	if File_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopologyGraph); i {
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
			RawDescriptor: file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_depIdxs,
		EnumInfos:         file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_enumTypes,
		MessageInfos:      file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto = out.File
	file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_rawDesc = nil
	file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_goTypes = nil
	file_github_com_rancher_opni_plugins_topology_apis_stream_stream_proto_depIdxs = nil
}
