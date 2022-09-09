package errors

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

// ClientConvertError 客户端转换成 kratos 错误
func ClientConvertError() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			if err != nil {
				err = errors.FromError(err)
			}
			return
		}
	}
}

// ServerConvertError 转换成kratos错误类型
func ServerConvertError() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			if err != nil {
				if se := new(errors.Error); !errors.As(err, &se) {
					err = errors.InternalServer("INTERNAL_ERROR", "service internal error")
				}
			}
			return
		}
	}
}
