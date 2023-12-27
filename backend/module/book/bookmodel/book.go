package bookmodel

import (
	"book-store-management-backend/common"
	"time"
)

type Book struct {
	ID          *string `json:"id" gorm:"column:id"`
	Name        *string `json:"name" gorm:"column:name"`
	BookTitleID *string `json:"bookTitleId" gorm:"column:booktitleid,fk"`
	PublisherID *string `json:"publisherId" gorm:"column:publisherid,fk"`
	Image       *string `json:"image" gorm:"column:imgUrl"`
	Edition     *int    `json:"edition" gorm:"column:edition"`
	Quantity    *int    `json:"quantity" gorm:"column:quantity"`
	ListedPrice *int    `json:"listedPrice" gorm:"column:listedPrice"`
	SellPrice   *int    `json:"sellPrice" gorm:"column:sellPrice"`
	ImportPrice *int    `json:"importPrice" gorm:"column:importPrice"`

	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"createdAt; column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"updatedAt; column:updatedAt;"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"deletedAt; column:deletedAt;"`
	IsActive  *bool      `json:"isActive,omitempty" gorm:"isActive; column:isActive; default:1"`
}

func (*Book) TableName() string {
	return common.TableBook
}
