package data

import (
	"context"

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

func (r *articleRepo) CreateArticle(ctx context.Context, article *articlesModel.Articles) error {
	return r.data.db.WithContext(ctx).Create(article).Error
}

func (r *articleRepo) UpdateArticle(ctx context.Context, oldSlug string, article *articlesModel.Articles) error {
	return r.data.db.WithContext(ctx).Where("slug = ?", oldSlug).Updates(article).Error
}

func (r *articleRepo) DeleteArticle(ctx context.Context, slug string) error {
	return r.data.db.WithContext(ctx).Where("slug = ?", slug).Delete(&articlesModel.Articles{}).Error
}
