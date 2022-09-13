package biz

import (
	interfacePb "conduit/api/interface/v1"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	GetArticleBySlug(ctx context.Context, slug string) (*interfacePb.SingleArticle, error)
}

type ArticleUsecase struct {
	ArticleRepo ArticleRepo
	UserRepo    UserRepo
	log         *log.Helper
}

func NewArticleUsecase(ar ArticleRepo, ur UserRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{
		ArticleRepo: ar,
		UserRepo:    ur,
		log:         log.NewHelper(logger),
	}
}

func (a *ArticleUsecase) GetArticleInfoBySlug(ctx context.Context, slug string) (*interfacePb.SingleArticle, error) {
	return a.ArticleRepo.GetArticleBySlug(ctx, slug)
}
