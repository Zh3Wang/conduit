package server

import (
	v1 "conduit/api/interface/v1"
	"conduit/app/interface/service/internal/service"
	"conduit/pkg/conf"
	"conduit/pkg/server"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, svc *service.ConduitInterface, logger log.Logger) *http.Server {
	srv := server.NewHttpServer(c)
	v1.RegisterConduitInterfaceHTTPServer(srv, svc)
	return srv
}
