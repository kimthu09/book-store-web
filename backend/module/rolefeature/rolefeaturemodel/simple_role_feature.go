package rolefeaturemodel

import "book-store-management-backend/common"

type SimpleRoleFeature struct {
	RoleId    string `json:"-" gorm:"column:roleId;"`
	FeatureId string `json:"featureId" gorm:"column:featureId;" example:"feature id"`
}

func (*SimpleRoleFeature) TableName() string {
	return common.TableRoleFeature
}
