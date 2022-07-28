package data

import (
	"context"

	articlePb "conduit/api/article/v1"
	"conduit/app/interface/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &ArticleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (a *ArticleRepo) GetArticleBySlug(ctx context.Context, slug string) (*articlePb.ArticleData, error) {
	reply, err := a.data.ac.GetArticleBySlug(ctx, &articlePb.GetArticleBySlugRequest{Slug: slug})
	if err != nil {
		return nil, err
	}
	return &articlePb.ArticleData{
		Slug:           reply.Article.Slug,
		Title:          reply.Article.Title,
		Description:    reply.Article.Description,
		Body:           reply.Article.Body,
		CreatedAt:      reply.Article.CreatedAt,
		UpdatedAt:      reply.Article.UpdatedAt,
		FavoritesCount: reply.Article.FavoritesCount,
		AuthorId:       reply.Article.AuthorId,
	}, nil
}
