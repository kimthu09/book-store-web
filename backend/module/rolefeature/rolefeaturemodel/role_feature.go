package rolefeaturemodel

import (
	"book-store-management-backend/common"
)

type RoleFeature struct {
	RoleId    string `json:"roleId" gorm:"column:roleId;" example:"role id"`
	FeatureId string `json:"featureId" gorm:"column:featureId;" example:"feature id"`
}

func (*RoleFeature) TableName() string {
	return common.TableRoleFeature
}
