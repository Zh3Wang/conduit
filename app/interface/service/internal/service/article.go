package service

import (
	"context"

	interfacePb "conduit/api/interface/v1"
)

func (c *ConduitInterface) GetArticle(ctx context.Context, req *interfacePb.GetArticleRequest) (*interfacePb.GetArticleReply, error) {
	result, err := c.ac.GetArticleInfoBySlug(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}
	return &interfacePb.GetArticleReply{Article: result}, nil
}
