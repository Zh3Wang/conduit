// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/interface/v1/interface_service.proto

package interfacePb

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

// ConduitInterfaceClient is the client API for ConduitInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConduitInterfaceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error)
}

type conduitInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewConduitInterfaceClient(cc grpc.ClientConnInterface) ConduitInterfaceClient {
	return &conduitInterfaceClient{cc}
}

func (c *conduitInterfaceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error) {
	out := new(GetArticleReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConduitInterfaceServer is the server API for ConduitInterface service.
// All implementations must embed UnimplementedConduitInterfaceServer
// for forward compatibility
type ConduitInterfaceServer interface {
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	Login(context.Context, *LoginRequest) (*UserReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error)
	mustEmbedUnimplementedConduitInterfaceServer()
}

// UnimplementedConduitInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedConduitInterfaceServer struct {
}

func (UnimplementedConduitInterfaceServer) Register(context.Context, *RegisterRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedConduitInterfaceServer) Login(context.Context, *LoginRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedConduitInterfaceServer) GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedConduitInterfaceServer) mustEmbedUnimplementedConduitInterfaceServer() {}

// UnsafeConduitInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConduitInterfaceServer will
// result in compilation errors.
type UnsafeConduitInterfaceServer interface {
	mustEmbedUnimplementedConduitInterfaceServer()
}

func RegisterConduitInterfaceServer(s grpc.ServiceRegistrar, srv ConduitInterfaceServer) {
	s.RegisterService(&ConduitInterface_ServiceDesc, srv)
}

func _ConduitInterface_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConduitInterface_ServiceDesc is the grpc.ServiceDesc for ConduitInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConduitInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interface.v1.ConduitInterface",
	HandlerType: (*ConduitInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ConduitInterface_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ConduitInterface_Login_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ConduitInterface_GetArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interface/v1/interface_service.proto",
}
