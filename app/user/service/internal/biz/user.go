package biz

import (
	"conduit/pkg/encrypt"
	"conduit/pkg/format"
	"conduit/pkg/middleware/auth"
	"context"
	"github.com/pkg/errors"
	"strconv"
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
	GetUser(ctx context.Context, keyword, stype string) (*usersModel.Users, error)
	CreateUser(context.Context, *usersModel.Users) error
	UpdateUser(context.Context, *usersModel.Users) error
	GetUserByEmail(ctx context.Context, email string) (*usersModel.Users, error)
	CreateFollowing(ctx context.Context, userId, followId int64) error
	DeleteFollowing(ctx context.Context, userId, followId int64) error
	IsFollowing(ctx context.Context, userId, authorId int64) (bool, error)
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
	// following?
	userId := auth.GetUserIdFromContext(ctx)
	following, err := uc.repo.IsFollowing(ctx, userId, int64(id))
	if err != nil {
		uc.log.WithContext(ctx).Errorf("IsFollowing err: %s", err.Error())
	}
	return &userPb.Profile{
		UserName:    r.Username,
		Bio:         r.Bio,
		Image:       r.Image,
		Following:   following,
		CreatedTime: format.ConvertTime(r.CreatedAt),
		UpdatedTime: format.ConvertTime(r.UpdatedAt),
	}, nil
}

func (uc *UserUsecase) GetProfileByUserName(ctx context.Context, username string) (*userPb.Profile, error) {
	r, err := uc.repo.GetUser(ctx, username, "username")
	if err != nil {
		return nil, err
	}
	// following?
	userId := auth.GetUserIdFromContext(ctx)
	following, err := uc.repo.IsFollowing(ctx, userId, r.ID)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("IsFollowing err: %s", err.Error())
	}
	return &userPb.Profile{
		UserName:    r.Username,
		Bio:         r.Bio,
		Image:       r.Image,
		Following:   following,
		CreatedTime: format.ConvertTime(r.CreatedAt),
		UpdatedTime: format.ConvertTime(r.UpdatedAt),
	}, nil
}

func (uc *UserUsecase) Register(ctx context.Context, user *userPb.RegisterModel) (*usersModel.Users, error) {
	// 判断邮箱是否存在
	u, err := uc.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if u != nil {
		// 已存在
		return nil, userPb.ErrorEmailAlreadyExist("邮箱已存在")
	}
	var userInfo = &usersModel.Users{
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		Email:        user.Email,
		Username:     user.UserName,
		Bio:          user.Bio,
		Image:        user.Image,
		PasswordHash: encrypt.Hash(user.Password),
	}
	err = uc.repo.CreateUser(ctx, userInfo)
	if err != nil {
		return nil, errors.WithMessagef(err, "create user repo")
	}

	return userInfo, nil
}

func (uc *UserUsecase) Login(ctx context.Context, user *userPb.LoginRequest) (*usersModel.Users, error) {
	d, err := uc.repo.GetUserByEmail(ctx, user.GetEmail())
	if err != nil {
		return nil, err
	}

	pass := encrypt.Verify(d.PasswordHash, user.Password)
	if pass != nil {
		return nil, userPb.ErrorUserNotFound("password invalid")
	}
	return d, nil
}

func (uc *UserUsecase) GetUser(ctx context.Context, keyword, stype string) (*usersModel.Users, error) {
	d, err := uc.repo.GetUser(ctx, keyword, stype)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *userPb.UpdateUserRequest) (*usersModel.Users, error) {
	var g = &usersModel.Users{
		ID:       user.UserId,
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
	}
	if len(user.Password) > 0 {
		g.PasswordHash = encrypt.Hash(user.Password)
	}
	err := uc.repo.UpdateUser(ctx, g)
	if err != nil {
		return nil, err
	}

	//查询更新后的用户信息
	gg, err := uc.repo.GetUser(ctx, strconv.Itoa(int(user.UserId)), "id")
	if err != nil {
		return nil, err
	}
	return gg, nil
}

func (uc *UserUsecase) FollowUser(ctx context.Context, followName string) (*usersModel.Users, error) {
	// 被关注用户的ID
	followUserInfo, err := uc.repo.GetUser(ctx, followName, "username")
	if err != nil {
		return nil, err
	}
	// 当前用户的ID
	userId := auth.GetUserIdFromContext(ctx)
	// 建立follow关系
	err = uc.repo.CreateFollowing(ctx, userId, followUserInfo.ID)
	if err != nil {
		return nil, err
	}
	// get profile
	r, err := uc.repo.GetProfile(ctx, int32(followUserInfo.ID))
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (uc *UserUsecase) UnFollowUser(ctx context.Context, followName string) (*usersModel.Users, error) {
	// 被关注用户的ID
	followUserInfo, err := uc.repo.GetUser(ctx, followName, "username")
	if err != nil {
		return nil, err
	}
	// 当前用户的ID
	userId := auth.GetUserIdFromContext(ctx)

	// 删除follow关系
	err = uc.repo.DeleteFollowing(ctx, userId, followUserInfo.ID)
	if err != nil {
		return nil, err
	}
	// get profile
	r, err := uc.repo.GetProfile(ctx, int32(followUserInfo.ID))
	if err != nil {
		return nil, err
	}
	return r, nil
}
