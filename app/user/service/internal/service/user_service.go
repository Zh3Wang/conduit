package service

import (
	"conduit/api/user/v1"
	v1 "conduit/api/user/v1"
	"conduit/app/user/service/internal/biz"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUsersServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

// GetProfileById 获取作者信息 by id
func (u *UserService) GetProfileById(ctx context.Context, req *userPb.GetProfileByIdRequest) (*userPb.GetProfileReply, error) {
	u.log.WithContext(ctx).Infof("GetProfileById Received ---- {%+v}", req)
	if req.GetId() <= 0 {
		return nil, errors.New(422, "PARAM_ILLEGAL", "非法参数")
	}
	r, err := u.uc.GetProfileById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &userPb.GetProfileReply{
		Profile: r,
	}, nil
}

func (u *UserService) Register(ctx context.Context, req *userPb.RegisterRequest) (*userPb.UserReply, error) {
	res, err := u.uc.Register(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &userPb.UserReply{
		User: &userPb.User{
			UserName:    res.Username,
			Bio:         res.Bio,
			Image:       res.Image,
			Email:       res.Email,
			CreatedTime: time.Unix(res.CreatedAt, 0).Format("2006/01/02 15:04:05"),
			UpdatedTime: time.Unix(res.UpdatedAt, 0).Format("2006/01/02 15:04:05"),
		},
	}, nil
}
