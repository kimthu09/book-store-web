package authormodel

import (
	"book-store-management-backend/common"
	"errors"
	"time"
)

type Author struct {
	Id        string     `json:"id" json:"column:id;"`
	Name      string     `json:"name" json:"column:name;"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"createdAt; column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"deletedAt; column:deletedAt;"`
	IsActive  *bool      `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1"`
}

func (*Author) TableName() string {
	return common.TableAuthor
}

var (
	ErrAuthorIdInvalid = common.NewCustomError(
		errors.New("id of Author is invalid"),
		`id of Author is invalid`,
		"ErrAuthorIdInvalid",
	)
	ErrAuthorNameEmpty = common.NewCustomError(
		errors.New("name of Author is empty"),
		"name of Author is empty",
		"ErrAuthorNameEmpty",
	)
	ErrAuthorIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of Author is duplicate"),
	)
	ErrAuthorCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create Author"),
	)
	ErrAuthorViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view Author"),
	)
	ErrAuthorUpdateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update Author"),
	)
	ErrAuthorDeleteNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to delete Author"),
	)
)

func (data *Author) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrAuthorNameEmpty
	}
	return nil
}
