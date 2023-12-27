package invoicedetailmodel

import (
	"book-store-management-backend/common"
)

type ReqCreateInvoiceDetail struct {
	InvoiceId string `json:"-" gorm:"column:invoiceId;" example:"123"`
	BookId    string `json:"bookId" gorm:"column:bookId;" example:"book id"`
	BookName  string `json:"-" gorm:"column:bookName;"`
	Quantity  int    `json:"qty" gorm:"column:qty;" example:"2"`
	UnitPrice int    `json:"-" gorm:"column:unitPrice" example:"60000"`
}

func (*ReqCreateInvoiceDetail) TableName() string {
	return common.TableInvoiceDetail
}

func (data *ReqCreateInvoiceDetail) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.BookId) {
		return ErrInvoiceDetailBookIdInvalid
	}
	if common.ValidateNotPositiveNumber(data.Quantity) {
		return ErrInvoiceDetailAmountIsNotPositiveNumber
	}
	return nil
}
