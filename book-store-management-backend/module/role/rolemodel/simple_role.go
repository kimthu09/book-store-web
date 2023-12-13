package rolemodel

import "book-store-management-backend/common"

type SimpleRole struct {
	Id   string `json:"id" gorm:"column:id;" example:"role id"`
	Name string `json:"name" gorm:"column:name;" example:"admin"`
}

func (*SimpleRole) TableName() string {
	return common.TableRole
}
