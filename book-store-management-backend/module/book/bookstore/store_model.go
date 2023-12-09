package bookstore

import "book-store-management-backend/common"

type BookDBModel struct {
	ID          *string `json:"id" gorm:"column:id;primaryKey"`
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"desc" gorm:"column:desc"`
	Edition     int     `json:"edition" gorm:"column:edition"`
	Quantity    int     `json:"quantity" gorm:"column:qty"`
	ListedPrice float64 `json:"listedPrice" gorm:"column:listedPrice"`
	SellPrice   float64 `json:"sellPrice" gorm:"column:sellPrice"`
	PublisherID string  `json:"publisherId" gorm:"column:publisherId"`
	AuthorIDs   string  `json:"authorIds" gorm:"column:authorIds"`
	CategoryIDs string  `json:"categoryIds" gorm:"column:categoryIds"`
	common.SQLModel
}
