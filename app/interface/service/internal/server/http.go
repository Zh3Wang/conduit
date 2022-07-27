package server

import (
	v1 "conduit/api/user/v1"
	"conduit/app/interface/service/internal/service"
	"conduit/pkg/conf"
	"conduit/pkg/server"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.UserService, logger log.Logger) *http.Server {
	srv := server.NewHttpServer(c)
	v1.RegisterUserHTTPServer(srv, greeter)
	return srv
}
