package bookmodel

import (
	"book-store-management-backend/common"
)

type ReqCreateBook struct {
	Name        string   `json:"name" gorm:"column:name"`
	Description string   `json:"desc" gorm:"column:desc"`
	Edition     int      `json:"edition" gorm:"column:edition"`
	Quantity    int      `json:"quantity" gorm:"column:qty"`
	Price       float64  `json:"price" gorm:"column:price"`
	SalePrice   float64  `json:"salePrice" gorm:"column:salePrice"`
	PublisherID string   `json:"publisherId" gorm:"column:publisherId"`
	AuthorIDs   []string `json:"authorIds" gorm:"column:authorIds"`
	CategoryIDs []string `json:"categoryIds" gorm:"column:categoryIds"`
}

//func (*ReqCreateBook) TableName() string {
//	return common.TableBook
//}

func (data *ReqCreateBook) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrBookNameEmpty
	}

	if data.Price <= 0 {
		return ErrBookPriceIsLessThanZero
	}

	if data.SalePrice <= 0 {
		return ErrBookSalePriceIsLessThanZero
	}

	if data.Quantity < 0 {
		return ErrBookQuantityIsNegativeNumber
	}

	if data.Edition < 0 {
		return ErrBookEditionIsNegativeNumber
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
