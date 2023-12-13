package rolebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/feature/featuremodel"
	"book-store-management-backend/module/role/rolemodel"
	"context"
)

type ListFeatureByRoleRepo interface {
	ListFeatureByRole(
		ctx context.Context,
		roleId string) ([]featuremodel.ResFeatureDetail, error)
}

type listFeatureByRoleBiz struct {
	repo      ListFeatureByRoleRepo
	requester middleware.Requester
}

func NewListFeatureByRoleBiz(
	repo ListFeatureByRoleRepo,
	requester middleware.Requester) *listFeatureByRoleBiz {
	return &listFeatureByRoleBiz{repo: repo, requester: requester}
}

func (biz *listFeatureByRoleBiz) ListFeatureByRole(
	ctx context.Context,
	roleId string) ([]featuremodel.ResFeatureDetail, error) {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return nil, rolemodel.ErrRoleViewNoPermission
	}

	result, err := biz.repo.ListFeatureByRole(ctx, roleId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
