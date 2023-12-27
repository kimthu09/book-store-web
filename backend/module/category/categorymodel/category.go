package categorymodel

import (
	"book-store-management-backend/common"
	"errors"
	"time"
)

type Category struct {
	Id        string     `json:"id" json:"column:id;"`
	Name      string     `json:"name" json:"column:name;"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"createdAt; column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"deletedAt; column:deletedAt;"`
	IsActive  *bool      `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1"`
}

func (*Category) TableName() string {
	return common.TableCategory
}

var (
	ErrCategoryIdInvalid = common.NewCustomError(
		errors.New("id of Category is invalid"),
		`Mã của thể loại không hợp lệ`,
		"ErrCategoryIdInvalid",
	)
	ErrCategoryNameEmpty = common.NewCustomError(
		errors.New("name of Category is empty"),
		"Tên của thể loại đang trống",
		"ErrCategoryNameEmpty",
	)
	ErrCategoryIdDuplicate = common.ErrDuplicateKey(
		errors.New("Thể loại đã tồn tại"),
	)
	ErrCategoryNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên thể loại đã tồn tại"),
	)
	ErrCategoryCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo thể loại mới"),
	)
	ErrCategoryViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem thể loại"),
	)
	ErrCategoryUpdateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin thể loại"),
	)
	ErrCategoryDeleteNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xóa thể loại"),
	)
)

func (data *Category) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrCategoryNameEmpty
	}
	return nil
}
