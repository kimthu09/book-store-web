package bookstore

import (
	"time"
)

type BookDBModel struct {
	ID          *string    `json:"id" gorm:"column:id;primaryKey"`
	Name        *string    `json:"name" gorm:"column:name"`
	Description *string    `json:"desc" gorm:"column:desc"`
	Edition     *int       `json:"edition" gorm:"column:edition"`
	Quantity    *int       `json:"quantity" gorm:"column:qty"`
	ListedPrice *float64   `json:"listedPrice" gorm:"column:listedPrice"`
	SellPrice   *float64   `json:"sellPrice" gorm:"column:sellPrice"`
	PublisherID *string    `json:"publisherId" gorm:"column:publisherId"`
	AuthorIDs   *string    `json:"authorIds" gorm:"column:authorIds"`
	CategoryIDs *string    `json:"categoryIds" gorm:"column:categoryIds"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"createdAt; column:createdAt;"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt   *time.Time `json:"deletedAt" gorm:"deletedAt; column:deletedAt;"`
	IsActive    *int       `json:"isActive" gorm:"isActive; column:isActive; default:1"`
}
