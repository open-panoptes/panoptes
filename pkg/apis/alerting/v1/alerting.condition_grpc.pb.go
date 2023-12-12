// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/pkg/apis/alerting/v1/alerting.condition.proto

package v1

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
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
	AlertConditions_ListAlertConditionGroups_FullMethodName      = "/alerting.AlertConditions/ListAlertConditionGroups"
	AlertConditions_CreateAlertCondition_FullMethodName          = "/alerting.AlertConditions/CreateAlertCondition"
	AlertConditions_GetAlertCondition_FullMethodName             = "/alerting.AlertConditions/GetAlertCondition"
	AlertConditions_ListAlertConditions_FullMethodName           = "/alerting.AlertConditions/ListAlertConditions"
	AlertConditions_UpdateAlertCondition_FullMethodName          = "/alerting.AlertConditions/UpdateAlertCondition"
	AlertConditions_ListAlertConditionChoices_FullMethodName     = "/alerting.AlertConditions/ListAlertConditionChoices"
	AlertConditions_DeleteAlertCondition_FullMethodName          = "/alerting.AlertConditions/DeleteAlertCondition"
	AlertConditions_AlertConditionStatus_FullMethodName          = "/alerting.AlertConditions/AlertConditionStatus"
	AlertConditions_ListAlertConditionsWithStatus_FullMethodName = "/alerting.AlertConditions/ListAlertConditionsWithStatus"
	AlertConditions_CloneTo_FullMethodName                       = "/alerting.AlertConditions/CloneTo"
	AlertConditions_ActivateSilence_FullMethodName               = "/alerting.AlertConditions/ActivateSilence"
	AlertConditions_DeactivateSilence_FullMethodName             = "/alerting.AlertConditions/DeactivateSilence"
	AlertConditions_Timeline_FullMethodName                      = "/alerting.AlertConditions/Timeline"
)

