package bookmodel

import "book-store-management-backend/common"

type BookUpdatePrice struct {
	Price *float32 `json:"price" gorm:"column:price;"`
}

func (*BookUpdatePrice) TableName() string {
	return common.TableBook
}
