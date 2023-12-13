package usermodel

import "book-store-management-backend/common"

type ReqResetPasswordUser struct {
	UserSenderPass string `json:"userSenderPass" gorm:"-" example:"mật khẩu người gửi"`
}

func (*ReqResetPasswordUser) TableName() string {
	return common.TableUser
}

func (data *ReqResetPasswordUser) Validate() error {
	if !common.ValidatePassword(&data.UserSenderPass) {
		return ErrUserSenderPassInvalid
	}
	return nil
}
