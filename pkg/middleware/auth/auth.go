package auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
)

func JWTAuthorization() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			//if tr, ok := transport.FromServerContext(ctx); ok {
			//	// Do something on entering
			//	header := tr.RequestHeader()
			//	defer func() {
			//		// Do something on exiting
			//	}()
			//}
			return handler(ctx, req)
		}
	}
}
