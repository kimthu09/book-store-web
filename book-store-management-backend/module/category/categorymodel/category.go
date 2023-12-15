package categorymodel

import (
	"book-store-management-backend/common"
	"errors"
	"time"
)

type Category struct {
	Id        string     `json:"id" json:"column:id;"`
	Name      string     `json:"name" json:"column:name;"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"createdAt; column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"deletedAt; column:deletedAt;"`
	IsActive  *bool      `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1"`
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
	if common.ValidateEmptyString(data.Name) {
		return ErrCategoryNameEmpty
	}
	return nil
}
