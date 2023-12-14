package importnotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"book-store-management-backend/module/user/usermodel"
	"time"
)

type ResDetailImportNote struct {
	Id            string                                   `json:"id" gorm:"column:id;" example:"import note id"`
	SupplierId    string                                   `json:"-" gorm:"column:supplierId;"`
	Supplier      SimpleSupplier                           `json:"supplier" gorm:"foreignKey:SupplierId;references:Id"`
	TotalPrice    int                                      `json:"totalPrice" gorm:"column:totalPrice;" example:"120000"`
	Status        *ImportNoteStatus                        `json:"status" gorm:"column:status;" example:"Done"`
	CreatedBy     string                                   `json:"-" gorm:"column:createdBy;"`
	CreatedByUser usermodel.SimpleUser                     `json:"createdBy" gorm:"foreignKey:CreatedBy"`
	ClosedBy      *string                                  `json:"-" gorm:"column:closedBy;"`
	ClosedByUser  *usermodel.SimpleUser                    `json:"closedBy" gorm:"foreignKey:ClosedBy"`
	CreatedAt     *time.Time                               `json:"createdAt" gorm:"column:createdAt;" example:"2023-12-03T15:02:19.62113565Z"`
	ClosedAt      *time.Time                               `json:"closedAt" gorm:"column:closedAt;" example:"2023-12-03T15:02:19.62113565Z"`
	Details       []importnotedetailmodel.ImportNoteDetail `json:"details"`
}

func (*ResDetailImportNote) TableName() string {
	return common.TableImportNote
}

func GetResDetailImportNoteFromImportNote(importNote *ImportNote) *ResDetailImportNote {
	var src ResDetailImportNote
	src.Id = importNote.Id
	src.SupplierId = importNote.SupplierId
	src.Supplier = importNote.Supplier
	src.TotalPrice = importNote.TotalPrice
	src.Status = importNote.Status
	src.CreatedBy = importNote.CreatedBy
	src.CreatedByUser = importNote.CreatedByUser
	src.ClosedBy = importNote.ClosedBy
	src.ClosedByUser = importNote.ClosedByUser
	src.CreatedAt = importNote.CreatedAt
	src.ClosedAt = importNote.ClosedAt
	return &src
}
