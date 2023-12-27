package userbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/hasher"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type ResetPasswordRepo interface {
	GetUser(
		ctx context.Context,
		id string,
	) (*usermodel.User, error)
	CheckUserStatusPermission(
		ctx context.Context,
		userId string,
	) error
	UpdateUserPassword(
		ctx context.Context,
		id string,
		pass string,
	) error
}

type resetPasswordBiz struct {
	repo      ResetPasswordRepo
	hasher    hasher.Hasher
	requester middleware.Requester
}

func NewResetPasswordBiz(
	repo ResetPasswordRepo,
	hasher hasher.Hasher,
	requester middleware.Requester) *resetPasswordBiz {
	return &resetPasswordBiz{
		repo:      repo,
		hasher:    hasher,
		requester: requester,
	}
}

func (biz *resetPasswordBiz) ResetPassword(
	ctx context.Context,
	id string,
	data *usermodel.ReqResetPasswordUser) error {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return usermodel.ErrUserResetPasswordNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	user, errGetUser := biz.repo.GetUser(ctx, biz.requester.GetUserId())
	if errGetUser != nil {
		return errGetUser
	}
	passwordHashed := biz.hasher.Hash(data.UserSenderPass + user.Salt)
	if user.Password != passwordHashed {
		return usermodel.ErrUserSenderPasswordWrong
	}

	if err := biz.repo.CheckUserStatusPermission(ctx, id); err != nil {
		return err
	}

	resetUser, errGetResetUser := biz.repo.GetUser(ctx, id)
	if errGetResetUser != nil {
		return errGetResetUser
	}
	newPasswordHashed := biz.hasher.Hash(common.DefaultPass + resetUser.Salt)

	if err := biz.repo.UpdateUserPassword(ctx, id, newPasswordHashed); err != nil {
		return err
	}

	return nil
}
