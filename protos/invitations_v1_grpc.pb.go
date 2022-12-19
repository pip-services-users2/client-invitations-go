// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/invitations_v1.proto

package protos

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

// InvitationsClient is the client API for Invitations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvitationsClient interface {
	GetInvitations(ctx context.Context, in *InvitationPageRequest, opts ...grpc.CallOption) (*InvitationPageReply, error)
	GetInvitationById(ctx context.Context, in *InvitationIdRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error)
	CreateInvitation(ctx context.Context, in *InvitationObjectRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error)
	DeleteInvitationById(ctx context.Context, in *InvitationIdRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error)
	ActivateInvitations(ctx context.Context, in *InvitationActivateRequest, opts ...grpc.CallOption) (*InvitationListReply, error)
	ApproveInvitation(ctx context.Context, in *InvitationApproveRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error)
	DenyInvitation(ctx context.Context, in *InvitationIdRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error)
	ResendInvitation(ctx context.Context, in *InvitationIdRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error)
	NotifyInvitation(ctx context.Context, in *InvitationObjectRequest, opts ...grpc.CallOption) (*InvitationEmptyReply, error)
}

type invitationsClient struct {
	cc grpc.ClientConnInterface
}

func NewInvitationsClient(cc grpc.ClientConnInterface) InvitationsClient {
	return &invitationsClient{cc}
}

