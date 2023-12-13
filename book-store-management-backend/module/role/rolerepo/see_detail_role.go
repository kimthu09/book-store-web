package rolerepo

import (
	"book-store-management-backend/module/role/rolemodel"
	"context"
)

type FindRoleStore interface {
	FindRole(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*rolemodel.ResSeeDetailRole, error)
}

type seeRoleDetailRepo struct {
	roleStore FindRoleStore
}

func NewSeeRoleDetailRepo(
	roleStore FindRoleStore) *seeRoleDetailRepo {
	return &seeRoleDetailRepo{
		roleStore: roleStore,
	}
}

func (biz *seeRoleDetailRepo) SeeRoleDetail(
	ctx context.Context,
	roleId string) (*rolemodel.ResSeeDetailRole, error) {
	role, errRole := biz.roleStore.FindRole(
		ctx, map[string]interface{}{"id": roleId})
	if errRole != nil {
		return nil, errRole
	}

	var roleDetail rolemodel.ResSeeDetailRole
	roleDetail.Id = role.Id
	roleDetail.Name = role.Name

	return &roleDetail, nil
}
