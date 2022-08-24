// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"conduit/app/interface/service/internal/biz"
	"conduit/app/interface/service/internal/data"
	"conduit/app/interface/service/internal/server"
	"conduit/app/interface/service/internal/service"
	"conduit/pkg/client"
	"conduit/pkg/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, confBiz *conf.Biz, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(confData, logger)
	usersClient := client.NewUserServiceClient(discovery)
	articleClient := client.NewArticleServiceClient(discovery)
	dataData, cleanup, err := data.NewData(logger, usersClient, articleClient)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(confBiz, userRepo, logger)
	articleRepo := data.NewArticleRepo(dataData, logger)
	articleUsecase := biz.NewArticleUsecase(articleRepo, userRepo, logger)
	conduitInterface := service.NewConduitInterface(userUsecase, articleUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, confBiz, conduitInterface, logger)
	registrar := data.NewRegistrar(confData, logger)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