func (c *invitationsClient) GetInvitations(ctx context.Context, in *InvitationPageRequest, opts ...grpc.CallOption) (*InvitationPageReply, error) {
	out := new(InvitationPageReply)
	err := c.cc.Invoke(ctx, "/invitations_v1.Invitations/get_invitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationsClient) GetInvitationById(ctx context.Context, in *InvitationIdRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error) {
	out := new(InvitationObjectReply)
	err := c.cc.Invoke(ctx, "/invitations_v1.Invitations/get_invitation_by_id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationsClient) CreateInvitation(ctx context.Context, in *InvitationObjectRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error) {
	out := new(InvitationObjectReply)
	err := c.cc.Invoke(ctx, "/invitations_v1.Invitations/create_invitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationsClient) DeleteInvitationById(ctx context.Context, in *InvitationIdRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error) {
	out := new(InvitationObjectReply)
	err := c.cc.Invoke(ctx, "/invitations_v1.Invitations/delete_invitation_by_id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationsClient) ActivateInvitations(ctx context.Context, in *InvitationActivateRequest, opts ...grpc.CallOption) (*InvitationListReply, error) {
	out := new(InvitationListReply)
	err := c.cc.Invoke(ctx, "/invitations_v1.Invitations/activate_invitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationsClient) ApproveInvitation(ctx context.Context, in *InvitationApproveRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error) {
	out := new(InvitationObjectReply)
	err := c.cc.Invoke(ctx, "/invitations_v1.Invitations/approve_invitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationsClient) DenyInvitation(ctx context.Context, in *InvitationIdRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error) {
	out := new(InvitationObjectReply)
	err := c.cc.Invoke(ctx, "/invitations_v1.Invitations/deny_invitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationsClient) ResendInvitation(ctx context.Context, in *InvitationIdRequest, opts ...grpc.CallOption) (*InvitationObjectReply, error) {
	out := new(InvitationObjectReply)
	err := c.cc.Invoke(ctx, "/invitations_v1.Invitations/resend_invitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationsClient) NotifyInvitation(ctx context.Context, in *InvitationObjectRequest, opts ...grpc.CallOption) (*InvitationEmptyReply, error) {
	out := new(InvitationEmptyReply)
	err := c.cc.Invoke(ctx, "/invitations_v1.Invitations/notify_invitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvitationsServer is the server API for Invitations service.
// All implementations must embed UnimplementedInvitationsServer
// for forward compatibility
type InvitationsServer interface {
	GetInvitations(context.Context, *InvitationPageRequest) (*InvitationPageReply, error)
	GetInvitationById(context.Context, *InvitationIdRequest) (*InvitationObjectReply, error)
	CreateInvitation(context.Context, *InvitationObjectRequest) (*InvitationObjectReply, error)
	DeleteInvitationById(context.Context, *InvitationIdRequest) (*InvitationObjectReply, error)
	ActivateInvitations(context.Context, *InvitationActivateRequest) (*InvitationListReply, error)
	ApproveInvitation(context.Context, *InvitationApproveRequest) (*InvitationObjectReply, error)
	DenyInvitation(context.Context, *InvitationIdRequest) (*InvitationObjectReply, error)
	ResendInvitation(context.Context, *InvitationIdRequest) (*InvitationObjectReply, error)
	NotifyInvitation(context.Context, *InvitationObjectRequest) (*InvitationEmptyReply, error)
	mustEmbedUnimplementedInvitationsServer()
}

// UnimplementedInvitationsServer must be embedded to have forward compatible implementations.
type UnimplementedInvitationsServer struct {
}

func (UnimplementedInvitationsServer) GetInvitations(context.Context, *InvitationPageRequest) (*InvitationPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvitations not implemented")
}
func (UnimplementedInvitationsServer) GetInvitationById(context.Context, *InvitationIdRequest) (*InvitationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvitationById not implemented")
}
func (UnimplementedInvitationsServer) CreateInvitation(context.Context, *InvitationObjectRequest) (*InvitationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvitation not implemented")
}
func (UnimplementedInvitationsServer) DeleteInvitationById(context.Context, *InvitationIdRequest) (*InvitationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInvitationById not implemented")
}
func (UnimplementedInvitationsServer) ActivateInvitations(context.Context, *InvitationActivateRequest) (*InvitationListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateInvitations not implemented")
}
func (UnimplementedInvitationsServer) ApproveInvitation(context.Context, *InvitationApproveRequest) (*InvitationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveInvitation not implemented")
}
func (UnimplementedInvitationsServer) DenyInvitation(context.Context, *InvitationIdRequest) (*InvitationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyInvitation not implemented")
}
func (UnimplementedInvitationsServer) ResendInvitation(context.Context, *InvitationIdRequest) (*InvitationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendInvitation not implemented")
}
func (UnimplementedInvitationsServer) NotifyInvitation(context.Context, *InvitationObjectRequest) (*InvitationEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyInvitation not implemented")
}
func (UnimplementedInvitationsServer) mustEmbedUnimplementedInvitationsServer() {}

// UnsafeInvitationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvitationsServer will
// result in compilation errors.
type UnsafeInvitationsServer interface {
	mustEmbedUnimplementedInvitationsServer()
}

func RegisterInvitationsServer(s grpc.ServiceRegistrar, srv InvitationsServer) {
	s.RegisterService(&Invitations_ServiceDesc, srv)
}

func _Invitations_GetInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationsServer).GetInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invitations_v1.Invitations/get_invitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationsServer).GetInvitations(ctx, req.(*InvitationPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invitations_GetInvitationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationsServer).GetInvitationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invitations_v1.Invitations/get_invitation_by_id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationsServer).GetInvitationById(ctx, req.(*InvitationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invitations_CreateInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationsServer).CreateInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invitations_v1.Invitations/create_invitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationsServer).CreateInvitation(ctx, req.(*InvitationObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invitations_DeleteInvitationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationsServer).DeleteInvitationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invitations_v1.Invitations/delete_invitation_by_id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationsServer).DeleteInvitationById(ctx, req.(*InvitationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invitations_ActivateInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationActivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationsServer).ActivateInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invitations_v1.Invitations/activate_invitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationsServer).ActivateInvitations(ctx, req.(*InvitationActivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invitations_ApproveInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationsServer).ApproveInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invitations_v1.Invitations/approve_invitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationsServer).ApproveInvitation(ctx, req.(*InvitationApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invitations_DenyInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationsServer).DenyInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invitations_v1.Invitations/deny_invitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationsServer).DenyInvitation(ctx, req.(*InvitationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invitations_ResendInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationsServer).ResendInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invitations_v1.Invitations/resend_invitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationsServer).ResendInvitation(ctx, req.(*InvitationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invitations_NotifyInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationsServer).NotifyInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invitations_v1.Invitations/notify_invitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationsServer).NotifyInvitation(ctx, req.(*InvitationObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Invitations_ServiceDesc is the grpc.ServiceDesc for Invitations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invitations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "invitations_v1.Invitations",
	HandlerType: (*InvitationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_invitations",
			Handler:    _Invitations_GetInvitations_Handler,
		},
		{
			MethodName: "get_invitation_by_id",
			Handler:    _Invitations_GetInvitationById_Handler,
		},
		{
			MethodName: "create_invitation",
			Handler:    _Invitations_CreateInvitation_Handler,
		},
		{
			MethodName: "delete_invitation_by_id",
			Handler:    _Invitations_DeleteInvitationById_Handler,
		},
		{
			MethodName: "activate_invitations",
			Handler:    _Invitations_ActivateInvitations_Handler,
		},
		{
			MethodName: "approve_invitation",
			Handler:    _Invitations_ApproveInvitation_Handler,
		},
		{
			MethodName: "deny_invitation",
			Handler:    _Invitations_DenyInvitation_Handler,
		},
		{
			MethodName: "resend_invitation",
			Handler:    _Invitations_ResendInvitation_Handler,
		},
		{
			MethodName: "notify_invitation",
			Handler:    _Invitations_NotifyInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/invitations_v1.proto",
}
