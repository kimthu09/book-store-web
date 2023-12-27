package usermodel

import "book-store-management-backend/common"

type ReqUpdateStatusUser struct {
	UserId   string `json:"userId" gorm:"column:id;"`
	IsActive *bool  `json:"isActive" gorm:"column:isActive;"`
}

func (*ReqUpdateStatusUser) TableName() string {
	return common.TableUser
}
