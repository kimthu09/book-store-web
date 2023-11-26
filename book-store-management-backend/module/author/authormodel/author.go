package authormodel

import (
	"book-store-management-backend/common"
	"errors"
)

type Author struct {
	Id   string `json:"id" json:"column:id;"`
	Name string `json:"name" json:"column:name;"`
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
