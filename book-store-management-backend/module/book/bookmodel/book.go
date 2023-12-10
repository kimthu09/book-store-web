package bookmodel

import (
	"book-store-management-backend/common"
)

type Book struct {
	ID          *string  `json:"id" gorm:"column:id;primaryKey"`
	Name        string   `json:"name" gorm:"column:name"`
	Description string   `json:"desc" gorm:"column:desc"`
	Edition     int      `json:"edition" gorm:"column:edition"`
	Quantity    int      `json:"quantity" gorm:"column:qty"`
	ListedPrice float64  `json:"listedPrice" gorm:"column:listedPrice"`
	SellPrice   float64  `json:"sellPrice" gorm:"column:sellPrice"`
	PublisherID string   `json:"publisherId" gorm:"column:publisherId"`
	AuthorIDs   []string `json:"authorIds" gorm:"column:authorIds"`
	CategoryIDs []string `json:"categoryIds" gorm:"column:categoryIds"`
	common.SQLModel
}

func (*Book) TableName() string {
	return common.TableBook
}

func (data *Book) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrBookNameEmpty
	}

	if data.ListedPrice <= 0 {
		return ErrBookListedPriceIsLessThanZero
	}

	if data.SellPrice <= 0 {
		if data.SellPrice == 0 {
			data.SellPrice = data.ListedPrice
		} else {
			return ErrBookSalePriceIsLessThanZero
		}
	}

	if data.Quantity < 0 {
		return ErrBookQuantityIsNegativeNumber
	}

	if data.Edition <= 0 {
		return ErrBookEditionNotPositiveNumber
	}

	if common.ValidateEmptyString(data.PublisherID) {
		return ErrBookPublisherIdEmpty
	}

	if len(data.AuthorIDs) == 0 {
		return ErrBookAuthorIdsEmpty
	}

	if len(data.CategoryIDs) == 0 {
		return ErrBookCategoryIdsEmpty
	}

	return nil
}
