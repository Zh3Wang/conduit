package data

import (
	"conduit/app/article/service/internal/biz"
	"context"
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

func (r *articleRepo) CreateArticle(ctx context.Context, g *biz.Article) error {
	return nil
}

func (r *articleRepo) UpdateArticle(ctx context.Context, g *biz.Article) error {
	return nil
}

func (r *articleRepo) GetArticle(ctx context.Context, articleId int32) (*biz.Article, error) {
	var d = biz.Article{}
	result := r.data.db.WithContext(ctx).Where("article_id = ?", articleId).First(&d)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &d, nil
}
