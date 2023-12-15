package importnotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
)

type ReqCreateImportNote struct {
	Id                *string                                        `json:"id" gorm:"column:id;" example:""`
	TotalPrice        int                                            `json:"-" gorm:"column:totalPrice;"`
	SupplierId        string                                         `json:"supplierId" gorm:"column:supplierId" example:"supplier id"`
	CreatedBy         string                                         `json:"-" gorm:"column:createdBy;"`
	ImportNoteDetails []importnotedetailmodel.ImportNoteDetailCreate `json:"details" gorm:"-"`
}

func (*ReqCreateImportNote) TableName() string {
	return common.TableImportNote
}

func (data *ReqCreateImportNote) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrImportNoteIdInvalid
	}
	if !common.ValidateNotNilId(&data.SupplierId) {
		return ErrImportNoteSupplierIdInvalid
	}
	if data.ImportNoteDetails == nil || len(data.ImportNoteDetails) == 0 {
		return ErrImportNoteDetailsEmpty
	}

	mapBookUpdatePriceTimes := make(map[string]int)
	for _, importNoteDetail := range data.ImportNoteDetails {
		if err := importNoteDetail.Validate(); err != nil {
			return err
		}
		mapBookUpdatePriceTimes[importNoteDetail.BookId]++
		if mapBookUpdatePriceTimes[importNoteDetail.BookId] > 1 {
			return ErrImportNoteHasSameBook
		}
	}
	return nil
}
