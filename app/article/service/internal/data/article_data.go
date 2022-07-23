package data

import (
	"blog/app/article/service/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	result := r.data.db.WithContext(ctx).Where("article_id = ?", articleId).Take(&d)
	return &d, result.Error
}
