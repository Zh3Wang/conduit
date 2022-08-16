package service

import (
	interfacePb "conduit/api/interface/v1"
	"context"
)

func (c *ConduitInterface) Register(ctx context.Context, req *interfacePb.RegisterRequest) (*interfacePb.UserReply, error) {
	user, err := c.uc.Register(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &interfacePb.UserReply{
		User: user,
	}, nil
}
