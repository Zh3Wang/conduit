package data

import (
	"time"

	userPb "conduit/api/user/v1"
	"conduit/pkg/client"
	"conduit/pkg/conf"
	"conduit/pkg/mysql"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientV3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewArticleRepo,
	NewDiscovery,
	NewRegistrar,
	client.NewUserServiceClient,
)

// Data .
type Data struct {
	userService userPb.UserClient
	db          *gorm.DB
	log         *log.Helper
}

// NewData .
func NewData(
	logger log.Logger,
	db *gorm.DB,
	us userPb.UserClient,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	l := log.NewHelper(log.With(logger, "module", "article-service/data"))

	data := Data{
		userService: us,
		db:          db,
		log:         l,
	}
	return &data, cleanup, nil
}

// NewRegistrar 注册服务 etcd
func NewRegistrar(c *conf.Data, logger log.Logger) registry.Registrar {
	l := log.NewHelper(log.With(logger, "module", "article-service/data/registrar"))
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
	l := log.NewHelper(log.With(logger, "module", "article-service/data/discovery"))
	cc, err := clientV3.New(clientV3.Config{
		Endpoints: c.Etcd.Addr,
	})
	if err != nil {
		l.Fatalf("failed connecting etcd: %s", err.Error())
	}
	r := etcd.New(cc, etcd.Namespace("conduit"))
	return r
}

// todo 初始化 redis

// NewDB 初始化 mysql
func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	l := log.NewHelper(log.With(logger, "module", "article-service/data/gorm"))
	mysqlOption := mysql.Options{
		Dsn:                   c.Database.Source,
		MaxIdleConnections:    int(c.Database.MaxIdleConnections),
		MaxOpenConnections:    int(c.Database.MaxOpenConnections),
		MaxConnectionLifeTime: time.Duration(c.Database.MaxConnectionLifeTime) * time.Second,
	}
	db, err := mysql.New(mysqlOption)
	if err != nil {
		l.Fatalf("failed opening connection to mysql: %s", err.Error())
	}

	return db
}
