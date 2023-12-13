package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type UpdatePasswordStore interface {
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

type updatePasswordRepo struct {
	userStore ResetPasswordStore
}

func NewUpdatePasswordRepo(
	userStore ResetPasswordStore) *updatePasswordRepo {
	return &updatePasswordRepo{
		userStore: userStore,
	}
}

func (repo *updatePasswordRepo) GetUser(
	ctx context.Context,
	userId string) (*usermodel.User, error) {
	currentUser, err := repo.userStore.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, err
	}
	return currentUser, nil
}

func (repo *updatePasswordRepo) UpdateUserPassword(
	ctx context.Context, id string, pass string) error {
	if err := repo.userStore.UpdatePasswordUser(ctx, id, pass); err != nil {
		return err
	}

	return nil
}
