package userbiz

import (
	"book-store-management-backend/component/hasher"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/user/usermodel"
	"context"
	"log"
)

type UpdatePasswordRepo interface {
	GetUser(
		ctx context.Context,
		userId string,
	) (*usermodel.User, error)
	UpdateUserPassword(
		ctx context.Context,
		id string,
		pass string,
	) error
}

type updatePasswordBiz struct {
	repo      UpdatePasswordRepo
	hasher    hasher.Hasher
	requester middleware.Requester
}

func NewUpdatePasswordBiz(
	repo UpdatePasswordRepo,
	hasher hasher.Hasher) *updatePasswordBiz {
	return &updatePasswordBiz{
		repo:   repo,
		hasher: hasher,
	}
}

func (biz *updatePasswordBiz) UpdatePassword(
	ctx context.Context,
	id string,
	data *usermodel.ReqUpdatePasswordUser) error {
	if err := data.Validate(); err != nil {
		return err
	}

	user, errGetUser := biz.repo.GetUser(ctx, id)
	if errGetUser != nil {
		return errGetUser
	}

	if !user.IsActive {
		return usermodel.ErrUserInactive
	}

	hashedPassword := biz.hasher.Hash(data.OldPassword + user.Salt)
	log.Print(data.OldPassword)
	log.Print(user.Salt)
	if hashedPassword != user.Password {
		return usermodel.ErrUserSenderPasswordWrong
	}

	newPasswordHashed := biz.hasher.Hash(data.NewPassword + user.Salt)
	if err := biz.repo.UpdateUserPassword(ctx, id, newPasswordHashed); err != nil {
		return err
	}

	return nil
}
