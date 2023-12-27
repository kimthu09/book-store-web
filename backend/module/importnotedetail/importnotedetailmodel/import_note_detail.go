package importnotedetailmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"errors"
)

type ImportNoteDetail struct {
	ImportNoteId   string               `json:"importNoteId" gorm:"column:importNoteId;" example:"import note id"`
	BookId         string               `json:"-" gorm:"column:bookId;"`
	Book           bookmodel.SimpleBook `json:"book"`
	QuantityImport int                  `json:"qtyImport" gorm:"column:qtyImport;"`
	Price          float32              `json:"price" gorm:"column:price;"`
}

func (*ImportNoteDetail) TableName() string {
	return common.TableImportNoteDetail
}

var (
	ErrImportDetailBookIdInvalid = common.NewCustomError(
		errors.New("id of book is invalid"),
		"Mã của sách không hợp lệ",
		"ErrImportDetailBookIdInvalid",
	)
	ErrImportDetailPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price of book is negative number"),
		"Giá nhập sách không được là số âm",
		"ErrImportDetailPriceIsNegativeNumber",
	)
	ErrImportDetailQuantityImportIsNotPositiveNumber = common.NewCustomError(
		errors.New("quantity import is not positive number"),
		"Số lượng nhập sách phải là số dương",
		"ErrImportDetailQuantityImportIsNotPositiveNumber",
	)
)
