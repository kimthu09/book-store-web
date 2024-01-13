package customermodel

import "book-store-management-backend/common"

type ReqUpdateInfoCustomer struct {
	Name  *string `json:"name" gorm:"column:name;" example:"Nếu không sửa tên thì xóa trường này"`
	Email *string `json:"email" gorm:"column:email;" example:"Nếu không sửa email thì xóa trường này"`
	Phone *string `json:"phone" gorm:"column:phone;" example:"Nếu không sửa số điện thoại thì xóa trường này"`
}

func (*ReqUpdateInfoCustomer) TableName() string {
	return common.TableCustomer
}

func (data *ReqUpdateInfoCustomer) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrCustomerNameEmpty
	}
	if data.Email != nil && *data.Email != "" && !common.ValidateEmail(*data.Email) {
		return ErrCustomerEmailInvalid
	}
	if data.Phone != nil && !common.ValidatePhone(*data.Phone) {
		return ErrCustomerPhoneInvalid
	}
	return nil
}
