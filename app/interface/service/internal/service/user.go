package service

import (
	interfacePb "conduit/api/interface/v1"
	"conduit/app/interface/service/internal/biz"
	"context"
)

// Register 注册
func (c *ConduitInterface) Register(ctx context.Context, req *interfacePb.RegisterRequest) (*interfacePb.UserReply, error) {
	user, err := c.uc.Register(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &interfacePb.UserReply{
		User: user,
	}, nil
}

// Login 登陆
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

// GetCurrentUser 当前用户数据
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

// UpdateUser 更新
func (c *ConduitInterface) UpdateUser(ctx context.Context, req *interfacePb.UpdateUserRequest) (*interfacePb.UserReply, error) {
	res, err := c.uc.UpdateUser(ctx, &biz.UpdateUser{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		UserName: req.GetUsername(),
		Bio:      req.GetBio(),
		Image:    req.GetImage(),
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

// GetProfile 个人信息
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

// FollowUser 关注
func (c *ConduitInterface) FollowUser(ctx context.Context, req *interfacePb.FollowUserRequest) (*interfacePb.ProfileReply, error) {
	res, err := c.uc.FollowUser(ctx, req.GetUsername())
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

// UnfollowUser 取关
func (c *ConduitInterface) UnfollowUser(ctx context.Context, req *interfacePb.UnfollowUserRequest) (*interfacePb.ProfileReply, error) {
	res, err := c.uc.UnFollowUser(ctx, req.GetUsername())
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
