package rolemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/feature/featuremodel"
)

type ResSeeDetailRole struct {
	Id   string `json:"id" gorm:"column:id;" example:"role id"`
	Name string `json:"name" gorm:"column:name;" example:"admin"`
	// Data contains the detailed information about features.
	Data []featuremodel.ResFeatureDetail `json:"data" gorm:"-"`
}

func (*ResSeeDetailRole) TableName() string {
	return common.TableRole
}
