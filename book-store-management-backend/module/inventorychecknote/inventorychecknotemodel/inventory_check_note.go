package inventorychecknotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"book-store-management-backend/module/user/usermodel"
	"errors"
	"time"
)

type InventoryCheckNote struct {
	Id                  string                                                   `json:"id" gorm:"column:id;"`
	QuantityDifferent   int                                                      `json:"qtyDifferent" gorm:"column:qtyDifferent;"`
	QuantityAfterAdjust int                                                      `json:"qtyAfterAdjust" gorm:"column:qtyAfterAdjust;"`
	CreateBy            string                                                   `json:"-" gorm:"column:createBy;"`
	CreateByUser        usermodel.SimpleUser                                     `json:"createBy" gorm:"foreignKey:CreateBy;references:Id"`
	CreateAt            *time.Time                                               `json:"createAt" gorm:"column:createAt;"`
	Details             []inventorychecknotedetailmodel.InventoryCheckNoteDetail `json:"details"`
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
	ErrInventoryCheckNoteModifyQuantityIsInvalid = common.NewCustomError(
		errors.New("the qty after modification is invalid"),
		"the qty after modification is invalid",
		"ErrInventoryCheckNoteModifyQuantityIsInvalid",
	)
	ErrInventoryCheckNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create inventory check note"),
	)
	ErrInventoryCheckNoteViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view inventory check note"),
	)
)
