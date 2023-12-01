package publishermodel

import (
	"book-store-management-backend/common"
	"errors"
)

type Publisher struct {
	Id   string `json:"id" json:"column:id;"`
	Name string `json:"name" json:"column:name;"`
}

func (*Publisher) TableName() string {
	return common.TablePublisher
}

var (
	ErrPublisherIdInvalid = common.NewCustomError(
		errors.New("id of Publisher is invalid"),
		`id of Publisher is invalid`,
		"ErrPublisherIdInvalid",
	)
	ErrPublisherNameEmpty = common.NewCustomError(
		errors.New("name of Publisher is empty"),
		"name of Publisher is empty",
		"ErrPublisherNameEmpty",
	)
	ErrPublisherIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of Publisher is duplicate"),
	)
	ErrPublisherCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create Publisher"),
	)
	ErrPublisherViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view Publisher"),
	)
	ErrPublisherUpdateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update Publisher"),
	)
	ErrPublisherDeleteNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to delete Publisher"),
	)
)

func (data *Publisher) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrPublisherNameEmpty
	}
	return nil
}
