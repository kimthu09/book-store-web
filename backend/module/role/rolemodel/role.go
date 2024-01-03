package rolemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/rolefeature/rolefeaturemodel"
	"errors"
)

type Role struct {
	Id           string                               `json:"id" gorm:"column:id;" example:"role id"`
	Name         string                               `json:"name" gorm:"column:name;" example:"admin"`
	RoleFeatures []rolefeaturemodel.SimpleRoleFeature `json:"features"`
}

func (*Role) TableName() string {
	return common.TableRole
}

var (
	ErrRoleNameEmpty = common.NewCustomError(
		errors.New("name of role is empty"),
		"Tên của vai trò đang trống",
		"ErrRoleNameEmpty",
	)
	ErrRoleFeaturesEmpty = common.NewCustomError(
		errors.New("features of role is empty"),
		"Danh sách chức năng của vai trò đang trống",
		"ErrRoleFeaturesEmpty",
	)
	ErrRoleFeatureInvalid = common.NewCustomError(
		errors.New("features of role is invalid"),
		"Chức năng không hợp lệ",
		"ErrRoleFeatureInvalid",
	)
	ErrRoleNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên của vai trò đã tồn tại"),
	)
	ErrRoleCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo vai trò mới"),
	)
	ErrRoleUpdateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin vai trò"),
	)
	ErrRoleViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem vai trò"),
	)
)
