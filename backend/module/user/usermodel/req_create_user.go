package usermodel

import (
	"book-store-management-backend/common"
)

type ReqCreateUser struct {
	Id       string `json:"-" gorm:"column:id;"`
	Name     string `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
	Email    string `json:"email" gorm:"column:email;" example:"a@gmail.com"`
	Phone    string `json:"phone" gorm:"column:phone;" example:"0919199112"`
	Address  string `json:"address" gorm:"column:address;" example:"HCM"`
	Password string `json:"-" gorm:"column:password;"`
	Salt     string `json:"-" gorm:"column:salt;"`
	RoleId   string `json:"roleId" gorm:"column:roleId;" example:"role id"`
	ImgUrl   string `json:"img" gorm:"column:imgUrl" example:"https://picsum.photos/200"`
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
	if !common.ValidatePhone(data.Phone) {
		return ErrUserPhoneInvalid
	}
	if !common.ValidateNotNilId(&data.RoleId) {
		return ErrUserRoleInvalid
	}
	if !common.ValidateImage(&data.ImgUrl, common.DefaultImageAvatar) {
		return ErrUserImageInvalid
	}
	return nil
}
