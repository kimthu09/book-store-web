package bookmodel

import (
	"book-store-management-backend/common"
	"errors"
)

type Book struct {
	Id        string  `json:"id" json:"column:id;"`
	Quantity  int     `json:"quantity" json:"column:qty;"`
	Edition   int     `json:"edition" json:"column:edition;"`
	Price     float64 `json:"price" json:"column:price;"`
	SalePrice float64 `json:"salePrice" json:"column:salePrice;"`
	IsActive  bool    `json:"isActive" json:"column:isActive;"`
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
