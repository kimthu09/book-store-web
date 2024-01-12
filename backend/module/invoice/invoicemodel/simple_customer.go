package invoicemodel

import "book-store-management-backend/common"

type SimpleCustomer struct {
	Id    string `json:"id" gorm:"column:id;" example:"customerId"`
	Name  string `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
	Phone string `json:"phone" gorm:"column:phone;" example:"0123456789"`
	Point int    `json:"point" gorm:"column:point;" example:"123"`
}

func (*SimpleCustomer) TableName() string {
	return common.TableCustomer
}
