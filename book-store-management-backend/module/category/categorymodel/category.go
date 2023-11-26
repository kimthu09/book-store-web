package categorymodel

import (
	"book-store-management-backend/common"
	"errors"
)

type Category struct {
	Id   string `json:"id" json:"column:id;"`
	Name string `json:"name" json:"column:name;"`
}

func (*Category) TableName() string {
	return common.TableCategory
}

var (
	ErrCategoryIdInvalid = common.NewCustomError(
		errors.New("id of Category is invalid"),
		`id of Category is invalid`,
		"ErrCategoryIdInvalid",
	)
	ErrCategoryNameEmpty = common.NewCustomError(
		errors.New("name of Category is empty"),
		"name of Category is empty",
		"ErrCategoryNameEmpty",
	)
	ErrCategoryIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of Category is duplicate"),
	)
	ErrCategoryCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create Category"),
	)
	ErrCategoryViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view Category"),
	)
	ErrCategoryUpdateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update Category"),
	)
	ErrCategoryDeleteNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to delete Category"),
	)
)

func (data *Category) Validate() *common.AppError {
	//if !common.ValidateId(data.Id) {
	//	return ErrBookIdInvalid
	//}
	//if common.ValidateEmptyString(data.Name) {
	//	return ErrBookNameEmpty
	//}
	//if data.Price < 0 {
	//	return ErrBookPriceIsNegativeNumber
	//}
	return nil
}
