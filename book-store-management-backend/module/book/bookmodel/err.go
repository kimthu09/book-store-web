package bookmodel

import (
	"book-store-management-backend/common"
	"errors"
)

var (
	ErrBookIdInvalid = common.NewCustomError(
		errors.New("id of Book is invalid"),
		"id of Book is invalid",
		"ErrBookIdInvalid",
	)

	ErrBookNameEmpty = common.NewCustomError(
		errors.New("name of Book is empty"),
		"name of Book is empty",
		"ErrBookNameEmpty",
	)

	ErrBookListedPriceIsLessThanZero = common.NewCustomError(
		errors.New("listed price of Book is less than zero"),
		"listed price of Book must be greater than 0",
		"ErrBookListedPriceIsLessThanZero",
	)

	ErrBookSalePriceIsLessThanZero = common.NewCustomError(
		errors.New("sell price of Book is less than zero"),
		"sell price of Book must be greater than 0",
		"ErrBookSalePriceIsLessThanZero",
	)

	ErrBookQuantityIsNegativeNumber = common.NewCustomError(
		errors.New("quantity of Book is a negative number"),
		"quantity of Book is a negative number",
		"ErrBookQuantityIsNegativeNumber",
	)

	ErrBookEditionIsNegativeNumber = common.NewCustomError(
		errors.New("edition number of Book is a negative number"),
		"edition number of Book is a negative number",
		"ErrBookEditionIsNegativeNumber",
	)

	ErrBookPublisherIdEmpty = common.NewCustomError(
		errors.New("publisher ID of Book is empty"),
		"publisher ID of Book is empty",
		"ErrBookPublisherIdEmpty",
	)

	ErrBookAuthorIdsEmpty = common.NewCustomError(
		errors.New("author IDs of Book are empty"),
		"author IDs of Book are empty",
		"ErrBookAuthorIdsEmpty",
	)

	ErrBookCategoryIdsEmpty = common.NewCustomError(
		errors.New("category IDs of Book are empty"),
		"category IDs of Book are empty",
		"ErrBookCategoryIdsEmpty",
	)

	ErrBookQtyUpdateInvalid = common.NewCustomError(
		errors.New("quantity need to update for the Book is invalid"),
		"quantity need to update for the Book is invalid",
		"ErrBookQtyUpdateInvalid",
	)
	ErrBookIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of Book is duplicate"),
	)
	ErrBookCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create Book"),
	)
	ErrBookViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view Book"),
	)
)
