package booktitlemodel

import (
	"book-store-management-backend/common"
)

type ReqCreateBook struct {
	Name        string   `json:"name" gorm:"column:name" example:"Tôi là Bêtô"`
	Description string   `json:"desc" gorm:"column:desc" example:"Tôi Là Bêtô là tác phẩm của nhà văn chuyên viết cho thanh thiếu niên Nguyễn Nhật Ánh."`
	ListedPrice float64  `json:"listedPrice" gorm:"column:listedPrice" example:"75000"`
	AuthorIDs   []string `json:"authorIds" gorm:"column:authorIds" example:"tgnna"`
	CategoryIDs []string `json:"categoryIds" gorm:"column:categoryIds" example:"dmtt,dmtruyen"`
}

func (*ReqCreateBook) TableName() string {
	return common.TableBookTitle
}
