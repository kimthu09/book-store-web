package booktitlestore

import (
	"time"
)

type BookTitleDBModel struct {
	ID          *string    `json:"id" gorm:"column:id;primaryKey"`
	Name        *string    `json:"name" gorm:"column:name"`
	Description *string    `json:"desc" gorm:"column:desc"`
	AuthorIDs   *string    `json:"authorIds" gorm:"column:authorIds"`
	CategoryIDs *string    `json:"categoryIds" gorm:"column:categoryIds"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"createdAt; column:createdAt;"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt   *time.Time `json:"deletedAt" gorm:"deletedAt; column:deletedAt;"`
	IsActive    *bool      `json:"isActive" gorm:"isActive; column:isActive; default:1"`
}
