package bookmodel

import (
	"book-store-management-backend/common"
)

type BookUpdateQuantity struct {
	QuantityUpdate int `json:"qtyUpdate" gorm:"-"`
}

func (*BookUpdateQuantity) TableName() string {
	return common.TableBook
}
