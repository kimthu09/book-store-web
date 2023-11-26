package inventorychecknotedetailmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"errors"
)

type InventoryCheckNoteDetail struct {
	InventoryCheckNoteId string         `json:"inventoryCheckNoteId" gorm:"column:inventoryCheckNoteId;"`
	BookId               string         `json:"-" gorm:"column:bookId;"`
	Book                 bookmodel.Book `json:"book"`
	Initial              float32        `json:"initial" gorm:"column:initial;"`
	Difference           float32        `json:"difference" gorm:"column:difference;"`
	Final                float32        `json:"final" gorm:"column:final;"`
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
