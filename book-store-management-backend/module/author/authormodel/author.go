package authormodel

import (
	"book-store-management-backend/common"
	"errors"
	"time"
)

type Author struct {
	Id        string     `json:"id" json:"column:id;"`
	Name      string     `json:"name" json:"column:name;"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"createdAt; column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"deletedAt; column:deletedAt;"`
	IsActive  *bool      `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1"`
}

func (*Author) TableName() string {
	return common.TableAuthor
}

var (
	ErrAuthorIdInvalid = common.NewCustomError(
		errors.New("id of Author is invalid"),
		`Mã của tác giả không hợp lệ`,
		"ErrAuthorIdInvalid",
	)
	ErrAuthorNameEmpty = common.NewCustomError(
		errors.New("name of Author is empty"),
		"Tên của tác giả đang trống",
		"ErrAuthorNameEmpty",
	)
	ErrAuthorIdDuplicate = common.ErrDuplicateKey(
		errors.New("Tác giả đã tồn tại"),
	)
	ErrAuthorNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên tác giả đã tồn tại"),
	)
	ErrAuthorCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo tác giả mới"),
	)
	ErrAuthorViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem tác giả"),
	)
	ErrAuthorUpdateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin tác giả"),
	)
	ErrAuthorDeleteNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xóa tác giả"),
	)
)

func (data *Author) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrAuthorNameEmpty
	}
	return nil
}
