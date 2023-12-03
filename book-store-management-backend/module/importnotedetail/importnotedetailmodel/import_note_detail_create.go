package importnotedetailmodel

import "book-store-management-backend/common"

type ImportNoteDetailCreate struct {
	ImportNoteId   string  `json:"-" gorm:"column:importNoteId;"`
	BookId         string  `json:"bookId" gorm:"column:bookId;" example:"book id"`
	QuantityImport float32 `json:"qtyImport" gorm:"column:qtyImport;" example:"100"`
	Price          float32 `json:"price" gorm:"column:price;" example:"60000"`
	IsReplacePrice bool    `json:"isReplacePrice" gorm:"-" example:"true"`
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

func (data *ImportNoteDetailCreate) Round() {
	common.CustomRound(&data.Price)
	common.CustomRound(&data.QuantityImport)
}
