package service

import (
	"context"

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

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements user.UserServer
func (s *UserService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "conduitError" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "FUCK User " + in.GetName()}, nil
}
