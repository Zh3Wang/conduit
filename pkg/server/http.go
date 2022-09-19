package server

import (
	"conduit/api/interface/v1"
	"conduit/pkg/conf"
	"conduit/pkg/middleware/auth"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func NewHttpServer(c *conf.Server, cb *conf.Biz, logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.ErrorEncoder(errorEncoder),
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			validate.Validator(),
			metadata.Server(),
			selector.Server(
				auth.JWTAuthorization(cb.JwtSecret),
			).Match(NewSkipRouterMatcher()).Build(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
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

	srv := http.NewServer(opts...)

	return srv
}

func NewSkipRouterMatcher() selector.MatchFunc {
	var skip = map[string]struct{}{
		// no authentication
		interfacePb.OperationConduitInterfaceLogin:      {},
		interfacePb.OperationConduitInterfaceRegister:   {},
		interfacePb.OperationConduitInterfaceGetArticle: {},
		interfacePb.OperationConduitInterfaceGetTags:    {},
		// optional authentication
		interfacePb.OperationConduitInterfaceGetComments: {},
		//interfacePb.OperationConduitInterfaceGetProfile:   {},
		//interfacePb.OperationConduitInterfaceListArticles: {},
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skip[operation]; ok {
			return false
		}
		return true
	}
}
