package bookmodel

import "book-store-management-backend/common"

type BookUpdateImportPrice struct {
	ImportPrice *int `json:"importPrice" gorm:"column:importPrice;"`
}

func (*BookUpdateImportPrice) TableName() string {
	return common.TableBook
}
