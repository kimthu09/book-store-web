package bookmodel

import (
	"book-store-management-backend/common"
	"errors"
)

var (
	ErrBookCreateNoPermission = common.NewCustomError(
		errors.New("no permission to create book"),
		"no permission to create book",
		"ErrBookCreateNoPermission",
	)
	ErrBookUpdateNoPermission = common.NewCustomError(
		errors.New("no permission to update book"),
		"no permission to update book",
		"ErrBookUpdateNoPermission",
	)
	ErrBookViewNoPermission = common.NewCustomError(
		errors.New("no permission to view book"),
		"no permission to view book",
		"ErrBookViewNoPermission",
	)
	ErrBookDeleteNoPermission = common.NewCustomError(
		errors.New("no permission to delete book"),
		"no permission to delete book",
		"ErrBookDeleteNoPermission",
	)

	ErrBookIdInvalid = common.NewCustomError(
		errors.New("id of Book is invalid"),
		"id of Book is invalid",
		"ErrBookIdInvalid",
	)
	ErrBookTitleIdInvalid = common.NewCustomError(
		errors.New("id of Book Title is invalid"),
		"id of Book Title is invalid",
		"ErrBookTitleIdInvalid",
	)
	ErrPublisherIdInvalid = common.NewCustomError(
		errors.New("id of Publisher is invalid"),
		"id of Publisher is invalid",
		"ErrPublisherIdInvalid",
	)
	ErrBookEditionInvalid = common.NewCustomError(
		errors.New("edition of Book is invalid"),
		"edition of Book is invalid",
		"ErrBookEditionInvalid",
	)
	ErrBookListedPriceInvalid = common.NewCustomError(
		errors.New("listed price of Book is invalid"),
		"listed price of Book is invalid",
		"ErrBookListedPriceInvalid",
	)
	ErrBookSellPriceInvalid = common.NewCustomError(
		errors.New("sell price of Book is invalid"),
		"sell price of Book is invalid",
		"ErrBookSellPriceInvalid",
	)

	ErrBookIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of Book is duplicate"),
	)
)
