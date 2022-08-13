package client

import (
	"context"

	articlePb "conduit/api/article/v1"
	userPb "conduit/api/user/v1"
	"conduit/pkg/service"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewArticleServiceClient(dis registry.Discovery) articlePb.ArticleClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery://conduit/"+service.ArticleService),
		grpc.WithDiscovery(dis),
	)
	if err != nil {
		panic(err)
	}

	return articlePb.NewArticleClient(conn)
}

func NewUserServiceClient(dis registry.Discovery) userPb.UsersClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery://conduit/"+service.UserService),
		grpc.WithDiscovery(dis),
	)
	if err != nil {
		panic(err)
	}

	return userPb.NewUsersClient(conn)
}
