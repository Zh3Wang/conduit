package data

import (
	interfacePb "conduit/api/interface/v1"
	userPb "conduit/api/user/v1"
	"conduit/app/interface/service/internal/biz"
	usersModel "conduit/model/users_model"
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) GetAuthorProfileById(ctx context.Context, authorId int32) (*userPb.Profile, error) {
	reply, err := u.data.uc.GetProfileById(ctx, &userPb.GetProfileByIdRequest{
		Id: authorId,
	})
	if err != nil {
		return nil, err
	}
	return &userPb.Profile{
		UserName:  reply.Profile.UserName,
		Bio:       reply.Profile.Bio,
		Image:     reply.Profile.Image,
		Following: reply.Profile.Following,
	}, nil
}

func (u *userRepo) GetProfileByUserName(ctx context.Context, username string) (*userPb.Profile, error) {
	reply, err := u.data.uc.GetProfileByUserName(ctx, &userPb.GetProfileByUserNameRequest{Username: username})
	if err != nil {
		return nil, err
	}
	return &userPb.Profile{
		UserName:  reply.Profile.UserName,
		Bio:       reply.Profile.Bio,
		Image:     reply.Profile.Image,
		Following: reply.Profile.Following,
	}, nil
}

func (u *userRepo) Login(ctx context.Context, email, password string) (*usersModel.Users, error) {
	d, err := u.data.uc.Login(ctx, &userPb.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, interfacePb.ErrorUserNotFound("login failed")
	}
	return &usersModel.Users{
		ID:       d.GetUser().UserId,
		Email:    d.GetUser().Email,
		Username: d.GetUser().UserName,
		Bio:      d.GetUser().Bio,
		Image:    d.GetUser().Image,
	}, nil
}

func (u *userRepo) CreateUser(ctx context.Context, info *interfacePb.RegisterUserModel) (*usersModel.Users, error) {
	reply, err := u.data.uc.Register(ctx, &userPb.RegisterRequest{
		User: &userPb.RegisterModel{
			UserName: info.Username,
			Email:    info.Email,
			Bio:      "default",
			Image:    "default",
			Password: info.Password,
		},
	})
	if err != nil {
		return nil, err
	}
	if reply == nil || reply.User == nil {
		return nil, interfacePb.ErrorContentMissing("RPC Register failed")
	}
	return &usersModel.Users{
		Email:    reply.User.Email,
		Username: reply.User.UserName,
		Bio:      reply.User.Bio,
		Image:    reply.User.Image,
		ID:       reply.User.UserId,
	}, nil
}

func (u *userRepo) GetUserById(ctx context.Context, id int64) (*usersModel.Users, error) {
	reply, err := u.data.uc.GetUser(ctx, &userPb.GetUserRequest{
		Type:    "id",
		Keyword: strconv.Itoa(int(id)),
	})
	if err != nil {
		return nil, err
	}
	if reply == nil || reply.User == nil {
		return nil, interfacePb.ErrorContentMissing("RPC GetUserById failed")
	}
	return &usersModel.Users{
		Email:    reply.User.Email,
		Username: reply.User.UserName,
		Bio:      reply.User.Bio,
		Image:    reply.User.Image,
		ID:       reply.User.UserId,
	}, nil
}

func (u *userRepo) UpdateUserInfo(ctx context.Context, userId int64, updateInfo *biz.UpdateUser) (*usersModel.Users, error) {
	reply, err := u.data.uc.UpdateUser(ctx, &userPb.UpdateUserRequest{
		Email:    updateInfo.Email,
		Username: updateInfo.UserName,
		Password: updateInfo.Password,
		Image:    updateInfo.Image,
		Bio:      updateInfo.Bio,
		UserId:   userId,
	})
	if err != nil {
		return nil, err
	}

	return &usersModel.Users{
		Email:    reply.User.Email,
		Username: reply.User.UserName,
		Bio:      reply.User.Bio,
		Image:    reply.User.Image,
		ID:       reply.User.UserId,
	}, nil
}

func (u *userRepo) FollowUser(ctx context.Context, username string) (*userPb.Profile, error) {
	// rpc 关注
	reply, err := u.data.uc.FollowUser(ctx, &userPb.FollowUserRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return reply.Profile, nil
}

func (u *userRepo) UnFollowUser(ctx context.Context, username string) (*userPb.Profile, error) {
	// rpc 取关
	reply, err := u.data.uc.UnfollowUser(ctx, &userPb.UnfollowUserRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return reply.Profile, nil
}
