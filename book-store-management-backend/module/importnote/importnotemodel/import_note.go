package importnotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/user/usermodel"
	"errors"
	"time"
)

type ImportNote struct {
	Id           string                `json:"id" gorm:"column:id;" example:"import note id"`
	SupplierId   string                `json:"-" gorm:"column:supplierId;"`
	Supplier     SimpleSupplier        `json:"supplier" gorm:"foreignKey:SupplierId;references:Id"`
	TotalPrice   float32               `json:"totalPrice" gorm:"column:totalPrice;" example:"120000"`
	Status       *ImportNoteStatus     `json:"status" gorm:"column:status;" example:"Done"`
	CreateBy     string                `json:"-" gorm:"column:createBy;"`
	CreateByUser usermodel.SimpleUser  `json:"createBy" gorm:"foreignKey:CreateBy"`
	CloseBy      *string               `json:"-" gorm:"column:closeBy;"`
	CloseByUser  *usermodel.SimpleUser `json:"closeBy" gorm:"foreignKey:CloseBy"`
	CreateAt     *time.Time            `json:"createAt" gorm:"column:createAt;" example:"2023-12-03T15:02:19.62113565Z"`
	CloseAt      *time.Time            `json:"closeAt" gorm:"column:closeAt;" example:"2023-12-03T15:02:19.62113565Z"`
}

type SimpleSupplier struct {
	Id   string `json:"id" gorm:"column:id;" example:"supplier id"`
	Name string `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
}

func (*SimpleSupplier) TableName() string {
	return common.TableSupplier
}

func (*ImportNote) TableName() string {
	return common.TableImportNote
}

var (
	ErrImportNoteIdInvalid = common.NewCustomError(
		errors.New("id of import note is invalid"),
		"id of import note is invalid",
		"ErrImportNoteIdInvalid",
	)
	ErrImportNoteSupplierIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"id of supplier is invalid",
		"ErrImportNoteSupplierIdInvalid",
	)
	ErrImportNoteDetailsEmpty = common.NewCustomError(
		errors.New("list import note details are empty"),
		"list import note details are empty",
		"ErrImportNoteDetailsEmpty",
	)
	ErrImportNoteStatusEmpty = common.NewCustomError(
		errors.New("import's status is empty"),
		"import's status is empty",
		"ErrImportNoteStatusEmpty",
	)
	ErrImportNoteStatusInvalid = common.NewCustomError(
		errors.New("import's status is invalid"),
		"import's status is invalid",
		"ErrImportNoteStatusInvalid",
	)
	ErrImportNoteHasSameBook = common.NewCustomError(
		errors.New("import note has duplicate book"),
		"import note has duplicate book",
		"ErrImportNoteHasSameBook",
	)
	ErrImportNoteClosed = common.NewCustomError(
		errors.New("import note has been closed"),
		"import note has been closed",
		"ErrImportNoteClosed",
	)
	ErrImportNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create import note"),
	)
	ErrImportNoteChangeStatusNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to change status import note"),
	)
	ErrImportNoteViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view import note"),
	)
)
