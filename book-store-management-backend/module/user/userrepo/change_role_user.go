package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type ChangeRoleUserStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
	UpdateRoleUser(
		ctx context.Context,
		id string,
		data *usermodel.ReqUpdateRoleUser) error
}

type changeRoleUserRepo struct {
	userStore ChangeRoleUserStore
}

func NewChangeRoleUserRepo(
	userStore ChangeRoleUserStore) *changeRoleUserRepo {
	return &changeRoleUserRepo{
		userStore: userStore,
	}
}

func (repo *changeRoleUserRepo) CheckUserStatusPermission(
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

func (repo *changeRoleUserRepo) UpdateRoleUser(
	ctx context.Context,
	userId string,
	data *usermodel.ReqUpdateRoleUser) error {
	if err := repo.userStore.UpdateRoleUser(ctx, userId, data); err != nil {
		return err
	}
	return nil
}
