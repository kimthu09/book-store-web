package bookmodel

import "book-store-management-backend/common"

type SimpleBook struct {
	ID   string `json:"id" gorm:"column:id;primaryKey" example:"book id"`
	Name string `json:"name" gorm:"column:name" example:"Những câu chuyện hay"`
}

func (*SimpleBook) TableName() string {
	return common.TableBook
}
