package bookmodel

import (
	"book-store-management-backend/common"
	"errors"
)

var (
	ErrBookCreateNoPermission = common.NewCustomError(
		errors.New("no permission to create book"),
		"Bạn không có quyền tạo sách mới",
		"ErrBookCreateNoPermission",
	)
	ErrBookUpdateNoPermission = common.NewCustomError(
		errors.New("no permission to update book"),
		"Bạn không có quyền chỉnh sửa thông tin sách",
		"ErrBookUpdateNoPermission",
	)
	ErrBookUpdateStatusNoPermission = common.NewCustomError(
		errors.New("no permission to update status book"),
		"Bạn không có quyền chỉnh sửa trạng thái sách",
		"ErrBookUpdateStatusNoPermission",
	)
	ErrBookViewNoPermission = common.NewCustomError(
		errors.New("no permission to view book"),
		"Bạn không có quyền xem sách",
		"ErrBookViewNoPermission",
	)
	ErrBookDeleteNoPermission = common.NewCustomError(
		errors.New("no permission to delete book"),
		"Bạn không có quyền xóa sách",
		"ErrBookDeleteNoPermission",
	)
	ErrBookStatusEmpty = common.NewCustomError(
		errors.New("book status is empty"),
		"Trạng thái sách đang trống",
		"ErrBookStatusEmpty",
	)
	ErrBookIdInvalid = common.NewCustomError(
		errors.New("id of Book is invalid"),
		"Mã của sách không hợp lệ",
		"ErrBookIdInvalid",
	)
	ErrBookTitleIdInvalid = common.NewCustomError(
		errors.New("id of Book Title is invalid"),
		"Đầu sách không hợp lệ",
		"ErrBookTitleIdInvalid",
	)
	ErrPublisherIdInvalid = common.NewCustomError(
		errors.New("id of Publisher is invalid"),
		"Nhà xuất bản không hợp lệ",
		"ErrPublisherIdInvalid",
	)
	ErrBookEditionInvalid = common.NewCustomError(
		errors.New("edition of Book is invalid"),
		"Lần tái bản không hợp lệ",
		"ErrBookEditionInvalid",
	)
	ErrBookListedPriceInvalid = common.NewCustomError(
		errors.New("listed price of Book is invalid"),
		"Giá niêm yết không hợp lệ",
		"ErrBookListedPriceInvalid",
	)
	ErrBookSellPriceInvalid = common.NewCustomError(
		errors.New("sell price of Book is invalid"),
		"Giá bán không hợp lệ",
		"ErrBookSellPriceInvalid",
	)
	ErrBookImageInvalid = common.NewCustomError(
		errors.New("image of book is invalid"),
		"Ảnh của sách không hợp lệ",
		"ErrBookImageInvalid",
	)
	ErrBookIdDuplicate = common.ErrDuplicateKey(
		errors.New("Sách đã tồn tại"),
	)
)
