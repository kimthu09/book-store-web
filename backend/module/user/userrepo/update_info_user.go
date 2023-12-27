package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type UpdateInfoUserStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
	UpdateInfoUser(
		ctx context.Context,
		id string,
		data *usermodel.ReqUpdateInfoUser,
	) error
}

type updateInfoUserRepo struct {
	userStore UpdateInfoUserStore
}

func NewUpdateInfoUserRepo(
	userStore UpdateInfoUserStore) *updateInfoUserRepo {
	return &updateInfoUserRepo{
		userStore: userStore,
	}
}

func (repo *updateInfoUserRepo) CheckUserStatusPermission(
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

func (repo *updateInfoUserRepo) UpdateInfoUser(
	ctx context.Context,
	userId string,
	data *usermodel.ReqUpdateInfoUser) error {
	if err := repo.userStore.UpdateInfoUser(ctx, userId, data); err != nil {
		return err
	}
	return nil
}
