package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	interfacePb "conduit/api/interface/v1"
)

func (c *ConduitInterface) GetArticle(ctx context.Context, req *interfacePb.GetArticleRequest) (*interfacePb.GetArticleReply, error) {
	result, err := c.ac.GetArticleInfoBySlug(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}
	return &interfacePb.GetArticleReply{Article: result}, nil
}

func (c *ConduitInterface) ListArticle(ctx context.Context, req *interfacePb.ListArticlesRequest) (*interfacePb.MultipleArticles, error) {
	result, err := c.ac.ListArticles(ctx, req)
	if err != nil {
		return nil, err
	}
	return &interfacePb.MultipleArticles{Articles: result}, nil
}

func (c *ConduitInterface) FeedArticle(ctx context.Context, req *interfacePb.FeedArticlesRequest) (*interfacePb.MultipleArticles, error) {
	result, err := c.ac.FeedArticles(ctx, req)
	if err != nil {
		return nil, err
	}
	return &interfacePb.MultipleArticles{Articles: result}, nil
}

func (c *ConduitInterface) CreateArticle(ctx context.Context, req *interfacePb.CreateArticleRequest) (*interfacePb.GetArticleReply, error) {
	result, err := c.ac.CreateArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	return &interfacePb.GetArticleReply{Article: result}, nil
}

func (c *ConduitInterface) UpdateArticle(ctx context.Context, req *interfacePb.UpdateArticleRequest) (*interfacePb.GetArticleReply, error) {
	result, err := c.ac.UpdateArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	return &interfacePb.GetArticleReply{Article: result}, nil
}

func (c *ConduitInterface) DeleteArticle(ctx context.Context, req *interfacePb.DeleteArticleRequest) (*empty.Empty, error) {
	err := c.ac.DeleteArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *ConduitInterface) GetTags(ctx context.Context, _ *empty.Empty) (*interfacePb.GetTagsReply, error) {
	result, err := c.ac.GetTags(ctx)
	if err != nil {
		return nil, err
	}
	return &interfacePb.GetTagsReply{Tags: result}, nil
}

func (c *ConduitInterface) AddComment(ctx context.Context, req *interfacePb.AddCommentRequest) (*interfacePb.SingleCommentReply, error) {
	result, err := c.ac.AddComment(ctx, req.GetSlug(), req.Comment.GetBody())
	if err != nil {
		return nil, err
	}
	return &interfacePb.SingleCommentReply{Comment: result}, nil
}

func (c *ConduitInterface) GetComments(ctx context.Context, req *interfacePb.GetCommentsRequest) (*interfacePb.MultipleCommentsReply, error) {
	result, err := c.ac.GetComments(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}
	return &interfacePb.MultipleCommentsReply{Comments: result}, nil
}

func (c *ConduitInterface) DeleteComment(ctx context.Context, req *interfacePb.DeleteCommentRequest) (*empty.Empty, error) {
	err := c.ac.DeleteComment(ctx, req.GetSlug(), req.GetId())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *ConduitInterface) FavoriteArticle(ctx context.Context, req *interfacePb.FavoriteArticleRequest) (*interfacePb.GetArticleReply, error) {
	result, err := c.ac.FavoriteArticle(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}
	return &interfacePb.GetArticleReply{Article: result}, nil
}

func (c *ConduitInterface) UnfavoriteArticle(ctx context.Context, req *interfacePb.UnfavoriteArticleRequest) (*interfacePb.GetArticleReply, error) {
	result, err := c.ac.UnFavoriteArticle(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}
	return &interfacePb.GetArticleReply{Article: result}, nil
}
