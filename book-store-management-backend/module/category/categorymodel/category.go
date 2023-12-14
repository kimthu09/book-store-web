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
	ErrCategoryNameEmpty = common.NewCustomError(
		errors.New("name of Category is empty"),
		"Tên danh mục đang trống",
		"ErrCategoryNameEmpty",
	)
	ErrCategoryIdDuplicate = common.ErrDuplicateKey(
		errors.New("Danh mục đã tồn tại"),
	)
	ErrCategoryNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên danh mục đã tồn tại"),
	)
)

func (data *Category) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrCategoryNameEmpty
	}
	return nil
}
