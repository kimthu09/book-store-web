package bookmodel

import "book-store-management-backend/common"

type SimpleBook struct {
	ID   string `json:"id" gorm:"column:id;primaryKey" example:"idOfBook"`
	Name string `json:"name" gorm:"column:name" example:"Doraemon"`
}

func (*SimpleBook) TableName() string {
	return common.TableBook
}
