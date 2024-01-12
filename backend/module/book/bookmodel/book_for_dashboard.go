package bookmodel

import "book-store-management-backend/common"

type BookForDashboard struct {
	Id   string `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Qty  int    `json:"qty" gorm:"-"`
	Sale int    `json:"sale" gorm:"-"`
}

func (*BookForDashboard) TableName() string {
	return common.TableBook
}
