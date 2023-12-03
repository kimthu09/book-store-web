package usermodel

import (
	"book-store-management-backend/common"
)

type ReqCreateUser struct {
	Id       string `json:"-" gorm:"column:id;"`
	Name     string `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
	Email    string `json:"email" gorm:"column:email;" example:"a@gmail.com"`
	Password string `json:"-" gorm:"column:password;"`
	Salt     string `json:"-" gorm:"column:salt;"`
	RoleId   string `json:"roleId" gorm:"column:roleId;" example:"user"`
}

func (*ReqCreateUser) TableName() string {
	return common.TableUser
}

func (data *ReqCreateUser) Validate() error {
	if common.ValidateEmptyString(data.Name) {
		return ErrUserNameEmpty
	}
	if !common.ValidateEmail(data.Email) {
		return ErrUserEmailInvalid
	}
	if !common.ValidateNotNilId(&data.RoleId) {
		return ErrUserRoleInvalid
	}
	return nil
}
