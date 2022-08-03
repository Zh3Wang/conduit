package biz

import (
	"context"

	userPb "conduit/api/user/v1"
	usersModel "conduit/model/users_model"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Hello string
}

type UserRepo interface {
	GetProfile(context.Context, int32) (*usersModel.Users, error)
	CreateUser(context.Context, *User) error
	UpdateUser(context.Context, *User) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetProfileById(ctx context.Context, id int32) (*userPb.UserProfile, error) {
	r, err := uc.repo.GetProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	return &userPb.UserProfile{
		UserName: r.Username,
		Bio:      r.Bio,
		Image:    r.Image,
	}, nil
}

func (uc *UserUsecase) Create(ctx context.Context, g *User) error {
	return uc.repo.CreateUser(ctx, g)
}

func (uc *UserUsecase) Update(ctx context.Context, g *User) error {
	return uc.repo.UpdateUser(ctx, g)
}
