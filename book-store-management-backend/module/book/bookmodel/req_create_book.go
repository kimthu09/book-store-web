package bookmodel

import (
	"book-store-management-backend/common"
)

type ReqCreateBook struct {
	Id          *string  `json:"id" gorm:"column:id"`
	Name        string   `json:"name" gorm:"column:name"`
	Quantity    int      `json:"quantity" gorm:"column:qty"`
	Edition     int      `json:"edition" gorm:"column:edition"`
	Price       float64  `json:"price" gorm:"column:price"`
	SalePrice   float64  `json:"salePrice" gorm:"column:salePrice"`
	AuthorIds   []string `json:"authorIds" gorm:"column:authorIds"`
	CategoryIds []string `json:"categoryIds" gorm:"column:categoryIds"`
	PublisherId string   `json:"publisherId" gorm:"column:publisherId"`
	Desc        string   `json:"desc" gorm:"column:desc"`
	IsActive    bool     `json:"isActive" gorm:"column:isActive"`
}

// func (*CreateBookRequest) TableName() string {
// 	return common.TableBook
// }

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

	if common.ValidateEmptyString(data.PublisherId) {
		return ErrBookPublisherIdEmpty
	}

	if len(data.AuthorIds) == 0 {
		return ErrBookAuthorIdsEmpty
	}

	if len(data.CategoryIds) == 0 {
		return ErrBookCategoryIdsEmpty
	}

	return nil
}
