package publishermodel

import (
	"book-store-management-backend/common"
	"errors"
)

type Publisher struct {
	Id   string `json:"id" json:"column:id;" example:"publisher id"`
	Name string `json:"name" json:"column:name;" example:"NXB Kim Đồng"`
}

func (*Publisher) TableName() string {
	return common.TablePublisher
}

var (
	ErrPublisherIdInvalid = common.NewCustomError(
		errors.New("id of Publisher is invalid"),
		`id của nhà xuất bản không hợp lệ`,
		"ErrPublisherIdInvalid",
	)
	ErrPublisherNameEmpty = common.NewCustomError(
		errors.New("name of Publisher is empty"),
		"Tên của nhà xuất bản đang trống",
		"ErrPublisherNameEmpty",
	)
	ErrPublisherIdDuplicate = common.ErrDuplicateKey(
		errors.New("Id của nhà xuất bản đã tồn tại"),
	)
	ErrPublisherNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên của nhà xuất bản đã tồn tại"),
	)
	ErrPublisherCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo nhà xuất bản mới"),
	)
	ErrPublisherViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem nhà xuất bản"),
	)
	ErrPublisherUpdateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin nhà xuất bản"),
	)
	ErrPublisherDeleteNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xóa nhà xuất bản"),
	)
)

func (data *Publisher) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrPublisherNameEmpty
	}
	return nil
}
