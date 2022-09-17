package service

import (
	"conduit/api/article/v1"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *ArticleService) AddComment(ctx context.Context, req *articlePb.AddCommentRequest) (*articlePb.GetCommentReply, error) {
	res, err := s.sc.AddComment(ctx, req)
	if err != nil {
		return nil, err
	}
	return &articlePb.GetCommentReply{Comment: res}, nil
}

func (s *ArticleService) GetComments(ctx context.Context, req *articlePb.GetCommentsRequest) (*articlePb.MultiGetCommentsReply, error) {
	res, err := s.sc.GetComments(ctx, req)
	if err != nil {
		return nil, err
	}
	return &articlePb.MultiGetCommentsReply{Comments: res}, nil
}

func (s *ArticleService) DeleteComment(ctx context.Context, req *articlePb.DeleteCommentRequest) (*empty.Empty, error) {
	err := s.sc.DeleteComment(ctx, req)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *ArticleService) FavoriteArticle(ctx context.Context, req *articlePb.FavoriteArticleRequest) (*articlePb.GetArticleReply, error) {
	res, err := s.sc.FavoriteArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	return &articlePb.GetArticleReply{Article: res}, nil
}

func (s *ArticleService) UnFavoriteArticle(ctx context.Context, req *articlePb.UnFavoriteArticleRequest) (*articlePb.GetArticleReply, error) {
	res, err := s.sc.UnFavoriteArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	return &articlePb.GetArticleReply{Article: res}, nil
}
