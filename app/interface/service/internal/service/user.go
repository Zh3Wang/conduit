package service

import (
	interfacePb "conduit/api/interface/v1"
	"conduit/app/interface/service/internal/biz"
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

func (c *ConduitInterface) GetCurrentUser(ctx context.Context, req *interfacePb.GetCurrentUserRequest) (*interfacePb.UserReply, error) {
	res, err := c.uc.GetCurrentUser(ctx)
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

func (c *ConduitInterface) UpdateUser(ctx context.Context, req *interfacePb.UpdateUserRequest) (*interfacePb.UserReply, error) {
	res, err := c.uc.UpdateUser(ctx, &biz.UpdateUser{
		Email:    req.User.GetEmail(),
		Password: req.User.GetPassword(),
		UserName: req.User.GetUsername(),
		Bio:      req.User.GetBio(),
		Image:    req.User.GetImage(),
	})
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

func (c *ConduitInterface) GetProfile(ctx context.Context, req *interfacePb.GetProfileRequest) (*interfacePb.ProfileReply, error) {
	res, err := c.uc.GetProfile(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &interfacePb.ProfileReply{
		Profile: &interfacePb.Profile{
			Username:  res.UserName,
			Bio:       res.Bio,
			Image:     res.Image,
			Following: res.Following,
		},
	}, nil
}
