package publishermodel

import "book-store-management-backend/common"

type ReqCreatePublisher struct {
	Name string `json:"name" json:"column:name;" example:"Kim Đồng"`
}

func (*ReqCreatePublisher) TableName() string {
	return common.TablePublisher
}
