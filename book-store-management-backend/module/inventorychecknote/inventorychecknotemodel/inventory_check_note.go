package inventorychecknotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"errors"
	"time"
)

type InventoryCheckNote struct {
	Id                string                                                   `json:"id" gorm:"column:id;"`
	AmountDifferent   float32                                                  `json:"amountDifferent" gorm:"column:amountDifferent;"`
	AmountAfterAdjust float32                                                  `json:"amountAfterAdjust" gorm:"column:amountAfterAdjust;"`
	CreateBy          string                                                   `json:"createBy" gorm:"column:createBy;"`
	CreateAt          *time.Time                                               `json:"createAt" gorm:"column:createAt;"`
	Details           []inventorychecknotedetailmodel.InventoryCheckNoteDetail `json:"details"`
}

func (*InventoryCheckNote) TableName() string {
	return common.TableInventoryCheckNote
}

var (
	ErrInventoryCheckNoteIdInvalid = common.NewCustomError(
		errors.New("id of inventory check note is invalid"),
		"id of inventory check note is invalid",
		"ErrInventoryCheckNoteIdInvalid",
	)
	ErrInventoryCheckNoteExistDuplicateBook = common.NewCustomError(
		errors.New("exist duplicate book"),
		"exist duplicate book",
		"ErrInventoryCheckNoteExistDuplicateBook",
	)
	ErrInventoryCheckNoteModifyAmountIsInvalid = common.NewCustomError(
		errors.New("the amount after modification is invalid"),
		"the amount after modification is invalid",
		"ErrInventoryCheckNoteModifyAmountIsInvalid",
	)
	ErrInventoryCheckNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create inventory check note"),
	)
	ErrInventoryCheckNoteViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view inventory check note"),
	)
)
