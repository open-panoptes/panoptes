// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v1.0.0
// source: github.com/open-panoptes/opni/pkg/apis/control/v1/remote.proto

package v1

import (
	_ "github.com/kralicky/totem"
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

type PatchOp int32

const (
	// revisions match
	PatchOp_None PatchOp = 0
	// same plugin exists on both
	PatchOp_Update PatchOp = 1
	// missing plugin on agent
	PatchOp_Create PatchOp = 2
	// outdated plugin on agent, with no version on the gateway
	PatchOp_Remove PatchOp = 3
	// same plugin contents and module name, but the file name is different
	PatchOp_Rename PatchOp = 4
)

// Enum value maps for PatchOp.
var (
	PatchOp_name = map[int32]string{
		0: "None",
		1: "Update",
		2: "Create",
		3: "Remove",
		4: "Rename",
	}
	PatchOp_value = map[string]int32{
		"None":   0,
		"Update": 1,
		"Create": 2,
		"Remove": 3,
		"Rename": 4,
	}
)

func (x PatchOp) Enum() *PatchOp {
	p := new(PatchOp)
	*p = x
	return p
}

func (x PatchOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PatchOp) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_enumTypes[0].Descriptor()
}

func (PatchOp) Type() protoreflect.EnumType {
	return &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_enumTypes[0]
}

func (x PatchOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PatchOp.Descriptor instead.
func (PatchOp) EnumDescriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescGZIP(), []int{0}
}

type SyncResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequiredPatches *PatchList `protobuf:"bytes,2,opt,name=requiredPatches,proto3" json:"requiredPatches,omitempty"`
}

func (x *SyncResults) Reset() {
	*x = SyncResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResults) ProtoMessage() {}

func (x *SyncResults) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResults.ProtoReflect.Descriptor instead.
func (*SyncResults) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescGZIP(), []int{0}
}

func (x *SyncResults) GetRequiredPatches() *PatchList {
	if x != nil {
		return x.RequiredPatches
	}
	return nil
}

type UpdateManifestEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	Path    string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Digest  string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *UpdateManifestEntry) Reset() {
	*x = UpdateManifestEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateManifestEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateManifestEntry) ProtoMessage() {}

func (x *UpdateManifestEntry) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateManifestEntry.ProtoReflect.Descriptor instead.
func (*UpdateManifestEntry) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateManifestEntry) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *UpdateManifestEntry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UpdateManifestEntry) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

type UpdateManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*UpdateManifestEntry `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *UpdateManifest) Reset() {
	*x = UpdateManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateManifest) ProtoMessage() {}

func (x *UpdateManifest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateManifest.ProtoReflect.Descriptor instead.
func (*UpdateManifest) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateManifest) GetItems() []*UpdateManifestEntry {
	if x != nil {
		return x.Items
	}
	return nil
}

type PluginArchiveEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *UpdateManifestEntry `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Data     []byte               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PluginArchiveEntry) Reset() {
	*x = PluginArchiveEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginArchiveEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginArchiveEntry) ProtoMessage() {}

func (x *PluginArchiveEntry) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginArchiveEntry.ProtoReflect.Descriptor instead.
func (*PluginArchiveEntry) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescGZIP(), []int{3}
}

func (x *PluginArchiveEntry) GetMetadata() *UpdateManifestEntry {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PluginArchiveEntry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PluginArchive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*PluginArchiveEntry `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PluginArchive) Reset() {
	*x = PluginArchive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginArchive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginArchive) ProtoMessage() {}

func (x *PluginArchive) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginArchive.ProtoReflect.Descriptor instead.
func (*PluginArchive) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescGZIP(), []int{4}
}

func (x *PluginArchive) GetItems() []*PluginArchiveEntry {
	if x != nil {
		return x.Items
	}
	return nil
}

// opPath should be empty when op == PatchRename
// bytes should be empty when op == PatchRemove
type PatchSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package   string  `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	Op        PatchOp `protobuf:"varint,2,opt,name=op,proto3,enum=control.PatchOp" json:"op,omitempty"`
	Data      []byte  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Path      string  `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	OldDigest string  `protobuf:"bytes,5,opt,name=oldDigest,proto3" json:"oldDigest,omitempty"`
	NewDigest string  `protobuf:"bytes,6,opt,name=newDigest,proto3" json:"newDigest,omitempty"`
}

func (x *PatchSpec) Reset() {
	*x = PatchSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchSpec) ProtoMessage() {}

func (x *PatchSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchSpec.ProtoReflect.Descriptor instead.
func (*PatchSpec) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescGZIP(), []int{5}
}

func (x *PatchSpec) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *PatchSpec) GetOp() PatchOp {
	if x != nil {
		return x.Op
	}
	return PatchOp_None
}

func (x *PatchSpec) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PatchSpec) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PatchSpec) GetOldDigest() string {
	if x != nil {
		return x.OldDigest
	}
	return ""
}

func (x *PatchSpec) GetNewDigest() string {
	if x != nil {
		return x.NewDigest
	}
	return ""
}

type PatchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*PatchSpec `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PatchList) Reset() {
	*x = PatchList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchList) ProtoMessage() {}

