package importnotedetailmodel

import "book-store-management-backend/common"

type ImportNoteDetailCreate struct {
	ImportNoteId   string `json:"-" gorm:"column:importNoteId;"`
	BookId         string `json:"bookId" gorm:"column:bookId;" example:"book id"`
	QuantityImport int    `json:"qtyImport" gorm:"column:qtyImport;" example:"100"`
	Price          int    `json:"price" gorm:"column:price;" example:"60000"`
	IsReplacePrice bool   `json:"isReplacePrice" gorm:"-" example:"true"`
}

func (*ImportNoteDetailCreate) TableName() string {
	return common.TableImportNoteDetail
}

func (data *ImportNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.BookId) {
		return ErrImportDetailBookIdInvalid
	}
	if common.ValidateNegativeNumber(data.Price) {
		return ErrImportDetailPriceIsNegativeNumber
	}
	if common.ValidateNotPositiveNumber(data.QuantityImport) {
		return ErrImportDetailQuantityImportIsNotPositiveNumber
	}
	return nil
}
