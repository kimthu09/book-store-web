package inventorychecknotedetailmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"errors"
)

type InventoryCheckNoteDetail struct {
	InventoryCheckNoteId string               `json:"inventoryCheckNoteId" gorm:"column:inventoryCheckNoteId;"`
	BookId               string               `json:"-" gorm:"column:bookId;"`
	Book                 bookmodel.SimpleBook `json:"book"`
	Initial              int                  `json:"initial" gorm:"column:initial;"`
	Difference           int                  `json:"difference" gorm:"column:difference;"`
	Final                int                  `json:"final" gorm:"column:final;"`
}

func (*InventoryCheckNoteDetail) TableName() string {
	return common.TableInventoryCheckNoteDetail
}

var (
	ErrInventoryCheckDetailBookIdInvalid = common.NewCustomError(
		errors.New("id of book is invalid"),
		"id of book is is invalid",
		"ErrInventoryCheckDetailBookIdInvalid",
	)
	ErrInventoryCheckDifferenceIsInvalid = common.NewCustomError(
		errors.New("difference is invalid"),
		"difference is is invalid",
		"ErrInventoryCheckDifferenceIsInvalid",
	)
)
