package bookmodel

import (
	"book-store-management-backend/common"
	"errors"
)

type Book struct {
	Id     string  `json:"id" json:"column:id;"`
	Amount float32 `json:"amount" json:"column:amount;"`
}

func (*Book) TableName() string {
	return common.TableBook
}

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
	ErrBookPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price of Book is negative number"),
		"price of Book is negative number",
		"ErrBookPriceIsNegativeNumber",
	)
	ErrBookMeasureTypeEmpty = common.NewCustomError(
		errors.New("measure type of Book is empty"),
		"measure type of Book is empty",
		"ErrBookMeasureTypeEmpty",
	)
	ErrBookAmountUpdateInvalid = common.NewCustomError(
		errors.New("amount need to update for the Book is invalid"),
		"amount need to update for the Book is invalid",
		"ErrBookAmountUpdateInvalid",
	)
	ErrBookIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of Book is duplicate"),
	)
	ErrBookNameDuplicate = common.ErrDuplicateKey(
		errors.New("name of Book is duplicate"),
	)
	ErrBookCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create Book"),
	)
	ErrBookViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view Book"),
	)
)
