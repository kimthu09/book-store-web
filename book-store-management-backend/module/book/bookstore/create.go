package bookstore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"fmt"
	"gorm.io/gorm"

	"golang.org/x/net/context"
)

type BookDBModel struct {
	gorm.Model
	Id          *string `json:"id" gorm:"column:id;"`
	Quantity    int     `json:"quantity" gorm:"column:qty;"`
	Edition     int     `json:"edition" gorm:"column:edition;"`
	ListedPrice float64 `json:"listedPrice" gorm:"column:listedPrice"`
	SellPrice   float64 `json:"sellPrice" gorm:"column:sellPrice"`
	IsActive    bool    `json:"isActive" gorm:"column:isActive;"`
}

func (*BookDBModel) TableName() string {
	return common.TableBook
}

func (s *sqlStore) CreateBook(ctx context.Context, data *bookmodel.Book) error {
	fmt.Println("=====================================\nStore Book\n=====================================\n")
	return nil

	db := s.db

	var tmpData BookDBModel = BookDBModel{
		Quantity:    data.Quantity,
		Edition:     data.Edition,
		ListedPrice: data.ListedPrice,
		SellPrice:   data.SellPrice,
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
