package biz

import "github.com/go-kratos/kratos/v2/log"

type ArticleRepo interface {
}

type ArticleUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}
