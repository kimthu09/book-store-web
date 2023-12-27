package usermodel

import "book-store-management-backend/common"

type ReqLoginUser struct {
	Email    string `json:"email" gorm:"column:email;" example:"admin@gmail.com"`
	Password string `json:"password" gorm:"-" example:"app123"`
}

func (*ReqLoginUser) TableName() string {
	return common.TableUser
}
