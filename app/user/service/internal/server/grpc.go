package server

import (
	v1 "conduit/api/user/v1"
	"conduit/app/user/service/internal/service"
	"conduit/pkg/conf"
	"conduit/pkg/server"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.UserService, logger log.Logger) *grpc.Server {
	srv := server.NewGrpcServer(c, logger)
	v1.RegisterUsersServer(srv, greeter)
	return srv
}
