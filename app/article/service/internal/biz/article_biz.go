package biz

import (
	"context"

	"conduit/model/articles_model"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	GetArticle(ctx context.Context, articleId int32) (*articlesModel.Articles, error)
	CreateArticle(context.Context, *articlesModel.Articles) error
	UpdateArticle(context.Context, *articlesModel.Articles) error
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUsecase) GetArticle(ctx context.Context, articleId int32) (*articlesModel.Articles, error) {
	return uc.repo.GetArticle(ctx, articleId)
}

func (uc *ArticleUsecase) Create(ctx context.Context, g *articlesModel.Articles) error {
	return uc.repo.CreateArticle(ctx, g)
}

func (uc *ArticleUsecase) Update(ctx context.Context, g *articlesModel.Articles) error {
	return uc.repo.UpdateArticle(ctx, g)
}
