package data

import (
	"time"

	"conduit/pkg/conf"
	"conduit/pkg/mysql"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	etcdClient "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewArticleRepo, NewEtcdRegistrar)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// 初始化 etcd
func NewEtcdRegistrar(c *conf.Data, logger log.Logger) registry.Registrar {
	l := log.NewHelper(log.With(logger, "module", "article-service/data/etcd"))
	client, err := etcdClient.New(etcdClient.Config{
		Endpoints: c.Etcd.Addr,
	})
	if err != nil {
		l.Fatalf("failed connecting etcd: %s", err.Error())
	}
	r := etcd.New(client, etcd.Namespace("conduit"))
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

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	l := log.NewHelper(log.With(logger, "module", "article-service/data"))

	data := Data{
		db:  db,
		log: l,
	}
	return &data, cleanup, nil
}
