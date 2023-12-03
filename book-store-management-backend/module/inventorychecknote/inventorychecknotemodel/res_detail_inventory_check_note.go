package inventorychecknotemodel

import (
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"book-store-management-backend/module/user/usermodel"
	"time"
)

type ResDetailInventoryCheckNote struct {
	Id                  string                                                   `json:"id" gorm:"column:id;" example:"inventory check note id"`
	QuantityDifferent   int                                                      `json:"qtyDifferent" gorm:"column:qtyDifferent;" example:"100"`
	QuantityAfterAdjust int                                                      `json:"qtyAfterAdjust" gorm:"column:qtyAfterAdjust;" example:"200"`
	CreateBy            string                                                   `json:"-" gorm:"column:createBy;"`
	CreateByUser        usermodel.SimpleUser                                     `json:"createBy" gorm:"foreignKey:CreateBy;references:Id"`
	CreateAt            *time.Time                                               `json:"createAt" gorm:"column:createAt;" example:"2023-12-03T15:02:19.62113565Z"`
	Details             []inventorychecknotedetailmodel.InventoryCheckNoteDetail `json:"details"`
}

func GetResDetailInventoryCheckNoteFromInventoryCheckNote(inventoryCheckNote *InventoryCheckNote) *ResDetailInventoryCheckNote {
	var src ResDetailInventoryCheckNote
	src.Id = inventoryCheckNote.Id
	src.QuantityDifferent = inventoryCheckNote.QuantityDifferent
	src.QuantityAfterAdjust = inventoryCheckNote.QuantityAfterAdjust
	src.CreateBy = inventoryCheckNote.CreateBy
	src.CreateByUser = inventoryCheckNote.CreateByUser
	src.CreateAt = inventoryCheckNote.CreateAt
	return &src
}
