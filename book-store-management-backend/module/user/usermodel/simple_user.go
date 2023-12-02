package usermodel

import "book-store-management-backend/common"

type SimpleUser struct {
	Id   string `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (*SimpleUser) TableName() string {
	return common.TableUser
}
