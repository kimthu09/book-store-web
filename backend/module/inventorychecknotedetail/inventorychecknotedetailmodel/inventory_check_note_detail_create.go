package inventorychecknotedetailmodel

import "book-store-management-backend/common"

type InventoryCheckNoteDetailCreate struct {
	InventoryCheckNoteId string `json:"-" gorm:"column:inventoryCheckNoteId;"`
	BookId               string `json:"bookId" gorm:"column:bookId;" example:"book id"`
	Initial              int    `json:"-" gorm:"column:initial;"`
	Difference           int    `json:"difference" gorm:"column:difference;" example:"100"`
	Final                int    `json:"-" gorm:"column:final;"`
}

func (*InventoryCheckNoteDetailCreate) TableName() string {
	return common.TableInventoryCheckNoteDetail
}

func (data *InventoryCheckNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.BookId) {
		return ErrInventoryCheckDetailBookIdInvalid
	}
	if data.Difference == 0 {
		return ErrInventoryCheckDifferenceIsInvalid
	}
	return nil
}
