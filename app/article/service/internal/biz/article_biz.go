package biz

import (
	articlePb "conduit/api/article/v1"
	"conduit/model/articles_model"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	GetArticle(ctx context.Context, slug string) (*articlesModel.Articles, error)
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

func (uc *ArticleUsecase) GetArticle(ctx context.Context, slug string) (*articlePb.ArticleData, error) {
	articles, err := uc.repo.GetArticle(ctx, slug)
	if err != nil {
		return nil, err
	}
	fmt.Println(timestamppb.New(articles.CreatedAt))
	return &articlePb.ArticleData{
		Slug:           articles.Slug,
		Title:          articles.Title,
		Description:    articles.Description,
		Body:           articles.Body,
		CreatedAt:      convertTime(articles.CreatedAt),
		UpdatedAt:      convertTime(articles.UpdatedAt),
		FavoritesCount: int32(articles.FavoritesCount),
		AuthorId:       int32(articles.AuthorID),
	}, nil
}

func (uc *ArticleUsecase) Create(ctx context.Context, g *articlesModel.Articles) error {
	return uc.repo.CreateArticle(ctx, g)
}

func (uc *ArticleUsecase) Update(ctx context.Context, g *articlesModel.Articles) error {
	return uc.repo.UpdateArticle(ctx, g)
}

func convertTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.999Z")
}
