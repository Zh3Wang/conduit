package data

import (
	interfacePb "conduit/api/interface/v1"
	userPb "conduit/api/user/v1"
	"conduit/app/interface/service/internal/biz"
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

func (u *userRepo) CreateUser(ctx context.Context, info *interfacePb.RegisterUserModel) (*interfacePb.User, error) {
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
	return &interfacePb.User{
		Email:    reply.User.Email,
		Username: reply.User.UserName,
		Bio:      reply.User.Bio,
		Image:    reply.User.Image,
	}, nil
}
