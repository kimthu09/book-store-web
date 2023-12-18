package booktitlemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/author/authormodel"
	"book-store-management-backend/module/category/categorymodel"
)

type SimpleBookTitle struct {
	ID          string                         `json:"id" gorm:"column:id;primaryKey" example:"book title id"`
	Name        string                         `json:"name" gorm:"column:name" example:"Cho tui 1 vé về tuổi thơ"`
	Description string                         `json:"desc" gorm:"column:desc" example:"Câu chuyên hay cảm động rớt nước mắt"`
	AuthorIDs   string                         `json:"-" gorm:"column:authorIds"`
	CategoryIDs string                         `json:"-" gorm:"column:categoryIds"`
	Authors     []authormodel.SimpleAuthor     `json:"authors" gorm:"-"`
	Categories  []categorymodel.SimpleCategory `json:"categories" gorm:"-"`
}

func (*SimpleBookTitle) TableName() string {
	return common.TableBookTitle
}
