package suppliermodel

import "book-store-management-backend/common"

type ReqUpdateInfoSupplier struct {
	Name  *string `json:"name" gorm:"column:name;" example:"tên đã đổi"`
	Email *string `json:"email" gorm:"column:email;" example:"b@gmail.com"`
	Phone *string `json:"phone" gorm:"column:phone;" example:"1234567890"`
}

func (*ReqUpdateInfoSupplier) TableName() string {
	return common.TableSupplier
}

func (data *ReqUpdateInfoSupplier) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrSupplierNameEmpty
	}
	if data.Email != nil && *data.Email != "" && !common.ValidateEmail(*data.Email) {
		return ErrSupplierEmailInvalid
	}
	if data.Phone != nil && !common.ValidatePhone(*data.Phone) {
		return ErrSupplierPhoneInvalid
	}
	return nil
}
