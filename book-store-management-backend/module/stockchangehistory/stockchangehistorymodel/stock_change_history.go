package stockchangehistorymodel

import "book-store-management-backend/common"

type StockChangeHistory struct {
	Id           string           `json:"id" gorm:"column:id;"`
	BookId       string           `json:"bookId" gorm:"column:bookId;"`
	Quantity     int              `json:"qty" gorm:"column:qty;"`
	QuantityLeft int              `json:"qtyLeft" gorm:"column:qtyLeft;"`
	Type         *StockChangeType `json:"type" gorm:"column:type;"`
}

func (*StockChangeHistory) TableName() string {
	return common.TableStockChangeHistory
}
