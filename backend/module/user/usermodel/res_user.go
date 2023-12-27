package usermodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/role/rolemodel"
)

type ResUser struct {
	Id       string               `json:"id" gorm:"column:id;" example:"user id"`
	Name     string               `json:"name" gorm:"column:name;" example:"Nguyễn Văn B"`
	Email    string               `json:"email" gorm:"column:email;" example:"b@gmail.com"`
	Phone    string               `json:"phone" gorm:"column:phone;" example:"0919199112"`
	Address  string               `json:"address" gorm:"column:address;" example:"HCM"`
	Password string               `json:"-" gorm:"column:password;"`
	Salt     string               `json:"-" gorm:"column:salt;"`
	RoleId   string               `json:"-" gorm:"column:roleId;"`
	Role     rolemodel.SimpleRole `json:"role" gorm:"foreignkey:roleId"`
	ImgUrl   string               `json:"img" gorm:"column:imgUrl" example:"https://picsum.photos/200"`
	IsActive bool                 `json:"isActive" gorm:"column:isActive;" example:"true"`
}

func (*ResUser) TableName() string {
	return common.TableUser
}
