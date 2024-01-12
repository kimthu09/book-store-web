package userrepo

import (
	"context"
)

type mailForgetPasswordRepo struct {
	userStore FindUserStore
}

func NewMailForgetPasswordRepo(
	userStore FindUserStore) *mailForgetPasswordRepo {
	return &mailForgetPasswordRepo{
		userStore: userStore,
	}
}

func (biz *mailForgetPasswordRepo) GetUserId(
	ctx context.Context,
	email string) (*string, error) {
	user, errUser := biz.userStore.FindUser(
		ctx, map[string]interface{}{"email": email})
	if errUser != nil {
		return nil, errUser
	}

	return &user.Id, nil
}
