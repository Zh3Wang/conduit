package service

import (
	"context"

	articlePb "conduit/api/article/v1"
)

// GetArticleBySlug returns an article
func (s *ArticleService) GetArticleBySlug(ctx context.Context, in *articlePb.GetArticleBySlugRequest) (*articlePb.GetArticleBySlugReply, error) {
	s.log.WithContext(ctx).Infof("GetArticle Received: %v", in.GetSlug())

	if in.GetSlug() == "" {
		return nil, articlePb.ErrorParamIllegal("slug is empty")
	}
	result, err := s.uc.GetArticle(ctx, in.GetSlug())
	if err != nil {
		return nil, err
	}
	data := &articlePb.GetArticleBySlugReply{
		Article: result,
	}
	return data, nil
}
