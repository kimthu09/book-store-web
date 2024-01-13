package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
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

func (biz *mailForgetPasswordRepo) VerifyUser(
	ctx context.Context,
	email string) error {
	user, errUser := biz.userStore.FindUser(
		ctx, map[string]interface{}{"email": email})
	if errUser != nil {
		return usermodel.ErrUserEmailNotExist
	}

	if !user.IsActive {
		return usermodel.ErrUserInactive
	}
	return nil
}
