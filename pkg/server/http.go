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
		interfacePb.OperationConduitInterfaceLogin:    {},
		interfacePb.OperationConduitInterfaceRegister: {},
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skip[operation]; ok {
			return false
		}
		return true
	}
}
