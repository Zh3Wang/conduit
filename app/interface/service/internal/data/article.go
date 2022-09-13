package data

import (
	"conduit/pkg/format"
	"context"

	"conduit/api/article/v1"
	"conduit/api/interface/v1"
	"conduit/api/user/v1"
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

func (a *ArticleRepo) GetArticleBySlug(ctx context.Context, slug string) (*interfacePb.SingleArticle, error) {
	reply, err := a.data.ac.GetArticleBySlug(ctx, &articlePb.GetArticleBySlugRequest{Slug: slug})
	if err != nil {
		return nil, err
	}
	// 作者信息
	authorInfo, err := a.data.uc.GetProfileById(ctx, &userPb.GetProfileByIdRequest{
		Id: int32(reply.Article.AuthorId),
	})
	if err != nil {
		return nil, err
	}

	// todo 点赞数量

	// todo 是否点赞

	return &interfacePb.SingleArticle{
		Slug:           reply.Article.Slug,
		Title:          reply.Article.Title,
		Description:    reply.Article.Description,
		Body:           reply.Article.Body,
		TagList:        nil,
		CreatedAt:      format.ConvertTime(reply.Article.CreatedAt),
		UpdatedAt:      format.ConvertTime(reply.Article.UpdatedAt),
		Favorited:      false,
		FavoritesCount: 0,
		Author: &interfacePb.Profile{
			Username:  authorInfo.Profile.UserName,
			Bio:       authorInfo.Profile.Bio,
			Image:     authorInfo.Profile.Image,
			Following: false,
		},
	}, nil
}
