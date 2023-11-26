package rolefeaturemodel

import (
	"book-store-management-backend/common"
	"errors"
)

type RoleFeature struct {
	RoleId    string `json:"roleId" gorm:"column:roleId;"`
	FeatureId string `json:"featureId" gorm:"column:featureId;"`
}

func (*RoleFeature) TableName() string {
	return common.TableRoleFeature
}

var (
	ErrRoleFeatureIdFeatureInvalid = common.NewCustomError(
		errors.New("id of feature is invalid"),
		"id of feature is invalid",
		"ErrRoleFeatureIdFeatureInvalid",
	)
)
