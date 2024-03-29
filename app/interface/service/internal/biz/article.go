package biz

import (
	interfacePb "conduit/api/interface/v1"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	GetArticleBySlug(ctx context.Context, slug string) (*interfacePb.SingleArticle, error)
	ListArticles(ctx context.Context, tag, author, favorited string, limit, offset int64) ([]*interfacePb.SingleArticle, error)
	FeedArticles(ctx context.Context, limit, offset int64) ([]*interfacePb.SingleArticle, error)
	CreateArticle(ctx context.Context, req *interfacePb.CreateArticleRequest) (*interfacePb.SingleArticle, error)
	UpdateArticle(ctx context.Context, req *interfacePb.UpdateArticleRequest) (*interfacePb.SingleArticle, error)
	DeleteArticle(ctx context.Context, req *interfacePb.DeleteArticleRequest) error

	AddComment(ctx context.Context, slug, body string) (*interfacePb.Comment, error)
	GetComments(ctx context.Context, slug string) ([]*interfacePb.Comment, error)
	DeleteComment(ctx context.Context, slug string, commentId int64) error

	FavoriteArticle(ctx context.Context, slug string) (*interfacePb.SingleArticle, error)
	UnFavoriteArticle(ctx context.Context, slug string) (*interfacePb.SingleArticle, error)

	GetTags(ctx context.Context) ([]string, error)
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
	return a.ArticleRepo.GetArticleBySlug(ctx, slug)
}

func (a *ArticleUsecase) ListArticles(ctx context.Context, req *interfacePb.ListArticlesRequest) ([]*interfacePb.SingleArticle, error) {
	return a.ArticleRepo.ListArticles(ctx, req.Tag, req.Author, req.Favorited, req.Limit, req.Offset)
}

func (a *ArticleUsecase) FeedArticles(ctx context.Context, req *interfacePb.FeedArticlesRequest) ([]*interfacePb.SingleArticle, error) {
	return a.ArticleRepo.FeedArticles(ctx, req.Limit, req.Offset)
}

func (a *ArticleUsecase) CreateArticle(ctx context.Context, req *interfacePb.CreateArticleRequest) (*interfacePb.SingleArticle, error) {
	return a.ArticleRepo.CreateArticle(ctx, req)
}

func (a *ArticleUsecase) UpdateArticle(ctx context.Context, req *interfacePb.UpdateArticleRequest) (*interfacePb.SingleArticle, error) {
	return a.ArticleRepo.UpdateArticle(ctx, req)
}

func (a *ArticleUsecase) DeleteArticle(ctx context.Context, req *interfacePb.DeleteArticleRequest) error {
	return a.ArticleRepo.DeleteArticle(ctx, req)
}

func (a *ArticleUsecase) GetTags(ctx context.Context) ([]string, error) {
	return a.ArticleRepo.GetTags(ctx)
}

func (a *ArticleUsecase) AddComment(ctx context.Context, slug, body string) (*interfacePb.Comment, error) {
	return a.ArticleRepo.AddComment(ctx, slug, body)
}

func (a *ArticleUsecase) GetComments(ctx context.Context, slug string) ([]*interfacePb.Comment, error) {
	return a.ArticleRepo.GetComments(ctx, slug)
}

func (a *ArticleUsecase) DeleteComment(ctx context.Context, slug string, commentId int64) error {
	return a.ArticleRepo.DeleteComment(ctx, slug, commentId)
}

func (a *ArticleUsecase) FavoriteArticle(ctx context.Context, slug string) (*interfacePb.SingleArticle, error) {
	return a.ArticleRepo.FavoriteArticle(ctx, slug)
}

func (a *ArticleUsecase) UnFavoriteArticle(ctx context.Context, slug string) (*interfacePb.SingleArticle, error) {
	return a.ArticleRepo.UnFavoriteArticle(ctx, slug)
}
