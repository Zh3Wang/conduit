package server

import (
	v1 "conduit/api/article/v1"
	"conduit/app/article/service/internal/service"
	"conduit/pkg/conf"
	"conduit/pkg/server"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, article *service.ArticleService, logger log.Logger) *http.Server {
	srv := server.NewHttpServer(c)
	v1.RegisterArticleHTTPServer(srv, article)
	return srv
}
