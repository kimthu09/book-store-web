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
	CreatedBy           string                                                   `json:"-" gorm:"column:createdBy;"`
	CreatedByUser       usermodel.SimpleUser                                     `json:"createdBy" gorm:"foreignKey:CreatedBy;references:Id"`
	CreatedAt           *time.Time                                               `json:"createdAt" gorm:"column:createdAt;" example:"2023-12-03T15:02:19.62113565Z"`
	Details             []inventorychecknotedetailmodel.InventoryCheckNoteDetail `json:"details"`
}

func GetResDetailInventoryCheckNoteFromInventoryCheckNote(inventoryCheckNote *InventoryCheckNote) *ResDetailInventoryCheckNote {
	var src ResDetailInventoryCheckNote
	src.Id = inventoryCheckNote.Id
	src.QuantityDifferent = inventoryCheckNote.QuantityDifferent
	src.QuantityAfterAdjust = inventoryCheckNote.QuantityAfterAdjust
	src.CreatedBy = inventoryCheckNote.CreatedBy
	src.CreatedByUser = inventoryCheckNote.CreatedByUser
	src.CreatedAt = inventoryCheckNote.CreatedAt
	return &src
}
