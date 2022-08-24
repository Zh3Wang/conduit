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

func (c *ConduitInterface) Login(ctx context.Context, req *interfacePb.LoginRequest) (*interfacePb.UserReply, error) {
	res, err := c.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &interfacePb.UserReply{User: &interfacePb.User{
		Email:    res.Email,
		Token:    res.Token,
		Username: res.Username,
		Bio:      res.Bio,
		Image:    res.Image,
	}}, nil
}
