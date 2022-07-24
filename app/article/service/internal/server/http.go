package server

import (
	v1 "conduit/api/article/v1"
	"conduit/app/article/service/internal/service"
	"conduit/pkg/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, article *service.ArticleService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	//opts = append(opts, http.ErrorEncoder(conduitError.ErrorEncoder))
	srv := http.NewServer(opts...)
	v1.RegisterArticleHTTPServer(srv, article)
	return srv
}
