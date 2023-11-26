package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"golang.org/x/net/context"
)

type BookDBModel struct {
	Id         *string `json:"id" gorm:"column:id;"`
	BookInfoId string  `json:"bookInfoId" gorm:"column:bookInfoId;"`
	Quantity   int     `json:"quantity" gorm:"column:qty;"`
	Edition    int     `json:"edition" gorm:"column:edition;"`
	Price      float64 `json:"price" gorm:"column:price;"`
	SalePrice  float64 `json:"salePrice" gorm:"column:salePrice;"`
	IsActive   bool    `json:"isActive" gorm:"column:isActive;"`
}

func (*BookDBModel) TableName() string {
	return "Book"
}

func (s *sqlStore) CreateBook(ctx context.Context, data *bookmodel.BookCreate) error {
	db := s.db

	var tmpData BookDBModel = BookDBModel{
		Id:         data.Id,
		BookInfoId: "JFK",
		Quantity:   data.Quantity,
		Edition:    data.Edition,
		Price:      data.Price,
		SalePrice:  data.SalePrice,
		IsActive:   data.IsActive,
	}
	if err := db.Create(tmpData).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
			case "PRIMARY":
				return bookmodel.ErrBookIdDuplicate
			}
		}
		return common.ErrDB(err)
	}

	//if err := db.Create(data).Error; err != nil {
	//	if gormErr := common.GetGormErr(err); gormErr != nil {
	//		switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
	//		case "PRIMARY":
	//			return bookmodel.ErrBookIdDuplicate
	//		}
	//	}
	//	return common.ErrDB(err)
	//}

	return nil
}
