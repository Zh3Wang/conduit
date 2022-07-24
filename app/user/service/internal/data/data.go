package data

import (
	"conduit/pkg/conf"
	"conduit/pkg/mysql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
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
