package importnotedetailmodel

import "book-store-management-backend/common"

type ImportNoteDetailCreate struct {
	ImportNoteId   string  `json:"importNoteId" gorm:"column:importNoteId;"`
	BookId         string  `json:"bookId" gorm:"column:bookId;"`
	QuantityImport float32 `json:"qtyImport" gorm:"column:qtyImport;"`
	Price          float32 `json:"price" gorm:"column:price;"`
	IsReplacePrice bool    `json:"isReplacePrice" gorm:"-"`
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
