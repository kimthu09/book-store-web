package categorymodel

import "book-store-management-backend/common"

type ReqCreateCategory struct {
	Name string `json:"name" json:"column:name;" example:"Trinh thám"`
}

func (*ReqCreateCategory) TableName() string {
	return common.TableCategory
}
