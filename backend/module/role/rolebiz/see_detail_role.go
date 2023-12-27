package rolebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/role/rolemodel"
	"context"
)

type SeeDetailRoleRepo interface {
	SeeRoleDetail(
		ctx context.Context,
		roleId string) (*rolemodel.ResSeeDetailRole, error)
}

type seeDetailRoleBiz struct {
	repo      SeeDetailRoleRepo
	requester middleware.Requester
}

func NewSeeDetailRoleBiz(
	repo SeeDetailRoleRepo,
	requester middleware.Requester) *seeDetailRoleBiz {
	return &seeDetailRoleBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeDetailRoleBiz) SeeDetailRole(
	ctx context.Context,
	roleId string) (*rolemodel.ResSeeDetailRole, error) {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return nil, rolemodel.ErrRoleViewNoPermission
	}

	role, err := biz.repo.SeeRoleDetail(ctx, roleId)
	if err != nil {
		return nil, err
	}

	return role, err
}
