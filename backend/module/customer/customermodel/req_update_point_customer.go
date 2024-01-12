package customermodel

import "book-store-management-backend/common"

type CustomerUpdatePoint struct {
	Amount *int `json:"amount" gorm:"-"`
}

func (*CustomerUpdatePoint) TableName() string {
	return common.TableCustomer
}
