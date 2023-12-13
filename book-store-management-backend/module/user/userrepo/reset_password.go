package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type ResetPasswordStore interface {
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

type resetPasswordRepo struct {
	userStore ResetPasswordStore
}

func NewResetPasswordRepo(
	userStore ResetPasswordStore) *resetPasswordRepo {
	return &resetPasswordRepo{
		userStore: userStore,
	}
}

func (repo *resetPasswordRepo) GetUser(
	ctx context.Context, id string) (*usermodel.User, error) {
	user, err := repo.userStore.FindUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *resetPasswordRepo) CheckUserStatusPermission(
	ctx context.Context,
	userId string) error {
	currentUser, err := repo.userStore.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return err
	}

	if !currentUser.IsActive {
		return usermodel.ErrUserInactive
	}
	return nil
}

func (repo *resetPasswordRepo) UpdateUserPassword(
	ctx context.Context, id string, pass string) error {
	if err := repo.userStore.UpdatePasswordUser(ctx, id, pass); err != nil {
		return err
	}

	return nil
}
