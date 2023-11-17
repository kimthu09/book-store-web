package userrepo

import (
	"book-store-management-backend/module/role/rolemodel"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type CheckRoleStore interface {
	FindRole(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*rolemodel.Role, error)
}

type createUserRepo struct {
	userStore CreateUserStore
	roleStore CheckRoleStore
}

func NewCreateUserRepo(
	userStore CreateUserStore,
	roleStore CheckRoleStore) *createUserRepo {
	return &createUserRepo{
		userStore: userStore,
		roleStore: roleStore,
	}
}

func (repo *createUserRepo) CheckRoleExist(ctx context.Context, roleId string) error {
	if _, err := repo.roleStore.FindRole(
		ctx,
		map[string]interface{}{
			"id": roleId,
		},
	); err != nil {
		return err
	}

	return nil
}

func (repo *createUserRepo) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := repo.userStore.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
