package usermodel

import "book-store-management-backend/component/tokenprovider"

type AccountWithoutRefresh struct {
	AccessToken *tokenprovider.Token `json:"accessToken"`
}

func NewAccountWithoutRefresh(at *tokenprovider.Token) *AccountWithoutRefresh {
	return &AccountWithoutRefresh{
		AccessToken: at,
	}
}
