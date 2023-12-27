package booktitlemodel

import (
	"book-store-management-backend/module/author/authormodel"
	"book-store-management-backend/module/category/categorymodel"
	"time"
)

type BookTitleDetail struct {
	ID          *string                  `json:"id" gorm:"column:id;primaryKey"`
	Name        string                   `json:"name" gorm:"column:name"`
	Description string                   `json:"desc" gorm:"column:desc"`
	Authors     []authormodel.Author     `json:"authors" gorm:"column:authorIds"`
	Categories  []categorymodel.Category `json:"categories" gorm:"column:categoryIds"`

	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"createdAt; column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"deletedAt; column:deletedAt;"`
	IsActive  *bool      `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1"`
}
