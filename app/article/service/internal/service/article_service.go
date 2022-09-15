package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	articlePb "conduit/api/article/v1"
)

// GetArticleBySlug returns an article
func (s *ArticleService) GetArticleBySlug(ctx context.Context, in *articlePb.GetArticleBySlugRequest) (*articlePb.GetArticleReply, error) {
	result, err := s.uc.GetArticle(ctx, in.GetSlug())
	if err != nil {
		return nil, err
	}
	data := &articlePb.GetArticleReply{
		Article: result,
	}
	return data, nil
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *articlePb.CreateArticleRequest) (*articlePb.GetArticleReply, error) {
	article, err := s.uc.CreateArticle(ctx, req)
	if err != nil {
		return nil, err
	}

	return &articlePb.GetArticleReply{
		Article: article,
	}, nil
}

func (s *ArticleService) UpdateArticle(ctx context.Context, req *articlePb.UpdateArticleRequest) (*articlePb.GetArticleReply, error) {
	article, err := s.uc.UpdateArticle(ctx, req)
	if err != nil {
		return nil, err
	}

	return &articlePb.GetArticleReply{
		Article: article,
	}, nil
}

func (s *ArticleService) DeleteArticle(ctx context.Context, req *articlePb.DeleteArticleRequest) (*empty.Empty, error) {
	err := s.uc.DeleteArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *ArticleService) ListArticles(ctx context.Context, req *articlePb.ListArticlesRequest) (*articlePb.GetMultipleArticleReply, error) {
	reply, err := s.uc.ListArticles(ctx, req)
	if err != nil {
		return nil, err
	}
	return &articlePb.GetMultipleArticleReply{Article: reply}, nil
}
