package importnotemodel

import "book-store-management-backend/common"

type SimpleSupplier struct {
	Id    string `json:"id" gorm:"column:id;" example:"123"`
	Name  string `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
	Phone string `json:"phone" gorm:"column:phone;" example:"0123456789"`
}

func (*SimpleSupplier) TableName() string {
	return common.TableSupplier
}
