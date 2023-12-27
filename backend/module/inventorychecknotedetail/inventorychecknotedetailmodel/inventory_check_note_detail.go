package inventorychecknotedetailmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"errors"
)

type InventoryCheckNoteDetail struct {
	InventoryCheckNoteId string               `json:"inventoryCheckNoteId" gorm:"column:inventoryCheckNoteId;" example:"inventory check note id"`
	BookId               string               `json:"-" gorm:"column:bookId;"`
	Book                 bookmodel.SimpleBook `json:"book"`
	Initial              int                  `json:"initial" gorm:"column:initial;" example:"100"`
	Difference           int                  `json:"difference" gorm:"column:difference;" example:"100"`
	Final                int                  `json:"final" gorm:"column:final;" example:"200"`
}

func (*InventoryCheckNoteDetail) TableName() string {
	return common.TableInventoryCheckNoteDetail
}

var (
	ErrInventoryCheckDetailBookIdInvalid = common.NewCustomError(
		errors.New("id of book is invalid"),
		"Mã của sách không hợp lệ",
		"ErrInventoryCheckDetailBookIdInvalid",
	)
	ErrInventoryCheckDifferenceIsInvalid = common.NewCustomError(
		errors.New("difference is invalid"),
		"Số lượng chỉnh sửa không hợp lệ",
		"ErrInventoryCheckDifferenceIsInvalid",
	)
)
