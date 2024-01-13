package userbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type ChangeStatusUserRepo interface {
	UpdateStatusUsers(
		ctx context.Context,
		data *usermodel.ReqUpdateStatusUsers,
	) error
}

type changeStatusUserBiz struct {
	repo      ChangeStatusUserRepo
	requester middleware.Requester
}

func NewChangeStatusUserBiz(
	repo ChangeStatusUserRepo,
	requester middleware.Requester) *changeStatusUserBiz {
	return &changeStatusUserBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *changeStatusUserBiz) ChangeStatusUser(
	ctx context.Context,
	data *usermodel.ReqUpdateStatusUsers) error {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return usermodel.ErrUserUpdateStatusNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.UpdateStatusUsers(ctx, data); err != nil {
		return err
	}

	return nil
}
