package usermodel

import (
	"book-store-management-backend/common"
)

type UserCreate struct {
	Id       string `json:"-" gorm:"column:id;"`
	Name     string `json:"name" gorm:"column:name;"`
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"-" gorm:"column:password;"`
	Salt     string `json:"-" gorm:"column:salt;"`
	RoleId   string `json:"roleId" gorm:"column:roleId;"`
}

func (*UserCreate) TableName() string {
	return common.TableUser
}

func (data *UserCreate) Validate() error {
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
