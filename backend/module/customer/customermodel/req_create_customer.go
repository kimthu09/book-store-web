package customermodel

import "book-store-management-backend/common"

type ReqCreateCustomer struct {
	Id    *string `json:"id" gorm:"column:id;" example:"customerId"`
	Name  string  `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
	Email string  `json:"email" gorm:"column:email;" example:"a@gmail.com"`
	Phone string  `json:"phone" gorm:"column:phone;" example:"0123456789"`
}

func (*ReqCreateCustomer) TableName() string {
	return common.TableCustomer
}

func (data *ReqCreateCustomer) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrCustomerIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrCustomerNameEmpty
	}
	if data.Email != "" && !common.ValidateEmail(data.Email) {
		return ErrCustomerEmailInvalid
	}
	if !common.ValidatePhone(data.Phone) {
		return ErrCustomerPhoneInvalid
	}
	return nil
}
