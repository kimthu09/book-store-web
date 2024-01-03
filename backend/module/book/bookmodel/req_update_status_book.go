package bookmodel

import "book-store-management-backend/common"

type ReqUpdateStatusBook struct {
	BookId   string `json:"bookId" gorm:"column:id;"`
	IsActive *bool  `json:"isActive" gorm:"column:isActive;"`
}

func (*ReqUpdateStatusBook) TableName() string {
	return common.TableBook
}
