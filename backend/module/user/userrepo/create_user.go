package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, data *usermodel.ReqCreateUser) error
}

type createUserRepo struct {
	userStore CreateUserStore
}

func NewCreateUserRepo(
	userStore CreateUserStore) *createUserRepo {
	return &createUserRepo{
		userStore: userStore,
	}
}

func (repo *createUserRepo) CreateUser(ctx context.Context, data *usermodel.ReqCreateUser) error {
	if err := repo.userStore.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
