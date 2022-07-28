package service

import (
	articlePb "conduit/api/article/v1"
	"conduit/app/article/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewArticleService)

// ArticleService is an Article service.
type ArticleService struct {
	articlePb.UnimplementedArticleServer

	uc  *biz.ArticleUsecase
	log *log.Helper
}

// NewArticleService new an Article service.
func NewArticleService(uc *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{uc: uc, log: log.NewHelper(logger)}
}
