package booktitlemodel

import (
	"book-store-management-backend/common"
	"errors"
)

var (
	ErrBookTitleNotFound = common.NewCustomError(
		errors.New("book Title not found"),
		"Không thể tìm thấy đầu sách",
		"ErrBookTitleNotFound",
	)
	ErrBookTitleIdInvalid = common.NewCustomError(
		errors.New("id of Book Title is invalid"),
		"Mã của đầu sách không hợp lệ",
		"ErrBookTitleIdInvalid",
	)

	ErrBookTitleNameEmpty = common.NewCustomError(
		errors.New("name of Book Title is empty"),
		"Tên của đầu sách đang trống",
		"ErrBookTitleNameEmpty",
	)
	ErrBookTitlePublisherIdEmpty = common.NewCustomError(
		errors.New("publisher ID of Book Title is empty"),
		"publisher ID of Book Title is empty",
		"ErrBookTitlePublisherIdEmpty",
	)
	ErrBookTitleAuthorIdsEmpty = common.NewCustomError(
		errors.New("author IDs of Book Title are empty"),
		"Thông tin nhà xuât bản đang trống",
		"ErrBookTitleAuthorIdsEmpty",
	)
	ErrBookTitleCategoryIdsEmpty = common.NewCustomError(
		errors.New("category IDs of Book Title are empty"),
		"Danh sách thể loại đang trống",
		"ErrBookTitleCategoryIdsEmpty",
	)

	ErrBookTitleValidatePublisher = common.NewCustomError(
		errors.New("publisher ID of Book Title is invalid"),
		"Nhà xuất bản không hợp lệ",
		"ErrBookTitleValidatePublisher",
	)
	ErrBookTitleValidateAuthor = common.NewCustomError(
		errors.New("author IDs of Book Title are invalid"),
		"Tác giả không hợp lệ",
		"ErrBookTitleValidateAuthor",
	)
	ErrBookTitleValidateCategory = common.NewCustomError(
		errors.New("category IDs of Book Title are invalid"),
		"Thể loại không hợp lệ",
		"ErrBookTitleValidateCategory",
	)

	ErrBookTitleIdDuplicate = common.ErrDuplicateKey(
		errors.New("Đầu sách đã tồn tại"),
	)

	ErrBookTitleCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo đầu sách mới"),
	)
	ErrBookTitleViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem đầu sách"),
	)
	ErrBookTitleUpdateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin đầu sách"),
	)
	ErrBookTitleDeleteNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xóa đầu sách"),
	)
)