func (x *PatchList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchList.ProtoReflect.Descriptor instead.
func (*PatchList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescGZIP(), []int{6}
}

func (x *PatchList) GetItems() []*PatchSpec {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_github_com_rancher_opni_pkg_apis_control_v1_remote_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x1a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x72, 0x61, 0x6c, 0x69, 0x63, 0x6b, 0x79, 0x2f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4b, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x3c, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x5b,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x62, 0x0a, 0x12, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x42, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x09, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x20, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x52,
	0x02, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x6f,
	0x6c, 0x64, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x6c, 0x64, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x77,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65,
	0x77, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x70, 0x65, 0x63, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x43,
	0x0a, 0x07, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e,
	0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x10, 0x04, 0x32, 0x3b, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x31, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x32, 0x46, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x12, 0x34, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x4b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x3d, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescData = file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_goTypes = []interface{}{
	(PatchOp)(0),                // 0: control.PatchOp
	(*SyncResults)(nil),         // 1: control.SyncResults
	(*UpdateManifestEntry)(nil), // 2: control.UpdateManifestEntry
	(*UpdateManifest)(nil),      // 3: control.UpdateManifest
	(*PluginArchiveEntry)(nil),  // 4: control.PluginArchiveEntry
	(*PluginArchive)(nil),       // 5: control.PluginArchive
	(*PatchSpec)(nil),           // 6: control.PatchSpec
	(*PatchList)(nil),           // 7: control.PatchList
	(*emptypb.Empty)(nil),       // 8: google.protobuf.Empty
	(*v1.Health)(nil),           // 9: core.Health
}
var file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_depIdxs = []int32{
	7, // 0: control.SyncResults.requiredPatches:type_name -> control.PatchList
	2, // 1: control.UpdateManifest.items:type_name -> control.UpdateManifestEntry
	2, // 2: control.PluginArchiveEntry.metadata:type_name -> control.UpdateManifestEntry
	4, // 3: control.PluginArchive.items:type_name -> control.PluginArchiveEntry
	0, // 4: control.PatchSpec.op:type_name -> control.PatchOp
	6, // 5: control.PatchList.items:type_name -> control.PatchSpec
	8, // 6: control.Health.GetHealth:input_type -> google.protobuf.Empty
	9, // 7: control.HealthListener.UpdateHealth:input_type -> core.Health
	3, // 8: control.UpdateSync.SyncManifest:input_type -> control.UpdateManifest
	9, // 9: control.Health.GetHealth:output_type -> core.Health
	8, // 10: control.HealthListener.UpdateHealth:output_type -> google.protobuf.Empty
	1, // 11: control.UpdateSync.SyncManifest:output_type -> control.SyncResults
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_init() }
func file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_init() {
	if File_github_com_rancher_opni_pkg_apis_control_v1_remote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResults); i {
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
		file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateManifestEntry); i {
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
		file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateManifest); i {
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
		file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginArchiveEntry); i {
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
		file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginArchive); i {
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
		file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchSpec); i {
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
		file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchList); i {
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
			RawDescriptor: file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_depIdxs,
		EnumInfos:         file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_enumTypes,
		MessageInfos:      file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_apis_control_v1_remote_proto = out.File
	file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_apis_control_v1_remote_proto_depIdxs = nil
}
