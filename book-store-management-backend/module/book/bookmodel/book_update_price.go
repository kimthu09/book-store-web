package bookmodel

//
//import "book-store-management-backend/common"
//
//type BookUpdatePrice struct {
//	Price *float32 `json:"price" gorm:"column:price;"`
//}
//
//func (*BookUpdatePrice) TableName() string {
//	return common.TableBook
//}
//
//func (data *BookUpdatePrice) Validate() *common.AppError {
//	if data.Price != nil && common.ValidateNegativeNumber(*data.Price) {
//		return ErrBookPriceIsNegativeNumber
//	}
//	return nil
//}
