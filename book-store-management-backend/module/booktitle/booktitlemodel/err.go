package booktitlemodel

import (
	"book-store-management-backend/common"
	"errors"
)

var (
	ErrBookTitleNotFound = common.NewCustomError(
		errors.New("book Title not found"),
		"BookTitle not found",
		"ErrBookTitleNotFound",
	)

	ErrBookTitleNameEmpty = common.NewCustomError(
		errors.New("name of Book Title is empty"),
		"name of Book Title is empty",
		"ErrBookTitleNameEmpty",
	)

	ErrBookTitleAuthorIdsEmpty = common.NewCustomError(
		errors.New("author IDs of Book Title are empty"),
		"author IDs of Book Title are empty",
		"ErrBookTitleAuthorIdsEmpty",
	)

	ErrBookTitleCategoryIdsEmpty = common.NewCustomError(
		errors.New("category IDs of Book Title are empty"),
		"category IDs of Book Title are empty",
		"ErrBookTitleCategoryIdsEmpty",
	)

	ErrBookTitleValidateAuthor = common.NewCustomError(
		errors.New("author IDs of Book Title are invalid"),
		"author IDs of Book Title are invalid",
		"ErrBookTitleValidateAuthor",
	)

	ErrBookTitleValidateCategory = common.NewCustomError(
		errors.New("category IDs of Book Title are invalid"),
		"category IDs of Book Title are invalid",
		"ErrBookTitleValidateCategory",
	)

	ErrBookTitleIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of Book Title is duplicate"),
	)
)