// AlertConditionsClient is the client API for AlertConditions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertConditionsClient interface {
	ListAlertConditionGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.ReferenceList, error)
	CreateAlertCondition(ctx context.Context, in *AlertCondition, opts ...grpc.CallOption) (*ConditionReference, error)
	GetAlertCondition(ctx context.Context, in *ConditionReference, opts ...grpc.CallOption) (*AlertCondition, error)
	ListAlertConditions(ctx context.Context, in *ListAlertConditionRequest, opts ...grpc.CallOption) (*AlertConditionList, error)
	UpdateAlertCondition(ctx context.Context, in *UpdateAlertConditionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListAlertConditionChoices(ctx context.Context, in *AlertDetailChoicesRequest, opts ...grpc.CallOption) (*ListAlertTypeDetails, error)
	DeleteAlertCondition(ctx context.Context, in *ConditionReference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AlertConditionStatus(ctx context.Context, in *ConditionReference, opts ...grpc.CallOption) (*AlertStatusResponse, error)
	ListAlertConditionsWithStatus(ctx context.Context, in *ListStatusRequest, opts ...grpc.CallOption) (*ListStatusResponse, error)
	CloneTo(ctx context.Context, in *CloneToRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// can only active silence when alert is in firing state (limitation of
	// alertmanager)
	ActivateSilence(ctx context.Context, in *SilenceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// id corresponds to conditionId
	DeactivateSilence(ctx context.Context, in *ConditionReference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Timeline(ctx context.Context, in *TimelineRequest, opts ...grpc.CallOption) (*TimelineResponse, error)
}

type alertConditionsClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertConditionsClient(cc grpc.ClientConnInterface) AlertConditionsClient {
	return &alertConditionsClient{cc}
}

func (c *alertConditionsClient) ListAlertConditionGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.ReferenceList, error) {
	out := new(v1.ReferenceList)
	err := c.cc.Invoke(ctx, AlertConditions_ListAlertConditionGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) CreateAlertCondition(ctx context.Context, in *AlertCondition, opts ...grpc.CallOption) (*ConditionReference, error) {
	out := new(ConditionReference)
	err := c.cc.Invoke(ctx, AlertConditions_CreateAlertCondition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) GetAlertCondition(ctx context.Context, in *ConditionReference, opts ...grpc.CallOption) (*AlertCondition, error) {
	out := new(AlertCondition)
	err := c.cc.Invoke(ctx, AlertConditions_GetAlertCondition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) ListAlertConditions(ctx context.Context, in *ListAlertConditionRequest, opts ...grpc.CallOption) (*AlertConditionList, error) {
	out := new(AlertConditionList)
	err := c.cc.Invoke(ctx, AlertConditions_ListAlertConditions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) UpdateAlertCondition(ctx context.Context, in *UpdateAlertConditionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AlertConditions_UpdateAlertCondition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) ListAlertConditionChoices(ctx context.Context, in *AlertDetailChoicesRequest, opts ...grpc.CallOption) (*ListAlertTypeDetails, error) {
	out := new(ListAlertTypeDetails)
	err := c.cc.Invoke(ctx, AlertConditions_ListAlertConditionChoices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) DeleteAlertCondition(ctx context.Context, in *ConditionReference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AlertConditions_DeleteAlertCondition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) AlertConditionStatus(ctx context.Context, in *ConditionReference, opts ...grpc.CallOption) (*AlertStatusResponse, error) {
	out := new(AlertStatusResponse)
	err := c.cc.Invoke(ctx, AlertConditions_AlertConditionStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) ListAlertConditionsWithStatus(ctx context.Context, in *ListStatusRequest, opts ...grpc.CallOption) (*ListStatusResponse, error) {
	out := new(ListStatusResponse)
	err := c.cc.Invoke(ctx, AlertConditions_ListAlertConditionsWithStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) CloneTo(ctx context.Context, in *CloneToRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AlertConditions_CloneTo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) ActivateSilence(ctx context.Context, in *SilenceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AlertConditions_ActivateSilence_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) DeactivateSilence(ctx context.Context, in *ConditionReference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AlertConditions_DeactivateSilence_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertConditionsClient) Timeline(ctx context.Context, in *TimelineRequest, opts ...grpc.CallOption) (*TimelineResponse, error) {
	out := new(TimelineResponse)
	err := c.cc.Invoke(ctx, AlertConditions_Timeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertConditionsServer is the server API for AlertConditions service.
// All implementations should embed UnimplementedAlertConditionsServer
// for forward compatibility
type AlertConditionsServer interface {
	ListAlertConditionGroups(context.Context, *emptypb.Empty) (*v1.ReferenceList, error)
	CreateAlertCondition(context.Context, *AlertCondition) (*ConditionReference, error)
	GetAlertCondition(context.Context, *ConditionReference) (*AlertCondition, error)
	ListAlertConditions(context.Context, *ListAlertConditionRequest) (*AlertConditionList, error)
	UpdateAlertCondition(context.Context, *UpdateAlertConditionRequest) (*emptypb.Empty, error)
	ListAlertConditionChoices(context.Context, *AlertDetailChoicesRequest) (*ListAlertTypeDetails, error)
	DeleteAlertCondition(context.Context, *ConditionReference) (*emptypb.Empty, error)
	AlertConditionStatus(context.Context, *ConditionReference) (*AlertStatusResponse, error)
	ListAlertConditionsWithStatus(context.Context, *ListStatusRequest) (*ListStatusResponse, error)
	CloneTo(context.Context, *CloneToRequest) (*emptypb.Empty, error)
	// can only active silence when alert is in firing state (limitation of
	// alertmanager)
	ActivateSilence(context.Context, *SilenceRequest) (*emptypb.Empty, error)
	// id corresponds to conditionId
	DeactivateSilence(context.Context, *ConditionReference) (*emptypb.Empty, error)
	Timeline(context.Context, *TimelineRequest) (*TimelineResponse, error)
}

// UnimplementedAlertConditionsServer should be embedded to have forward compatible implementations.
type UnimplementedAlertConditionsServer struct {
}

func (UnimplementedAlertConditionsServer) ListAlertConditionGroups(context.Context, *emptypb.Empty) (*v1.ReferenceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertConditionGroups not implemented")
}
func (UnimplementedAlertConditionsServer) CreateAlertCondition(context.Context, *AlertCondition) (*ConditionReference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertCondition not implemented")
}
func (UnimplementedAlertConditionsServer) GetAlertCondition(context.Context, *ConditionReference) (*AlertCondition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertCondition not implemented")
}
func (UnimplementedAlertConditionsServer) ListAlertConditions(context.Context, *ListAlertConditionRequest) (*AlertConditionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertConditions not implemented")
}
func (UnimplementedAlertConditionsServer) UpdateAlertCondition(context.Context, *UpdateAlertConditionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlertCondition not implemented")
}
func (UnimplementedAlertConditionsServer) ListAlertConditionChoices(context.Context, *AlertDetailChoicesRequest) (*ListAlertTypeDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertConditionChoices not implemented")
}
func (UnimplementedAlertConditionsServer) DeleteAlertCondition(context.Context, *ConditionReference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertCondition not implemented")
}
func (UnimplementedAlertConditionsServer) AlertConditionStatus(context.Context, *ConditionReference) (*AlertStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlertConditionStatus not implemented")
}
func (UnimplementedAlertConditionsServer) ListAlertConditionsWithStatus(context.Context, *ListStatusRequest) (*ListStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertConditionsWithStatus not implemented")
}
func (UnimplementedAlertConditionsServer) CloneTo(context.Context, *CloneToRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneTo not implemented")
}
func (UnimplementedAlertConditionsServer) ActivateSilence(context.Context, *SilenceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateSilence not implemented")
}
func (UnimplementedAlertConditionsServer) DeactivateSilence(context.Context, *ConditionReference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateSilence not implemented")
}
func (UnimplementedAlertConditionsServer) Timeline(context.Context, *TimelineRequest) (*TimelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Timeline not implemented")
}

// UnsafeAlertConditionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertConditionsServer will
// result in compilation errors.
type UnsafeAlertConditionsServer interface {
	mustEmbedUnimplementedAlertConditionsServer()
}

func RegisterAlertConditionsServer(s grpc.ServiceRegistrar, srv AlertConditionsServer) {
	s.RegisterService(&AlertConditions_ServiceDesc, srv)
}

func _AlertConditions_ListAlertConditionGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).ListAlertConditionGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_ListAlertConditionGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).ListAlertConditionGroups(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_CreateAlertCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertCondition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).CreateAlertCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_CreateAlertCondition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).CreateAlertCondition(ctx, req.(*AlertCondition))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_GetAlertCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConditionReference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).GetAlertCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_GetAlertCondition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).GetAlertCondition(ctx, req.(*ConditionReference))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_ListAlertConditions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).ListAlertConditions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_ListAlertConditions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).ListAlertConditions(ctx, req.(*ListAlertConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_UpdateAlertCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).UpdateAlertCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_UpdateAlertCondition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).UpdateAlertCondition(ctx, req.(*UpdateAlertConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_ListAlertConditionChoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertDetailChoicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).ListAlertConditionChoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_ListAlertConditionChoices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).ListAlertConditionChoices(ctx, req.(*AlertDetailChoicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_DeleteAlertCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConditionReference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).DeleteAlertCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_DeleteAlertCondition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).DeleteAlertCondition(ctx, req.(*ConditionReference))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_AlertConditionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConditionReference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).AlertConditionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_AlertConditionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).AlertConditionStatus(ctx, req.(*ConditionReference))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_ListAlertConditionsWithStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).ListAlertConditionsWithStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_ListAlertConditionsWithStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).ListAlertConditionsWithStatus(ctx, req.(*ListStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_CloneTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).CloneTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_CloneTo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).CloneTo(ctx, req.(*CloneToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_ActivateSilence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SilenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).ActivateSilence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_ActivateSilence_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).ActivateSilence(ctx, req.(*SilenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_DeactivateSilence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConditionReference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).DeactivateSilence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_DeactivateSilence_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).DeactivateSilence(ctx, req.(*ConditionReference))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertConditions_Timeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertConditionsServer).Timeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertConditions_Timeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertConditionsServer).Timeline(ctx, req.(*TimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertConditions_ServiceDesc is the grpc.ServiceDesc for AlertConditions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertConditions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alerting.AlertConditions",
	HandlerType: (*AlertConditionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAlertConditionGroups",
			Handler:    _AlertConditions_ListAlertConditionGroups_Handler,
		},
		{
			MethodName: "CreateAlertCondition",
			Handler:    _AlertConditions_CreateAlertCondition_Handler,
		},
		{
			MethodName: "GetAlertCondition",
			Handler:    _AlertConditions_GetAlertCondition_Handler,
		},
		{
			MethodName: "ListAlertConditions",
			Handler:    _AlertConditions_ListAlertConditions_Handler,
		},
		{
			MethodName: "UpdateAlertCondition",
			Handler:    _AlertConditions_UpdateAlertCondition_Handler,
		},
		{
			MethodName: "ListAlertConditionChoices",
			Handler:    _AlertConditions_ListAlertConditionChoices_Handler,
		},
		{
			MethodName: "DeleteAlertCondition",
			Handler:    _AlertConditions_DeleteAlertCondition_Handler,
		},
		{
			MethodName: "AlertConditionStatus",
			Handler:    _AlertConditions_AlertConditionStatus_Handler,
		},
		{
			MethodName: "ListAlertConditionsWithStatus",
			Handler:    _AlertConditions_ListAlertConditionsWithStatus_Handler,
		},
		{
			MethodName: "CloneTo",
			Handler:    _AlertConditions_CloneTo_Handler,
		},
		{
			MethodName: "ActivateSilence",
			Handler:    _AlertConditions_ActivateSilence_Handler,
		},
		{
			MethodName: "DeactivateSilence",
			Handler:    _AlertConditions_DeactivateSilence_Handler,
		},
		{
			MethodName: "Timeline",
			Handler:    _AlertConditions_Timeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/apis/alerting/v1/alerting.condition.proto",
}
