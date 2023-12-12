package booktitlemodel

import (
	"book-store-management-backend/common"
	"errors"
)

var (
	ErrBookTitleNotFound = common.NewCustomError(
		errors.New("Book Title not found"),
		"BookTitle not found",
		"ErrBookTitleNotFound",
	)
	ErrBookTitleIdInvalid = common.NewCustomError(
		errors.New("id of Book Title is invalid"),
		"id of Book Title is invalid",
		"ErrBookTitleIdInvalid",
	)

	ErrBookTitleNameEmpty = common.NewCustomError(
		errors.New("name of Book Title is empty"),
		"name of Book Title is empty",
		"ErrBookTitleNameEmpty",
	)

	ErrBookTitlePublisherIdEmpty = common.NewCustomError(
		errors.New("publisher ID of Book Title is empty"),
		"publisher ID of Book Title is empty",
		"ErrBookTitlePublisherIdEmpty",
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

	ErrBookTitleValidatePublisher = common.NewCustomError(
		errors.New("publisher ID of Book Title is invalid"),
		"publisher ID of Book Title is invalid",
		"ErrBookTitleValidatePublisher",
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

	ErrBookTitleCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create Book Title"),
	)
	ErrBookTitleViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view Book Title"),
	)
	ErrBookTitleUpdateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update Book Title"),
	)
	ErrBookTitleDeleteNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to delete Book Title"),
	)
)
