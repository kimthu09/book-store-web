package usermodel

import "book-store-management-backend/common"

type ReqUpdatePasswordUser struct {
	OldPassword string `json:"oldPassword" gorm:"-" example:"mật khẩu cũ"`
	NewPassword string `json:"newPassword" gorm:"-" example:"mật khẩu mới"`
}

func (*ReqUpdatePasswordUser) TableName() string {
	return common.TableUser
}

func (data *ReqUpdatePasswordUser) Validate() error {
	if !common.ValidatePassword(&data.OldPassword) {
		return ErrUserUpdatedPassInvalid
	}
	if !common.ValidatePassword(&data.NewPassword) {
		return ErrUserUpdatedPassInvalid
	}
	return nil
}
