package data

import (
	"context"

	userPb "conduit/api/user/v1"
	"conduit/app/interface/service/internal/biz"

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
