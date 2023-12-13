package userbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type UpdateRoleUserRepo interface {
	CheckUserStatusPermission(
		ctx context.Context,
		userId string,
	) error
	UpdateRoleUser(
		ctx context.Context,
		userId string,
		data *usermodel.ReqUpdateRoleUser) error
}

type changeRoleUserBiz struct {
	repo      UpdateRoleUserRepo
	requester middleware.Requester
}

func NewChangeRoleUserBiz(
	repo UpdateRoleUserRepo,
	requester middleware.Requester) *changeRoleUserBiz {
	return &changeRoleUserBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *changeRoleUserBiz) ChangeRoleUser(
	ctx context.Context,
	id string,
	data *usermodel.ReqUpdateRoleUser) error {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return usermodel.ErrUserUpdateRoleNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.CheckUserStatusPermission(ctx, id); err != nil {
		return err
	}

	if err := biz.repo.UpdateRoleUser(ctx, id, data); err != nil {
		return err
	}

	return nil
}
