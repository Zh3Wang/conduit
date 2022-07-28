package biz

import (
	"context"

	userPb "conduit/api/user/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetAuthorProfileById(ctx context.Context, authorId int32) (*userPb.Profile, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}
