package biz

import (
	articlePb "conduit/api/article/v1"
	"conduit/model/articles_model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"regexp"
	"strings"
	"time"
)

type ArticleRepo interface {
	GetArticle(ctx context.Context, slug string) (*articlesModel.Articles, error)
	CreateArticle(context.Context, *articlesModel.Articles) error
	UpdateArticle(ctx context.Context, oldSlug string, article *articlesModel.Articles) error
	DeleteArticle(ctx context.Context, slug string) error
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func slugify(title string) string {
	re, _ := regexp.Compile(`[^\w]`)
	return strings.ToLower(re.ReplaceAllString(title, "-"))
}

func (uc *ArticleUsecase) GetArticle(ctx context.Context, slug string) (*articlePb.ArticleData, error) {
	articles, err := uc.repo.GetArticle(ctx, slug)
	if err != nil {
		return nil, err
	}
	return &articlePb.ArticleData{
		Slug:        articles.Slug,
		Title:       articles.Title,
		Description: articles.Description,
		Body:        articles.Body,
		CreatedAt:   articles.CreatedAt,
		UpdatedAt:   articles.UpdatedAt,
		AuthorId:    articles.AuthorID,
	}, nil
}

func (uc *ArticleUsecase) CreateArticle(ctx context.Context, req *articlePb.CreateArticleRequest) (*articlePb.ArticleData, error) {
	var g = &articlesModel.Articles{
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		Slug:        req.Title,
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
		AuthorID:    req.AuthorId,
	}
	g.Slug = slugify(req.GetTitle())
	err := uc.repo.CreateArticle(ctx, g)
	if err != nil {
		return nil, err
	}

	// todo tag

	return uc.GetArticle(ctx, g.Slug)
}

func (uc *ArticleUsecase) UpdateArticle(ctx context.Context, req *articlePb.UpdateArticleRequest) (*articlePb.ArticleData, error) {
	var g = &articlesModel.Articles{
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		Slug:        slugify(req.GetTitle()),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Body:        req.GetBody(),
	}
	err := uc.repo.UpdateArticle(ctx, req.GetSlug(), g)
	if err != nil {
		return nil, err
	}

	return uc.GetArticle(ctx, g.Slug)
}

func (uc *ArticleUsecase) DeleteArticle(ctx context.Context, req *articlePb.DeleteArticleRequest) error {
	return uc.repo.DeleteArticle(ctx, req.GetSlug())
}
