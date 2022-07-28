package data

import (
	"context"

	articlePb "conduit/api/article/v1"
	userPb "conduit/api/user/v1"
	"conduit/app/article/service/internal/biz"
	"conduit/model/articles_model"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) CreateArticle(ctx context.Context, g *articlesModel.Articles) error {
	return nil
}

func (r *articleRepo) UpdateArticle(ctx context.Context, g *articlesModel.Articles) error {
	return nil
}

func (r *articleRepo) GetArticle(ctx context.Context, slug string) (*articlesModel.Articles, error) {
	var d = articlesModel.Articles{}
	result := r.data.db.WithContext(ctx).Where("slug = ?", slug).First(&d)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &d, nil
}

func (r *articleRepo) GetProfileById(ctx context.Context, id int32) (*articlePb.ArticleDataAuthorInfo, error) {
	resp, err := r.data.userService.GetProfileById(ctx, &userPb.GetProfileByIdRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return &articlePb.ArticleDataAuthorInfo{
		Username:  resp.Profile.UserName,
		Bio:       resp.Profile.Bio,
		Images:    resp.Profile.Image,
		Following: false,
	}, nil
}
