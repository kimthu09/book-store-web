package featuremodel

import (
	"book-store-management-backend/common"
	"errors"
)

type Feature struct {
	Id          string `json:"id" gorm:"column:id;" example:"feature id"`
	Description string `json:"description" gorm:"column:description;" example:"Xem nhân viên"`
}

func (*Feature) TableName() string {
	return common.TableFeature
}

var (
	ErrFeatureViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view feature"),
	)
)
