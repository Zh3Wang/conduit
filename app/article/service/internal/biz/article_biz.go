package biz

import (
	articlePb "conduit/api/article/v1"
	"conduit/model/articles_model"
	tagsModel "conduit/model/tags_model"
	"conduit/pkg/mysql"
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
	BatchGetArticles(ctx context.Context, opts ...mysql.QueryOption) ([]*articlesModel.Articles, error)
	CreateArticleTags(ctx context.Context, articleId int64, tags []string) error
	GetTagsFromArticleId(ctx context.Context, articleId int64) ([]tagsModel.Tags, error)
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

	// tag
	if len(req.GetTagList()) > 0 {
		err = uc.repo.CreateArticleTags(ctx, g.ID, req.GetTagList())
		if err != nil {
			return nil, err
		}
	}

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

func (uc *ArticleUsecase) ListArticles(ctx context.Context, req *articlePb.ListArticlesRequest) ([]*articlePb.ArticleData, error) {
	var (
		queryOpts = make([]mysql.QueryOption, 0, 5)
	)
	if req.GetLimit() == 0 {
		req.Limit = 20
	}
	if req.GetOffset() == 0 {
		req.Offset = 0
	}
	if req.GetTag() != "" {
		queryOpts = append(queryOpts, mysql.WithTag(req.GetTag()))
	}
	if req.GetAuthor() != "" {
		queryOpts = append(queryOpts, mysql.WithAuthorName(req.GetAuthor()))
	}
	if req.GetFavorited() != "" {
		queryOpts = append(queryOpts, mysql.WithFavorited(req.GetFavorited()))
	}
	queryOpts = append(queryOpts, mysql.WithLimit(req.GetLimit()), mysql.WithOffset(req.GetOffset()))
	res, err := uc.repo.BatchGetArticles(ctx, queryOpts...)
	if err != nil {
		return nil, err
	}

	var nr = make([]*articlePb.ArticleData, 0, req.GetLimit())
	for _, v := range res {
		// 获取tag
		tags, err := uc.repo.GetTagsFromArticleId(ctx, v.ID)
		if err != nil {
			uc.log.WithContext(ctx).Errorf("GetTagsFromArticleId err: %s", err.Error())
		}
		var tagList []string
		if len(tags) > 0 {
			for _, vv := range tags {
				tagList = append(tagList, vv.Name)
			}
		}
		nr = append(nr, &articlePb.ArticleData{
			Slug:        v.Slug,
			Title:       v.Title,
			Description: v.Description,
			Body:        v.Body,
			TagList:     tagList,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			AuthorId:    v.AuthorID,
		})
	}

	return nr, nil
}
