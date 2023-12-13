package rolerepo

import (
	"book-store-management-backend/module/role/rolemodel"
	"context"
)

type ListRoleStore interface {
	ListRole(
		ctx context.Context,
	) ([]rolemodel.SimpleRole, error)
}

type listRoleRepo struct {
	roleStore ListRoleStore
}

func NewListRoleRepo(
	roleStore ListRoleStore) *listRoleRepo {
	return &listRoleRepo{
		roleStore: roleStore,
	}
}

func (biz *listRoleRepo) ListRole(
	ctx context.Context) ([]rolemodel.SimpleRole, error) {
	roles, errRole := biz.roleStore.ListRole(ctx)
	if errRole != nil {
		return nil, errRole
	}

	return roles, nil
}
