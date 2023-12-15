package inventorychecknotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
)

type ReqCreateInventoryCheckNote struct {
	Id                  *string                                                        `json:"id" gorm:"column:id;" example:""`
	QuantityDifferent   int                                                            `json:"-" gorm:"column:qtyDifferent;"`
	QuantityAfterAdjust int                                                            `json:"-" gorm:"column:qtyAfterAdjust;"`
	CreatedBy           string                                                         `json:"-" gorm:"column:createdBy;"`
	Details             []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate `json:"details" gorm:"-"`
}

func (*ReqCreateInventoryCheckNote) TableName() string {
	return common.TableInventoryCheckNote
}

func (data *ReqCreateInventoryCheckNote) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrInventoryCheckNoteIdInvalid
	}

	mapExits := make(map[string]int)
	for _, detail := range data.Details {
		if err := detail.Validate(); err != nil {
			return err
		}
		mapExits[detail.BookId]++
		if mapExits[detail.BookId] >= 2 {
			return ErrInventoryCheckNoteExistDuplicateBook
		}
	}
	return nil
}
