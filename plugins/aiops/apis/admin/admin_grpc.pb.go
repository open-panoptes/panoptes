// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/aiops/apis/admin/admin.proto

package admin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AIAdmin_GetAISettings_FullMethodName     = "/admin.AIAdmin/GetAISettings"
	AIAdmin_PutAISettings_FullMethodName     = "/admin.AIAdmin/PutAISettings"
	AIAdmin_DeleteAISettings_FullMethodName  = "/admin.AIAdmin/DeleteAISettings"
	AIAdmin_UpgradeAvailable_FullMethodName  = "/admin.AIAdmin/UpgradeAvailable"
	AIAdmin_DoUpgrade_FullMethodName         = "/admin.AIAdmin/DoUpgrade"
	AIAdmin_GetRuntimeClasses_FullMethodName = "/admin.AIAdmin/GetRuntimeClasses"
)

// AIAdminClient is the client API for AIAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AIAdminClient interface {
	GetAISettings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AISettings, error)
	PutAISettings(ctx context.Context, in *AISettings, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAISettings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpgradeAvailable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UpgradeAvailableResponse, error)
	DoUpgrade(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRuntimeClasses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RuntimeClassResponse, error)
}

type aIAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewAIAdminClient(cc grpc.ClientConnInterface) AIAdminClient {
	return &aIAdminClient{cc}
}

func (c *aIAdminClient) GetAISettings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AISettings, error) {
	out := new(AISettings)
	err := c.cc.Invoke(ctx, AIAdmin_GetAISettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIAdminClient) PutAISettings(ctx context.Context, in *AISettings, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AIAdmin_PutAISettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIAdminClient) DeleteAISettings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AIAdmin_DeleteAISettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIAdminClient) UpgradeAvailable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UpgradeAvailableResponse, error) {
	out := new(UpgradeAvailableResponse)
	err := c.cc.Invoke(ctx, AIAdmin_UpgradeAvailable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIAdminClient) DoUpgrade(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AIAdmin_DoUpgrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIAdminClient) GetRuntimeClasses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RuntimeClassResponse, error) {
	out := new(RuntimeClassResponse)
	err := c.cc.Invoke(ctx, AIAdmin_GetRuntimeClasses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AIAdminServer is the server API for AIAdmin service.
// All implementations must embed UnimplementedAIAdminServer
// for forward compatibility
type AIAdminServer interface {
	GetAISettings(context.Context, *emptypb.Empty) (*AISettings, error)
	PutAISettings(context.Context, *AISettings) (*emptypb.Empty, error)
	DeleteAISettings(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	UpgradeAvailable(context.Context, *emptypb.Empty) (*UpgradeAvailableResponse, error)
	DoUpgrade(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetRuntimeClasses(context.Context, *emptypb.Empty) (*RuntimeClassResponse, error)
	mustEmbedUnimplementedAIAdminServer()
}

// UnimplementedAIAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAIAdminServer struct {
}

func (UnimplementedAIAdminServer) GetAISettings(context.Context, *emptypb.Empty) (*AISettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAISettings not implemented")
}
func (UnimplementedAIAdminServer) PutAISettings(context.Context, *AISettings) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutAISettings not implemented")
}
func (UnimplementedAIAdminServer) DeleteAISettings(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAISettings not implemented")
}
func (UnimplementedAIAdminServer) UpgradeAvailable(context.Context, *emptypb.Empty) (*UpgradeAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeAvailable not implemented")
}
func (UnimplementedAIAdminServer) DoUpgrade(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoUpgrade not implemented")
}
func (UnimplementedAIAdminServer) GetRuntimeClasses(context.Context, *emptypb.Empty) (*RuntimeClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRuntimeClasses not implemented")
}
func (UnimplementedAIAdminServer) mustEmbedUnimplementedAIAdminServer() {}

// UnsafeAIAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AIAdminServer will
// result in compilation errors.
type UnsafeAIAdminServer interface {
	mustEmbedUnimplementedAIAdminServer()
}

func RegisterAIAdminServer(s grpc.ServiceRegistrar, srv AIAdminServer) {
	s.RegisterService(&AIAdmin_ServiceDesc, srv)
}

func _AIAdmin_GetAISettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIAdminServer).GetAISettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIAdmin_GetAISettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIAdminServer).GetAISettings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIAdmin_PutAISettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AISettings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIAdminServer).PutAISettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIAdmin_PutAISettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIAdminServer).PutAISettings(ctx, req.(*AISettings))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIAdmin_DeleteAISettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIAdminServer).DeleteAISettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIAdmin_DeleteAISettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIAdminServer).DeleteAISettings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIAdmin_UpgradeAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIAdminServer).UpgradeAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIAdmin_UpgradeAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIAdminServer).UpgradeAvailable(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIAdmin_DoUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIAdminServer).DoUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIAdmin_DoUpgrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIAdminServer).DoUpgrade(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIAdmin_GetRuntimeClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIAdminServer).GetRuntimeClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIAdmin_GetRuntimeClasses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIAdminServer).GetRuntimeClasses(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AIAdmin_ServiceDesc is the grpc.ServiceDesc for AIAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AIAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AIAdmin",
	HandlerType: (*AIAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAISettings",
			Handler:    _AIAdmin_GetAISettings_Handler,
		},
		{
			MethodName: "PutAISettings",
			Handler:    _AIAdmin_PutAISettings_Handler,
		},
		{
			MethodName: "DeleteAISettings",
			Handler:    _AIAdmin_DeleteAISettings_Handler,
		},
		{
			MethodName: "UpgradeAvailable",
			Handler:    _AIAdmin_UpgradeAvailable_Handler,
		},
		{
			MethodName: "DoUpgrade",
			Handler:    _AIAdmin_DoUpgrade_Handler,
		},
		{
			MethodName: "GetRuntimeClasses",
			Handler:    _AIAdmin_GetRuntimeClasses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/aiops/apis/admin/admin.proto",
}
