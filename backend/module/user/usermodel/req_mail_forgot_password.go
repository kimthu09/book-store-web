package usermodel

import "book-store-management-backend/common"

type ReqMailForgotPassword struct {
	Email string `json:"email" gorm:"column:email;"`
}

func (*ReqMailForgotPassword) TableName() string {
	return common.TableUser
}

func (data *ReqMailForgotPassword) Validate() error {
	if !common.ValidateEmail(data.Email) {
		return ErrUserEmailInvalid
	}
	return nil
}
