package importnotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/user/usermodel"
	"errors"
	"time"
)

type ImportNote struct {
	Id            string                `json:"id" gorm:"column:id;" example:"import note id"`
	SupplierId    string                `json:"-" gorm:"column:supplierId;"`
	Supplier      SimpleSupplier        `json:"supplier" gorm:"foreignKey:SupplierId;references:Id"`
	TotalPrice    int                   `json:"totalPrice" gorm:"column:totalPrice;" example:"120000"`
	Status        *ImportNoteStatus     `json:"status" gorm:"column:status;" example:"Done"`
	CreatedBy     string                `json:"-" gorm:"column:createdBy;"`
	CreatedByUser usermodel.SimpleUser  `json:"createdBy" gorm:"foreignKey:CreatedBy"`
	ClosedBy      *string               `json:"-" gorm:"column:closedBy;"`
	ClosedByUser  *usermodel.SimpleUser `json:"closedBy" gorm:"foreignKey:ClosedBy"`
	CreatedAt     *time.Time            `json:"createdAt" gorm:"column:createdAt;" example:"2023-12-03T15:02:19.62113565Z"`
	ClosedAt      *time.Time            `json:"closedAt" gorm:"column:closedAt;" example:"2023-12-03T15:02:19.62113565Z"`
}

func (*ImportNote) TableName() string {
	return common.TableImportNote
}

var (
	ErrImportNoteIdInvalid = common.NewCustomError(
		errors.New("id of import note is invalid"),
		"Mã của phiếu nhập không hợp lệ",
		"ErrImportNoteIdInvalid",
	)
	ErrImportNoteSupplierIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"Mã của nhà cung cấp hợp lệ",
		"ErrImportNoteSupplierIdInvalid",
	)
	ErrImportNoteDetailsEmpty = common.NewCustomError(
		errors.New("list import note details are empty"),
		"Danh sách sản phẩm muốn nhập đang trống",
		"ErrImportNoteDetailsEmpty",
	)
	ErrImportNoteStatusEmpty = common.NewCustomError(
		errors.New("import's status is empty"),
		"Trạng thái phiếu nhập muốn chuyển sang đang trống",
		"ErrImportNoteStatusEmpty",
	)
	ErrImportNoteStatusInvalid = common.NewCustomError(
		errors.New("import's status is invalid"),
		"Trạng thái phiếu nhập muốn chuyển sang không hợp lệ",
		"ErrImportNoteStatusInvalid",
	)
	ErrImportNoteHasSameBook = common.NewCustomError(
		errors.New("import note has duplicate book"),
		"Trong phiếu nhập đang có 2 sách giống nhau",
		"ErrImportNoteHasSameBook",
	)
	ErrImportNoteClosed = common.NewCustomError(
		errors.New("import note has been closed"),
		"Phiếu nhập đã đóng",
		"ErrImportNoteClosed",
	)
	ErrImportNoteIdDuplicate = common.ErrDuplicateKey(errors.New(
		"Phiếu nhập đã tồn tại"),
	)
	ErrImportNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo phiếu nhập mới"),
	)
	ErrImportNoteChangeStatusNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền thay đổi trạng thái phiếu nhập"),
	)
	ErrImportNoteViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem phiếu nhập"),
	)
)
