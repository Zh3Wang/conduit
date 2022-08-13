package data

import (
	articlePb "conduit/api/article/v1"
	userPb "conduit/api/user/v1"
	"conduit/pkg/client"
	"conduit/pkg/conf"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientV3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRegistrar, NewDiscovery,
	client.NewUserServiceClient, client.NewArticleServiceClient, NewArticleRepo, NewUserRepo)

// Data .
type Data struct {
	log *log.Helper
	uc  userPb.UsersClient
	ac  articlePb.ArticleClient
}

// NewData .
func NewData(logger log.Logger, uc userPb.UsersClient, ac articlePb.ArticleClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	l := log.NewHelper(log.With(logger, "module", "conduit-interface-service/data"))

	data := Data{
		log: l,
		uc:  uc,
		ac:  ac,
	}
	return &data, cleanup, nil
}

// NewRegistrar 注册服务 etcd
func NewRegistrar(c *conf.Data, logger log.Logger) registry.Registrar {
	l := log.NewHelper(log.With(logger, "module", "conduit-interface-service/data/registrar"))
	cc, err := clientV3.New(clientV3.Config{
		Endpoints: c.Etcd.Addr,
	})
	if err != nil {
		l.Fatalf("failed connecting etcd: %s", err.Error())
	}
	r := etcd.New(cc, etcd.Namespace("conduit"))
	return r
}

// NewDiscovery 服务发现
func NewDiscovery(c *conf.Data, logger log.Logger) registry.Discovery {
	l := log.NewHelper(log.With(logger, "module", "conduit-interface-service/data/discovery"))
	cc, err := clientV3.New(clientV3.Config{
		Endpoints: c.Etcd.Addr,
	})
	if err != nil {
		l.Fatalf("failed connecting etcd: %s", err.Error())
	}
	r := etcd.New(cc, etcd.Namespace("conduit"))
	return r
}
