package data

import (
	"context"

	"conduit/app/user/service/internal/biz"
	usersModel "conduit/model/users_model"

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

func (r *userRepo) GetProfile(ctx context.Context, id int32) (*usersModel.Users, error) {
	d := &usersModel.Users{ID: int64(id)}
	err := r.data.db.WithContext(ctx).Where("id = ?", id).First(d).Error
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *userRepo) CreateUser(ctx context.Context, g *biz.User) error {
	return nil
}

func (r *userRepo) UpdateUser(ctx context.Context, g *biz.User) error {
	return nil
}
