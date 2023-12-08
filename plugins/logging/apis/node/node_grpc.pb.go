// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/open-panoptes/opni/plugins/logging/apis/node/node.proto

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
	NodeLoggingCapability_Sync_FullMethodName = "/node.logging.NodeLoggingCapability/Sync"
)

// NodeLoggingCapabilityClient is the client API for NodeLoggingCapability service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeLoggingCapabilityClient interface {
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
}

type nodeLoggingCapabilityClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeLoggingCapabilityClient(cc grpc.ClientConnInterface) NodeLoggingCapabilityClient {
	return &nodeLoggingCapabilityClient{cc}
}

func (c *nodeLoggingCapabilityClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, NodeLoggingCapability_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeLoggingCapabilityServer is the server API for NodeLoggingCapability service.
// All implementations must embed UnimplementedNodeLoggingCapabilityServer
// for forward compatibility
type NodeLoggingCapabilityServer interface {
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	mustEmbedUnimplementedNodeLoggingCapabilityServer()
}

// UnimplementedNodeLoggingCapabilityServer must be embedded to have forward compatible implementations.
type UnimplementedNodeLoggingCapabilityServer struct {
}

func (UnimplementedNodeLoggingCapabilityServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedNodeLoggingCapabilityServer) mustEmbedUnimplementedNodeLoggingCapabilityServer() {}

// UnsafeNodeLoggingCapabilityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeLoggingCapabilityServer will
// result in compilation errors.
type UnsafeNodeLoggingCapabilityServer interface {
	mustEmbedUnimplementedNodeLoggingCapabilityServer()
}

func RegisterNodeLoggingCapabilityServer(s grpc.ServiceRegistrar, srv NodeLoggingCapabilityServer) {
	s.RegisterService(&NodeLoggingCapability_ServiceDesc, srv)
}

func _NodeLoggingCapability_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeLoggingCapabilityServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeLoggingCapability_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeLoggingCapabilityServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeLoggingCapability_ServiceDesc is the grpc.ServiceDesc for NodeLoggingCapability service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeLoggingCapability_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.logging.NodeLoggingCapability",
	HandlerType: (*NodeLoggingCapabilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _NodeLoggingCapability_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/open-panoptes/opni/plugins/logging/apis/node/node.proto",
}
