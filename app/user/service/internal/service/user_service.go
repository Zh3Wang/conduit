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
			UserId:      res.ID,
			UserName:    res.Username,
			Bio:         res.Bio,
			Image:       res.Image,
			Email:       res.Email,
			CreatedTime: convertTime(res.CreatedAt),
			UpdatedTime: convertTime(res.UpdatedAt),
		},
	}, nil
}

func (u *UserService) Login(ctx context.Context, req *userPb.LoginRequest) (*userPb.UserReply, error) {
	res, err := u.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &userPb.UserReply{
		User: &userPb.User{
			UserId:      res.ID,
			UserName:    res.Username,
			Bio:         res.Bio,
			Image:       res.Image,
			Email:       res.Email,
			CreatedTime: convertTime(res.CreatedAt),
			UpdatedTime: convertTime(res.UpdatedAt),
		},
	}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *userPb.GetUserRequest) (*userPb.UserReply, error) {
	res, err := u.uc.GetUser(ctx, req.GetKeyword(), req.GetType())
	if err != nil {
		return nil, err
	}
	return &userPb.UserReply{
		User: &userPb.User{
			UserId:      res.ID,
			UserName:    res.Username,
			Bio:         res.Bio,
			Image:       res.Image,
			Email:       res.Email,
			CreatedTime: convertTime(res.CreatedAt),
			UpdatedTime: convertTime(res.UpdatedAt),
		},
	}, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *userPb.UpdateUserRequest) (*userPb.UserReply, error) {
	res, err := u.uc.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &userPb.UserReply{
		User: &userPb.User{
			UserName:    res.Username,
			Bio:         res.Bio,
			Image:       res.Image,
			Email:       res.Email,
			CreatedTime: convertTime(res.CreatedAt),
			UpdatedTime: convertTime(res.UpdatedAt),
			UserId:      res.ID,
		},
	}, nil
}

func convertTime(t int64) string {
	return time.Unix(t, 0).Format("2006/01/02 15:04:05")
}
