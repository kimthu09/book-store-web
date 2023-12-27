package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type GetAllUserStore interface {
	GetAllUser(
		ctx context.Context,
		moreKeys ...string) ([]usermodel.SimpleUser, error)
}

type getAllUserRepo struct {
	store GetAllUserStore
}

func NewGetAllUserRepo(store GetAllUserStore) *getAllUserRepo {
	return &getAllUserRepo{store: store}
}

func (repo *getAllUserRepo) GetAllUser(
	ctx context.Context) ([]usermodel.SimpleUser, error) {
	result, err := repo.store.GetAllUser(
		ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
