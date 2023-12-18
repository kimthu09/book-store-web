package booktitlemodel

import "book-store-management-backend/common"

type ReqUpdateBookInfo struct {
	Name        *string   `json:"name" gorm:"column:name" example:"Tôi là Bêtô"`
	Description *string   `json:"desc" gorm:"column:desc" example:"Tôi Là Bêtô là tác phẩm của nhà văn chuyên viết cho thanh thiếu niên Nguyễn Nhật Ánh."`
	AuthorIDs   *[]string `json:"authorIds" gorm:"column:authorIds" example:"tgnna"`
	CategoryIDs *[]string `json:"categoryIds" gorm:"column:categoryIds" example:"dmtt,dmtruyen"`
}

func (*ReqUpdateBookInfo) TableName() string {
	return common.TableBookTitle
}
