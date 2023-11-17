package importnotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
)

type ImportNoteCreate struct {
	Id                *string                                        `json:"id" gorm:"column:id;"`
	TotalPrice        float32                                        `json:"-" gorm:"column:totalPrice;"`
	SupplierId        string                                         `json:"supplierId" gorm:"column:supplierId"`
	CreateBy          string                                         `json:"-" gorm:"column:createBy;"`
	ImportNoteDetails []importnotedetailmodel.ImportNoteDetailCreate `json:"details" gorm:"-"`
}

func (*ImportNoteCreate) TableName() string {
	return common.TableImportNote
}

func (data *ImportNoteCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrImportNoteIdInvalid
	}
	if !common.ValidateNotNilId(&data.SupplierId) {
		return ErrImportNoteSupplierIdInvalid
	}
	if data.ImportNoteDetails == nil || len(data.ImportNoteDetails) == 0 {
		return ErrImportNoteDetailsEmpty
	}

	mapIngredientUpdatePriceTimes := make(map[string]int)
	for _, importNoteDetail := range data.ImportNoteDetails {
		if err := importNoteDetail.Validate(); err != nil {
			return err
		}
		mapIngredientUpdatePriceTimes[importNoteDetail.BookId]++
		if mapIngredientUpdatePriceTimes[importNoteDetail.BookId] > 1 {
			return ErrImportNoteHasSameBook
		}
	}
	return nil
}
