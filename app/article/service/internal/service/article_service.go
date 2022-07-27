package service

import (
	"context"

	articlePb "conduit/api/article/v1"
	"conduit/app/article/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// ArticleService is an Article service.
type ArticleService struct {
	articlePb.UnimplementedArticleServer

	uc  *biz.ArticleUsecase
	log *log.Helper
}

// NewArticleService new an Article service.
func NewArticleService(uc *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{uc: uc, log: log.NewHelper(logger)}
}

// GetArticle returns an article
func (s *ArticleService) GetArticle(ctx context.Context, in *articlePb.GetArticleRequest) (*articlePb.GetArticleReply, error) {
	s.log.WithContext(ctx).Infof("GetArticle Received: %v", in.GetSlug())

	if in.GetSlug() == "" {
		return nil, articlePb.ErrorParamIllegal("slug is empty")
	}
	result, err := s.uc.GetArticle(ctx, in.GetSlug())
	if err != nil {
		return nil, err
	}
	data := &articlePb.GetArticleReply{
		Article: result,
	}
	return data, nil
}
