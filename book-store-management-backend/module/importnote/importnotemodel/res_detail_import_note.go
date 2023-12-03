package importnotemodel

import (
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"book-store-management-backend/module/user/usermodel"
	"time"
)

type ResDetailImportNote struct {
	Id           string                                   `json:"id" gorm:"column:id;" example:"import note id"`
	SupplierId   string                                   `json:"-" gorm:"column:supplierId;"`
	Supplier     SimpleSupplier                           `json:"supplier" gorm:"foreignKey:SupplierId;references:Id"`
	TotalPrice   float32                                  `json:"totalPrice" gorm:"column:totalPrice;" example:"120000"`
	Status       *ImportNoteStatus                        `json:"status" gorm:"column:status;" example:"Done"`
	CreateBy     string                                   `json:"-" gorm:"column:createBy;"`
	CreateByUser usermodel.SimpleUser                     `json:"createBy" gorm:"foreignKey:CreateBy"`
	CloseBy      *string                                  `json:"-" gorm:"column:closeBy;"`
	CloseByUser  *usermodel.SimpleUser                    `json:"closeBy" gorm:"foreignKey:CloseBy"`
	CreateAt     *time.Time                               `json:"createAt" gorm:"column:createAt;" example:"2023-12-03T15:02:19.62113565Z"`
	CloseAt      *time.Time                               `json:"closeAt" gorm:"column:closeAt;" example:"2023-12-03T15:02:19.62113565Z"`
	Details      []importnotedetailmodel.ImportNoteDetail `json:"details"`
}

func GetResDetailImportNoteFromImportNote(importNote *ImportNote) *ResDetailImportNote {
	var src ResDetailImportNote
	src.Id = importNote.Id
	src.SupplierId = importNote.SupplierId
	src.Supplier = importNote.Supplier
	src.TotalPrice = importNote.TotalPrice
	src.Status = importNote.Status
	src.CreateBy = importNote.CreateBy
	src.CreateByUser = importNote.CreateByUser
	src.CloseBy = importNote.CloseBy
	src.CloseByUser = importNote.CloseByUser
	src.CreateAt = importNote.CreateAt
	src.CloseAt = importNote.CloseAt
	return &src
}
