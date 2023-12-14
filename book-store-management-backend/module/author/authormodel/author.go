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
	ErrAuthorNameEmpty = common.NewCustomError(
		errors.New("name of Author is empty"),
		"Tên tác giả đang trống",
		"ErrAuthorNameEmpty",
	)
	ErrAuthorIdDuplicate = common.ErrDuplicateKey(
		errors.New("Tác giả đã tồn tại"),
	)
	ErrAuthorNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên tác giả đã tồn tại"),
	)
)

func (data *Author) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrAuthorNameEmpty
	}
	return nil
}
