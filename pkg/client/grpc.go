package client

import (
	articlePb "conduit/api/article/v1"
	userPb "conduit/api/user/v1"
	"conduit/pkg/middleware/errors"
	"conduit/pkg/service"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/metadata"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewArticleServiceClient(dis registry.Discovery) articlePb.ArticleClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		NewClientOption(dis, service.ArticleService)...,
	)
	if err != nil {
		panic(err)
	}

	return articlePb.NewArticleClient(conn)
}

func NewUserServiceClient(dis registry.Discovery) userPb.UsersClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		NewClientOption(dis, service.UserService)...,
	)
	if err != nil {
		panic(err)
	}

	return userPb.NewUsersClient(conn)
}

func NewClientOption(dis registry.Discovery, serviceName string) []grpc.ClientOption {
	return []grpc.ClientOption{
		grpc.WithEndpoint("discovery://conduit/" + serviceName),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			errors.ClientConvertError(),
			metadata.Client(),
		),
	}
}
