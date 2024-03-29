// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/interface/v1/interface_service.proto

package interfacePb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
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
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticles, error)
	FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticles, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error)
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error)
	UnfavoriteArticle(ctx context.Context, in *UnfavoriteArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error)
	GetTags(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTagsReply, error)
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

func (c *conduitInterfaceClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/FollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/UnfollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticles, error) {
	out := new(MultipleArticles)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/ListArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticles, error) {
	out := new(MultipleArticles)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/FeedArticles", in, out, opts...)
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

func (c *conduitInterfaceClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error) {
	out := new(GetArticleReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error) {
	out := new(GetArticleReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error) {
	out := new(SingleCommentReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error) {
	out := new(MultipleCommentsReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/GetComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error) {
	out := new(GetArticleReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/FavoriteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) UnfavoriteArticle(ctx context.Context, in *UnfavoriteArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error) {
	out := new(GetArticleReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/UnfavoriteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitInterfaceClient) GetTags(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTagsReply, error) {
	out := new(GetTagsReply)
	err := c.cc.Invoke(ctx, "/interface.v1.ConduitInterface/GetTags", in, out, opts...)
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
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error)
	FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error)
	UnfollowUser(context.Context, *UnfollowUserRequest) (*ProfileReply, error)
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticles, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticles, error)
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*GetArticleReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*GetArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*empty.Empty, error)
	AddComment(context.Context, *AddCommentRequest) (*SingleCommentReply, error)
	GetComments(context.Context, *GetCommentsRequest) (*MultipleCommentsReply, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*empty.Empty, error)
	FavoriteArticle(context.Context, *FavoriteArticleRequest) (*GetArticleReply, error)
	UnfavoriteArticle(context.Context, *UnfavoriteArticleRequest) (*GetArticleReply, error)
	GetTags(context.Context, *empty.Empty) (*GetTagsReply, error)
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
func (UnimplementedConduitInterfaceServer) GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedConduitInterfaceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedConduitInterfaceServer) GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedConduitInterfaceServer) FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedConduitInterfaceServer) UnfollowUser(context.Context, *UnfollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUser not implemented")
}
func (UnimplementedConduitInterfaceServer) ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedConduitInterfaceServer) FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}
func (UnimplementedConduitInterfaceServer) GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedConduitInterfaceServer) CreateArticle(context.Context, *CreateArticleRequest) (*GetArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedConduitInterfaceServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*GetArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedConduitInterfaceServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedConduitInterfaceServer) AddComment(context.Context, *AddCommentRequest) (*SingleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedConduitInterfaceServer) GetComments(context.Context, *GetCommentsRequest) (*MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedConduitInterfaceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedConduitInterfaceServer) FavoriteArticle(context.Context, *FavoriteArticleRequest) (*GetArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteArticle not implemented")
}
func (UnimplementedConduitInterfaceServer) UnfavoriteArticle(context.Context, *UnfavoriteArticleRequest) (*GetArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfavoriteArticle not implemented")
}
func (UnimplementedConduitInterfaceServer) GetTags(context.Context, *empty.Empty) (*GetTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
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

func _ConduitInterface_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/FollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_UnfollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).UnfollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/UnfollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).UnfollowUser(ctx, req.(*UnfollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/ListArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_FeedArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).FeedArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/FeedArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).FeedArticles(ctx, req.(*FeedArticlesRequest))
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

func _ConduitInterface_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/GetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_FavoriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).FavoriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/FavoriteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).FavoriteArticle(ctx, req.(*FavoriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_UnfavoriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfavoriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).UnfavoriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/UnfavoriteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).UnfavoriteArticle(ctx, req.(*UnfavoriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConduitInterface_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitInterfaceServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.ConduitInterface/GetTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitInterfaceServer).GetTags(ctx, req.(*empty.Empty))
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
			MethodName: "GetCurrentUser",
			Handler:    _ConduitInterface_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ConduitInterface_UpdateUser_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _ConduitInterface_GetProfile_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _ConduitInterface_FollowUser_Handler,
		},
		{
			MethodName: "UnfollowUser",
			Handler:    _ConduitInterface_UnfollowUser_Handler,
		},
		{
			MethodName: "ListArticles",
			Handler:    _ConduitInterface_ListArticles_Handler,
		},
		{
			MethodName: "FeedArticles",
			Handler:    _ConduitInterface_FeedArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ConduitInterface_GetArticle_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _ConduitInterface_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _ConduitInterface_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ConduitInterface_DeleteArticle_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _ConduitInterface_AddComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _ConduitInterface_GetComments_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _ConduitInterface_DeleteComment_Handler,
		},
		{
			MethodName: "FavoriteArticle",
			Handler:    _ConduitInterface_FavoriteArticle_Handler,
		},
		{
			MethodName: "UnfavoriteArticle",
			Handler:    _ConduitInterface_UnfavoriteArticle_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _ConduitInterface_GetTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interface/v1/interface_service.proto",
}
