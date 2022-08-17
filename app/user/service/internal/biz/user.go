package biz

import (
	"conduit/pkg/encrypt"
	"context"
	"github.com/pkg/errors"
	"time"

	userPb "conduit/api/user/v1"
	usersModel "conduit/model/users_model"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Hello string
}

type UserRepo interface {
	GetProfile(context.Context, int32) (*usersModel.Users, error)
	CreateUser(context.Context, *usersModel.Users) error
	UpdateUser(context.Context, *usersModel.Users) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetProfileById(ctx context.Context, id int32) (*userPb.Profile, error) {
	r, err := uc.repo.GetProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	return &userPb.Profile{
		UserName:    r.Username,
		Bio:         r.Bio,
		Image:       r.Image,
		Following:   false,
		CreatedTime: time.Unix(r.CreatedAt, 0).Format("2006/01/02 15:04:05"),
		UpdatedTime: time.Unix(r.UpdatedAt, 0).Format("2006/01/02 15:04:05"),
	}, nil
}

func (uc *UserUsecase) Register(ctx context.Context, user *userPb.RegisterModel) (*usersModel.Users, error) {
	var userInfo = &usersModel.Users{
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		Email:        user.Email,
		Username:     user.UserName,
		Bio:          user.Bio,
		Image:        user.Image,
		PasswordHash: encrypt.Hash(user.Password),
		Following:    0,
	}
	err := uc.repo.CreateUser(ctx, userInfo)
	if err != nil {
		return nil, errors.Wrapf(err, "create user repo")
	}

	return userInfo, nil
}
