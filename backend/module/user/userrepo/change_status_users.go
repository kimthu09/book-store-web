package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type ChangeStatusUserStore interface {
	UpdateStatusUsers(
		ctx context.Context,
		id string,
		data *usermodel.ReqUpdateStatusUser) error
}

type changeStatusUserRepo struct {
	userStore ChangeStatusUserStore
}

func NewChangeStatusUserRepo(
	userStore ChangeStatusUserStore) *changeStatusUserRepo {
	return &changeStatusUserRepo{
		userStore: userStore,
	}
}

func (repo *changeStatusUserRepo) UpdateStatusUsers(
	ctx context.Context,
	data *usermodel.ReqUpdateStatusUsers) error {
	for _, v := range data.UserIds {
		updateModel := usermodel.ReqUpdateStatusUser{
			UserId:   v,
			IsActive: data.IsActive,
		}
		if err := repo.userStore.UpdateStatusUsers(ctx, v, &updateModel); err != nil {
			return err
		}
	}

	return nil
}
