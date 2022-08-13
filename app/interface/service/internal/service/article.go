package service

import (
	"context"

	interfacePb "conduit/api/interface/v1"
)

func (c *ConduitInterface) GetArticle(ctx context.Context, req *interfacePb.GetArticleRequest) (*interfacePb.GetArticleReply, error) {
	if req.GetSlug() == "" {
		return nil, interfacePb.ErrorParamIllegal("slug is empty")
	}
	result, err := c.ac.GetArticleInfoBySlug(ctx, req.GetSlug())
	if err != nil {
		c.log.Errorf("GetArticle err: %s", err.Error())
		return nil, err
	}
	return &interfacePb.GetArticleReply{Article: result}, nil
}
