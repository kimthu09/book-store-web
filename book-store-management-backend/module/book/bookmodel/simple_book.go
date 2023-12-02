package bookmodel

import "book-store-management-backend/common"

type SimpleBook struct {
	ID   string `json:"id" gorm:"column:id;primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}

func (*SimpleBook) TableName() string {
	return common.TableBook
}
