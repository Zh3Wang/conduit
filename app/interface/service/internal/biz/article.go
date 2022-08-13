package biz

import (
	"context"

	articlePb "conduit/api/article/v1"
	interfacePb "conduit/api/interface/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	GetArticleBySlug(ctx context.Context, slug string) (*articlePb.ArticleData, error)
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
	// 文章基础信息
	articles, err := a.ArticleRepo.GetArticleBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	// 作者信息
	authorInfo, err := a.UserRepo.GetAuthorProfileById(ctx, articles.AuthorId)
	if err != nil {
		return nil, err
	}

	return &interfacePb.SingleArticle{
		Slug:           articles.Slug,
		Title:          articles.Title,
		Description:    articles.Description,
		Body:           articles.Body,
		TagList:        nil,
		CreatedAt:      articles.CreatedAt,
		UpdatedAt:      articles.UpdatedAt,
		Favorited:      false,
		FavoritesCount: articles.FavoritesCount,
		Author: &interfacePb.Profile{
			Username:  authorInfo.UserName,
			Bio:       authorInfo.Bio,
			Image:     authorInfo.Image,
			Following: authorInfo.Following,
		},
	}, nil
}
