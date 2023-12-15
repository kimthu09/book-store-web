package invoicedetailmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"errors"
)

type InvoiceDetail struct {
	InvoiceId string               `json:"-" gorm:"column:invoiceId;" example:"123"`
	BookId    string               `json:"-" gorm:"column:bookId;"`
	Book      bookmodel.SimpleBook `json:"book" gorm:"foreignKey:bookId;references:ID"`
	Quantity  int                  `json:"qty" gorm:"column:qty;" example:"2"`
	UnitPrice int                  `json:"unitPrice" gorm:"column:unitPrice" example:"60000"`
}

func (*InvoiceDetail) TableName() string {
	return common.TableInvoiceDetail
}

var (
	ErrInvoiceDetailBookIdInvalid = common.NewCustomError(
		errors.New("id of book is invalid"),
		"Mã của sách không hợp lệ",
		"ErrInvoiceDetailBookIdInvalid",
	)
	ErrInvoiceDetailAmountIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount import is not positive number"),
		"Số lượng sản phẩm bán ra phải là số dương",
		"ErrInvoiceDetailAmountIsNotPositiveNumber",
	)
	ErrInvoiceDetailBookIsInactive = common.NewCustomError(
		errors.New("book is inactive"),
		"Tồn tại sách không còn bán nữa",
		"ErrInvoiceDetailBookIsInactive",
	)
)
