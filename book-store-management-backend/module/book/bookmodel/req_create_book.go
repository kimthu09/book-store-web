package bookmodel

import "book-store-management-backend/common"

type CreateBookRequest struct {
	Id        *string `json:"id" gorm:"column:id;"`
	Name      string  `json:"name" gorm:"column:name;"`
	Quantity  int     `json:"quantity" gorm:"column:qty;"`
	Edition   int     `json:"edition" gorm:"column:edition;"`
	Price     float64 `json:"price" gorm:"column:price;"`
	SalePrice float64 `json:"salePrice" gorm:"column:salePrice;"`
	IsActive  bool    `json:"isActive" gorm:"column:isActive;"`
}

func (*CreateBookRequest) TableName() string {
	return common.TableBook
}

func (data *CreateBookRequest) Validate() *common.AppError {
	//if !common.ValidateId(data.Id) {
	//	return ErrBookIdInvalid
	//}
	//if common.ValidateEmptyString(data.Name) {
	//	return ErrBookNameEmpty
	//}
	//if data.Price < 0 {
	//	return ErrBookPriceIsNegativeNumber
	//}
	return nil
}
