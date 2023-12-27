package featuremodel

import (
	"book-store-management-backend/common"
	"errors"
)

type Feature struct {
	Id          string `json:"id" gorm:"column:id;" example:"feature id"`
	Description string `json:"description" gorm:"column:description;" example:"Xem nhân viên"`
	GroupName   string `json:"groupName" gorm:"column:groupName" example:"Nhân viên"`
}

func (*Feature) TableName() string {
	return common.TableFeature
}

var (
	ErrFeatureViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem các chức năng"),
	)
)
