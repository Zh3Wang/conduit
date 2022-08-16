package biz

import (
	interfacePb "conduit/api/interface/v1"
	userPb "conduit/api/user/v1"
	"conduit/pkg/conf"
	"conduit/pkg/middleware/auth"
	"context"
	"github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetAuthorProfileById(ctx context.Context, authorId int32) (*userPb.Profile, error)
}

type UserUsecase struct {
	secret string
	repo   UserRepo
	log    *log.Helper
}

func NewUserUsecase(conf *conf.Biz, repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger), secret: conf.JwtSecret}
}

func (u *UserUsecase) Register(ctx context.Context, info *interfacePb.RegisterUserModel) (*interfacePb.User, error) {
	// 生成jwt token
	token, err := u.generateToken(info.Username, info.Email)
	if err != nil {
		return nil, errors.WithMessagef(err, "generate token err")
	}
	return &interfacePb.User{
		Email:    info.Email,
		Token:    token,
		Username: info.Username,
		Bio:      "bio",
		Image:    "image",
	}, nil
}

func (u *UserUsecase) generateToken(username, email string) (string, error) {
	return auth.GenerateJwtToken(u.secret, username, email)
}
