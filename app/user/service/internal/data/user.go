package data

import (
	"conduit/api/user/v1"
	"conduit/app/user/service/internal/biz"
	followingsModel "conduit/model/followings_model"
	usersModel "conduit/model/users_model"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
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

func (r *userRepo) GetProfile(ctx context.Context, id int32) (*usersModel.Users, error) {
	d := &usersModel.Users{ID: int64(id)}
	err := r.data.db.WithContext(ctx).Where("id = ?", id).First(d).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *userRepo) GetUser(ctx context.Context, keyword, stype string) (*usersModel.Users, error) {
	var (
		d   = &usersModel.Users{}
		err error
	)
	cond := fmt.Sprintf("%s = ?", stype)
	err = r.data.db.WithContext(ctx).Where(cond, keyword).First(d).Error
	if err == gorm.ErrRecordNotFound {
		return nil, userPb.ErrorUserNotFound("用户不存在")
	}
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u *usersModel.Users) error {
	err := r.data.db.WithContext(ctx).Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) UpdateUser(ctx context.Context, g *usersModel.Users) error {
	err := r.data.db.WithContext(ctx).Model(usersModel.Users{}).Where("id = ?", g.ID).Updates(g).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*usersModel.Users, error) {
	d := &usersModel.Users{}
	err := r.data.db.WithContext(ctx).Where("email = ? ", email).First(d).Error
	if err == gorm.ErrRecordNotFound {
		return nil, userPb.ErrorUserNotFound("用户不存在")
	}
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *userRepo) CreateFollowing(ctx context.Context, userId, followId int64) error {
	d := &followingsModel.Followings{
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		UserID:      userId,
		FollowingID: followId,
	}
	err := r.data.db.WithContext(ctx).Create(d).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) DeleteFollowing(ctx context.Context, userId, followId int64) error {
	err := r.data.db.WithContext(ctx).Where("user_id = ? and following_id = ?", userId, followId).Delete(&followingsModel.Followings{}).Error
	if err != nil {
		return err
	}
	return nil
}
