package bookmodel

import (
	"book-store-management-backend/common"
)

type ReqCreateBook struct {
	Id string `json:"id" gorm:"column:id;primaryKey" example:"bookId"`
}

func (*ReqCreateBook) TableName() string {
	return common.TableBook
}
