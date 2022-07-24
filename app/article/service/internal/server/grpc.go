package server

import (
	v1 "conduit/api/article/v1"
	"conduit/app/article/service/internal/service"
	"conduit/pkg/conf"
	"conduit/pkg/server"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, article *service.ArticleService, logger log.Logger) *grpc.Server {
	srv := server.NewGrpcServer(c)
	v1.RegisterArticleServer(srv, article)
	return srv
}
