package biz

import (
	interfacePb "conduit/api/interface/v1"
	userPb "conduit/api/user/v1"
	usersModel "conduit/model/users_model"
	"conduit/pkg/conf"
	"conduit/pkg/middleware/auth"
	"context"
	"github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetUserById(ctx context.Context, id int64) (*usersModel.Users, error)
	GetAuthorProfileById(ctx context.Context, authorId int32) (*userPb.Profile, error)
	GetProfileByUserName(ctx context.Context, username string) (*userPb.Profile, error)
	CreateUser(ctx context.Context, info *interfacePb.RegisterUserModel) (*usersModel.Users, error)
	Login(ctx context.Context, email, password string) (*usersModel.Users, error)
	UpdateUserInfo(ctx context.Context, userId int64, updateInfo *UpdateUser) (*usersModel.Users, error)
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
	res, err := u.repo.CreateUser(ctx, info)
	if err != nil {
		return nil, errors.WithMessagef(err, "CreateUser repo err")
	}
	// 生成jwt token
	token, err := u.generateToken(res.Username, res.ID)
	if err != nil {
		return nil, errors.WithMessagef(err, "generate token err")
	}
	return &interfacePb.User{
		Email:    res.Email,
		Token:    token,
		Username: res.Username,
		Bio:      res.Bio,
		Image:    res.Image,
	}, nil
}

func (u *UserUsecase) generateToken(username string, id int64) (string, error) {
	return auth.GenerateJwtToken(u.secret, username, id)
}

func (u *UserUsecase) Login(ctx context.Context, email, password string) (*interfacePb.User, error) {
	res, err := u.repo.Login(ctx, email, password)
	if err != nil {
		return nil, err
	}
	// 生成jwt token
	token, err := u.generateToken(res.Username, res.ID)
	if err != nil {
		return nil, errors.WithMessagef(err, "generate token err")
	}
	return &interfacePb.User{
		Email:    res.Email,
		Token:    token,
		Username: res.Username,
		Bio:      res.Bio,
		Image:    res.Image,
	}, nil
}

func (u *UserUsecase) GetCurrentUser(ctx context.Context) (*interfacePb.User, error) {
	uInfo := auth.FromContext(ctx)
	if uInfo == nil {
		return nil, interfacePb.ErrorUserNotFound("GetCurrentUser FromContext failed")
	}
	res, err := u.repo.GetUserById(ctx, uInfo.UserId)
	if err != nil {
		return nil, err
	}
	return &interfacePb.User{
		Email:    res.Email,
		Username: res.Username,
		Bio:      res.Bio,
		Image:    res.Image,
	}, nil
}

func (u *UserUsecase) UpdateUser(ctx context.Context, updateInfo *UpdateUser) (*interfacePb.User, error) {
	uInfo := auth.FromContext(ctx)
	if uInfo == nil {
		return nil, interfacePb.ErrorUserNotFound("UpdateUser FromContext failed")
	}
	res, err := u.repo.UpdateUserInfo(ctx, uInfo.UserId, updateInfo)
	if err != nil {
		return nil, err
	}
	return &interfacePb.User{
		Email:    res.Email,
		Username: res.Username,
		Bio:      res.Bio,
		Image:    res.Image,
	}, nil
}

func (u *UserUsecase) GetProfile(ctx context.Context, username string) (*userPb.Profile, error) {
	res, err := u.repo.GetProfileByUserName(ctx, username)
	if err != nil {
		return nil, err
	}
	return res, nil
}
