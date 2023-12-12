package bookmodel

import "book-store-management-backend/common"

type ReqUpdateBookInfo struct {
	Id          *string  `json:"id" gorm:"column:id" example:"bookId"`
	Name        *string  `json:"name" gorm:"column:name" example:"Tôi là Bêtô"`
	Description *string  `json:"desc" gorm:"column:desc" example:"Tôi Là Bêtô là tác phẩm của nhà văn chuyên viết cho thanh thiếu niên Nguyễn Nhật Ánh."`
	Edition     *int     `json:"edition" gorm:"column:edition" example:"1"`
	Quantity    *int     `json:"quantity" gorm:"column:qty" example:"0"`
	ListedPrice *float64 `json:"listedPrice" gorm:"column:listedPrice" example:"75000"`
	SellPrice   *float64 `json:"sellPrice" gorm:"column:sellPrice" example:"80000"`
	PublisherID *string  `json:"publisherId" gorm:"column:publisherId" example:"nxbdk"`
	AuthorIDs   []string `json:"authorIds" gorm:"column:authorIds" example:"tgnna"`
	CategoryIDs []string `json:"categoryIds" gorm:"column:categoryIds" example:"dmtt,dmtruyen"`
}

func (*ReqUpdateBookInfo) TableName() string {
	return common.TableBook
}
