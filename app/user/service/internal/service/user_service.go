package service

import (
	"context"

	"conduit/api/user/v1"
	v1 "conduit/api/user/v1"
	"conduit/app/user/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

// GetProfileById 获取作者信息 by id
func (u *UserService) GetProfileById(context.Context, *userPb.GetProfileByIdRequest) (*userPb.GetProfileReply, error) {

	return &userPb.GetProfileReply{
		Profile: &userPb.Profile{
			UserName:  "",
			Bio:       "",
			Image:     "",
			Following: false,
		},
	}, nil
}
