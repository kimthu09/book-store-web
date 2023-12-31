package inventorychecknotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/user/usermodel"
	"errors"
	"time"
)

type InventoryCheckNote struct {
	Id                  string               `json:"id" gorm:"column:id;" example:"inventory check note id"`
	QuantityDifferent   int                  `json:"qtyDifferent" gorm:"column:qtyDifferent;" example:"100"`
	QuantityAfterAdjust int                  `json:"qtyAfterAdjust" gorm:"column:qtyAfterAdjust;" example:"200"`
	CreatedBy           string               `json:"-" gorm:"column:createdBy;"`
	CreatedByUser       usermodel.SimpleUser `json:"createdBy" gorm:"foreignKey:CreatedBy;references:Id"`
	CreatedAt           *time.Time           `json:"createdAt" gorm:"column:createdAt;" example:"2023-12-03T15:02:19.62113565Z"`
}

func (*InventoryCheckNote) TableName() string {
	return common.TableInventoryCheckNote
}

var (
	ErrInventoryCheckNoteIdInvalid = common.NewCustomError(
		errors.New("id of inventory check note is invalid"),
		"Mã của phiếu kiểm kho không hợp lệ",
		"ErrInventoryCheckNoteIdInvalid",
	)
	ErrInventoryCheckNoteExistDuplicateBook = common.NewCustomError(
		errors.New("exist duplicate book"),
		"Trong phiếu nhập đang có 2 sách giống nhau",
		"ErrInventoryCheckNoteExistDuplicateBook",
	)
	ErrInventoryCheckNoteModifyQuantityIsInvalid = common.NewCustomError(
		errors.New("the qty after modification is invalid"),
		"Số lượng sau khi điều chỉnh không hợp lệ",
		"ErrInventoryCheckNoteModifyQuantityIsInvalid",
	)
	ErrInventoryCheckNoteIdDuplicate = common.ErrDuplicateKey(errors.New(
		"Phiếu kiểm kho đã tồn tại"),
	)
	ErrInventoryCheckNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo phiếu kiểm kho mới"),
	)
	ErrInventoryCheckNoteViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem phiếu kiểm kho"),
	)
)
