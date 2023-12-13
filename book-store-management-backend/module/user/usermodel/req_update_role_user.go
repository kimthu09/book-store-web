package usermodel

import "book-store-management-backend/common"

type ReqUpdateRoleUser struct {
	RoleId string `json:"roleId" gorm:"column:roleId;" example:"role id"`
}

func (*ReqUpdateRoleUser) TableName() string {
	return common.TableUser
}

func (data *ReqUpdateRoleUser) Validate() error {
	if !common.ValidateNotNilId(&data.RoleId) {
		return ErrUserRoleInvalid
	}
	return nil
}
