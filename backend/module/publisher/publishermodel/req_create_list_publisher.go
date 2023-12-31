package publishermodel

import "book-store-management-backend/common"

type ReqCreateListPublisher struct {
	Names []string `json:"names"`
}

func (*ReqCreateListPublisher) TableName() string {
	return common.TablePublisher
}
