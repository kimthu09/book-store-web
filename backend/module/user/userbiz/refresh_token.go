package userbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/tokenprovider"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type refreshTokenBiz struct {
	expiry        int
	tokenProvider tokenprovider.Provider
}

func NewRefreshTokenBiz(
	expiry int,
	tokenProvider tokenprovider.Provider) *refreshTokenBiz {
	return &refreshTokenBiz{
		expiry:        expiry,
		tokenProvider: tokenProvider,
	}
}

func (biz *refreshTokenBiz) RefreshToken(
	ctx context.Context, refreshToken *usermodel.ReqRefreshToken) (*usermodel.AccountWithoutRefresh, error) {
	payload, err := biz.tokenProvider.Validate(refreshToken.RefreshToken)
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
