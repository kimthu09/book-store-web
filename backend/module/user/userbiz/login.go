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
	appCtx             appctx.AppContext
	repo               LoginRepo
	accessTokenExpiry  int
	refreshTokenExpiry int
	tokenProvider      tokenprovider.Provider
	hasher             hasher.Hasher
}

func NewLoginBiz(
	appCtx appctx.AppContext,
	repo LoginRepo,
	accessTokenExpiry int,
	refreshTokenExpiry int,
	tokenProvider tokenprovider.Provider,
	hasher hasher.Hasher) *loginBiz {
	return &loginBiz{
		appCtx:             appCtx,
		repo:               repo,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
		tokenProvider:      tokenProvider,
		hasher:             hasher,
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

	accessToken, errAccessToken :=
		biz.tokenProvider.Generate(payload, biz.accessTokenExpiry)
	if errAccessToken != nil {
		return nil, common.ErrInternal(errAccessToken)
	}

	refreshToken, errRefreshToken := biz.tokenProvider.Generate(payload, biz.refreshTokenExpiry)
	if errRefreshToken != nil {
		return nil, common.ErrInternal(errRefreshToken)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
