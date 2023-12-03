package userbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/hasher"
	"book-store-management-backend/component/tokenprovider"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type LoginRepo interface {
	FindUserByEmail(
		ctx context.Context,
		email string,
	) (*usermodel.User, error)
}

type loginBiz struct {
	appCtx        appctx.AppContext
	repo          LoginRepo
	expiry        int
	tokenProvider tokenprovider.Provider
	hasher        hasher.Hasher
}

func NewLoginBiz(
	appCtx appctx.AppContext,
	repo LoginRepo,
	expiry int,
	tokenProvider tokenprovider.Provider,
	hasher hasher.Hasher) *loginBiz {
	return &loginBiz{
		appCtx:        appCtx,
		repo:          repo,
		expiry:        expiry,
		tokenProvider: tokenProvider,
		hasher:        hasher,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.ReqLoginUser) (*usermodel.Account, error) {
	user, err := biz.repo.FindUserByEmail(ctx, data.Email)

	if err != nil {
		return nil, usermodel.ErrUserEmailOrPasswordInvalid
	}

	passwordHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passwordHashed {
		return nil, usermodel.ErrUserEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role.Id,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
