package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type VerifyForgetPasswordStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
	UpdatePasswordUser(
		ctx context.Context,
		id string,
		password string,
	) error
}

type verifyForgetPassRepo struct {
	userStore VerifyForgetPasswordStore
}

func NewVerifyForgetPasswordRepo(
	userStore ResetPasswordStore) *verifyForgetPassRepo {
	return &verifyForgetPassRepo{
		userStore: userStore,
	}
}

func (repo *verifyForgetPassRepo) GetUser(
	ctx context.Context,
	email string) (*usermodel.User, error) {
	currentUser, err := repo.userStore.FindUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, err
	}
	return currentUser, nil
}

func (repo *verifyForgetPassRepo) UpdateUserPassword(
	ctx context.Context, id string, pass string) error {
	if err := repo.userStore.UpdatePasswordUser(ctx, id, pass); err != nil {
		return err
	}

	return nil
}
