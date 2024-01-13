package usermodel

import "book-store-management-backend/common"

type ReqForgetPassword struct {
	NewPassword string `json:"newPassword" gorm:"-" example:"mật khẩu mới"`
}

func (*ReqForgetPassword) TableName() string {
	return common.TableUser
}

func (data *ReqForgetPassword) Validate() error {
	if !common.ValidatePassword(&data.NewPassword) {
		return ErrUserUpdatedPassInvalid
	}
	return nil
}
