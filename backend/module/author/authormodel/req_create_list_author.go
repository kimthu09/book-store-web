package authormodel

import "book-store-management-backend/common"

type ReqCreateListAuthor struct {
	Names []string `json:"names"`
}

func (*ReqCreateListAuthor) TableName() string {
	return common.TableAuthor
}
