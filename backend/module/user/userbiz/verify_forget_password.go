package userbiz

import (
	"book-store-management-backend/component/hasher"
	"book-store-management-backend/component/tokenprovider"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type VerifyForgetPasswordRepo interface {
	GetUser(
		ctx context.Context,
		email string,
	) (*usermodel.User, error)
	UpdateUserPassword(
		ctx context.Context,
		id string,
		pass string,
	) error
}

type verifyForgetPasswordBiz struct {
	repo          VerifyForgetPasswordRepo
	hasher        hasher.Hasher
	tokenProvider tokenprovider.Provider
}

func NewVerifyForgetPasswordBiz(
	repo VerifyForgetPasswordRepo,
	hasher hasher.Hasher,
	tokenProvider tokenprovider.Provider) *verifyForgetPasswordBiz {
	return &verifyForgetPasswordBiz{
		repo:          repo,
		hasher:        hasher,
		tokenProvider: tokenProvider,
	}
}

func (biz *verifyForgetPasswordBiz) VerifyForgetPassword(
	ctx context.Context,
	forgetPasswordToken string,
	data *usermodel.ReqForgetPassword) error {
	if err := data.Validate(); err != nil {
		return err
	}

	payload, err := biz.tokenProvider.ValidateTokenForPayLoadEmail(forgetPasswordToken)
	if err != nil {
		panic(err)
	}
	if err := data.Validate(); err != nil {
		return err
	}

	user, errGetUser := biz.repo.GetUser(ctx, payload.Email)
	if errGetUser != nil {
		return errGetUser
	}
	if !user.IsActive {
		return usermodel.ErrUserInactive
	}

	newPasswordHashed := biz.hasher.Hash(data.NewPassword + user.Salt)
	if err := biz.repo.UpdateUserPassword(ctx, user.Id, newPasswordHashed); err != nil {
		return err
	}

	return nil
}
