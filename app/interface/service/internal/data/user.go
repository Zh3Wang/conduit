package data

import (
	interfacePb "conduit/api/interface/v1"
	userPb "conduit/api/user/v1"
	"conduit/app/interface/service/internal/biz"
	usersModel "conduit/model/users_model"
	"context"

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
