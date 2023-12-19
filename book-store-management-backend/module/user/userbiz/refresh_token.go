package userbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/tokenprovider"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type refreshTokenBiz struct {
	appCtx        appctx.AppContext
	expiry        int
	tokenProvider tokenprovider.Provider
}

func NewRefreshTokenBiz(
	appCtx appctx.AppContext,
	expiry int,
	tokenProvider tokenprovider.Provider) *refreshTokenBiz {
	return &refreshTokenBiz{
		appCtx:        appCtx,
		expiry:        expiry,
		tokenProvider: tokenProvider,
	}
}

func (biz *refreshTokenBiz) RefreshToken(
	ctx context.Context, refreshToken string) (*usermodel.AccountWithoutRefresh, error) {
	payload, err := biz.tokenProvider.Validate(refreshToken)
	if err != nil {
		panic(err)
	}

	accessToken, err := biz.tokenProvider.Generate(*payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccountWithoutRefresh(accessToken)

	return account, nil
}
