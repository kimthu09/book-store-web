package suppliermodel

import "book-store-management-backend/common"

type ReqCreateSupplier struct {
	Id    *string `json:"id" gorm:"column:id;" example:"123"`
	Name  string  `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
	Email string  `json:"email" gorm:"column:email;" example:"a@gmail.com"`
	Phone string  `json:"phone" gorm:"column:phone;" example:"0123456789"`
	Debt  int     `json:"debt" gorm:"column:debt" example:"-100000"`
}

func (*ReqCreateSupplier) TableName() string {
	return common.TableSupplier
}

func (data *ReqCreateSupplier) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrSupplierIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrSupplierNameEmpty
	}
	if data.Email != "" && !common.ValidateEmail(data.Email) {
		return ErrSupplierEmailInvalid
	}
	if !common.ValidatePhone(data.Phone) {
		return ErrSupplierPhoneInvalid
	}
	if common.ValidateNegativeNumber(data.Debt) {
		return ErrSupplierInitDebtInvalid
	}
	return nil
}
