package categorymodel

import "book-store-management-backend/common"

type ReqCreateListCategory struct {
	Names []string `json:"names"`
}

func (*ReqCreateListCategory) TableName() string {
	return common.TableCategory
}
