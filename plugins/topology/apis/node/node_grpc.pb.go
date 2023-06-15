// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/topology/apis/node/node.proto

package node

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NodeTopologyCapability_Sync_FullMethodName = "/node.topology.NodeTopologyCapability/Sync"
)

// NodeTopologyCapabilityClient is the client API for NodeTopologyCapability service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeTopologyCapabilityClient interface {
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
}

type nodeTopologyCapabilityClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeTopologyCapabilityClient(cc grpc.ClientConnInterface) NodeTopologyCapabilityClient {
	return &nodeTopologyCapabilityClient{cc}
}

func (c *nodeTopologyCapabilityClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, NodeTopologyCapability_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeTopologyCapabilityServer is the server API for NodeTopologyCapability service.
// All implementations must embed UnimplementedNodeTopologyCapabilityServer
// for forward compatibility
type NodeTopologyCapabilityServer interface {
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	mustEmbedUnimplementedNodeTopologyCapabilityServer()
}

// UnimplementedNodeTopologyCapabilityServer must be embedded to have forward compatible implementations.
type UnimplementedNodeTopologyCapabilityServer struct {
}

func (UnimplementedNodeTopologyCapabilityServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedNodeTopologyCapabilityServer) mustEmbedUnimplementedNodeTopologyCapabilityServer() {
}

// UnsafeNodeTopologyCapabilityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeTopologyCapabilityServer will
// result in compilation errors.
type UnsafeNodeTopologyCapabilityServer interface {
	mustEmbedUnimplementedNodeTopologyCapabilityServer()
}

func RegisterNodeTopologyCapabilityServer(s grpc.ServiceRegistrar, srv NodeTopologyCapabilityServer) {
	s.RegisterService(&NodeTopologyCapability_ServiceDesc, srv)
}

func _NodeTopologyCapability_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTopologyCapabilityServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeTopologyCapability_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTopologyCapabilityServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeTopologyCapability_ServiceDesc is the grpc.ServiceDesc for NodeTopologyCapability service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeTopologyCapability_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.topology.NodeTopologyCapability",
	HandlerType: (*NodeTopologyCapabilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _NodeTopologyCapability_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/topology/apis/node/node.proto",
}
