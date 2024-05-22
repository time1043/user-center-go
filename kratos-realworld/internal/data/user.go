package data

import (
	"context"

	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// --------------------------------------------------------------------------------------------------
type userRepo struct {
	data *Data
	log  *log.Helper
}

type ProfileRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &ProfileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// --------------------------------------------------------------------------------------------------
func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	return nil
}
