package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type LoginStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type loginRepo struct {
	userStore LoginStore
}

func NewLoginRepo(userStore LoginStore) *loginRepo {
	return &loginRepo{
		userStore: userStore,
	}
}

func (repo *loginRepo) FindUserByEmail(
	ctx context.Context,
	email string) (*usermodel.User, error) {
	user, err := repo.userStore.FindUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, err
	}
	return user, nil
}
