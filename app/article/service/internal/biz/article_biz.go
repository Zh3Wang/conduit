package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	ArticleId    int32  `json:"article_id"`
	ArticleTitle string `json:"article_title"`
	UserId       int32  `json:"user_id"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
	StarNum      int32  `json:"star_num"`
	ReadNum      int32  `json:"read_num"`
}

func (a *Article) TableName() string {
	return "conduit.conduit_article_info"
}

type ArticleRepo interface {
	GetArticle(ctx context.Context, articleId int32) (*Article, error)
	CreateArticle(context.Context, *Article) error
	UpdateArticle(context.Context, *Article) error
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUsecase) GetArticle(ctx context.Context, articleId int32) (*Article, error) {
	return uc.repo.GetArticle(ctx, articleId)
}

func (uc *ArticleUsecase) Create(ctx context.Context, g *Article) error {
	return uc.repo.CreateArticle(ctx, g)
}

func (uc *ArticleUsecase) Update(ctx context.Context, g *Article) error {
	return uc.repo.UpdateArticle(ctx, g)
}
