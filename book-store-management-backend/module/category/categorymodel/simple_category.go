package categorymodel

import (
	"book-store-management-backend/common"
)

type SimpleCategory struct {
	Id   string `json:"id" json:"column:id;" example:"category id"`
	Name string `json:"name" json:"column:name;" example:"Tiểu thuyết"`
}

func (*SimpleCategory) TableName() string {
	return common.TableCategory
}
