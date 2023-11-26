package importnotedetailmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"errors"
)

type ImportNoteDetail struct {
	ImportNoteId string         `json:"importNoteId" gorm:"column:importNoteId;"`
	BookId       string         `json:"bookId" gorm:"column:bookId;"`
	Book         bookmodel.Book `json:"book"`
	AmountImport float32        `json:"amountImport" gorm:"column:amountImport;"`
	Price        float32        `json:"price" gorm:"column:price;"`
}

func (*ImportNoteDetail) TableName() string {
	return common.TableImportNoteDetail
}

var (
	ErrImportDetailBookIdInvalid = common.NewCustomError(
		errors.New("id of book is invalid"),
		"id of book is invalid",
		"ErrImportDetailBookIdInvalid",
	)
	ErrImportDetailPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price of ingredient is negative number"),
		"price of ingredient is negative number",
		"ErrImportDetailPriceIsNegativeNumber",
	)
	ErrImportDetailAmountImportIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount import is not positive number"),
		"amount import is not positive number",
		"ErrImportDetailAmountImportIsNotPositiveNumber",
	)
)
