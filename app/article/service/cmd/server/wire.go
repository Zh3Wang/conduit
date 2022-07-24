//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"conduit/app/article/service/internal/biz"
	"conduit/app/article/service/internal/data"
	"conduit/app/article/service/internal/server"
	"conduit/app/article/service/internal/service"
	"conduit/pkg/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

//go:generate go run github.com/google/wire/cmd/wire

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
