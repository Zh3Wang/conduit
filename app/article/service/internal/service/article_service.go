package service

import (
	articlePb "conduit/api/article/v1"
	"conduit/app/article/service/internal/biz"
	"context"
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
	s.log.WithContext(ctx).Infof("GetArticle Received: %v", in.GetArticleId())

	if in.GetArticleId() == 0 {
		return nil, articlePb.ErrorParamIllegal("article id is empty")
	}
	result, err := s.uc.GetArticle(ctx, in.GetArticleId())
	if err != nil {
		return nil, err
	}
	data := &articlePb.GetArticleReply{
		Message: &articlePb.ArticleData{
			ArticleId:  result.ArticleId,
			Title:      result.ArticleTitle,
			UserId:     result.UserId,
			CreateTime: result.CreateTime,
			UpdateTime: result.UpdateTime,
			StarNum:    result.StarNum,
			ReadNum:    result.ReadNum,
		},
	}

	return data, nil
}
