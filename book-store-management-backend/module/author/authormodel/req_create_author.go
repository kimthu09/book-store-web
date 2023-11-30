package authormodel

import "book-store-management-backend/common"

type ReqCreateAuthor struct {
	Name string `json:"name" json:"column:name;" example:"Nguyễn Nhật Ánh"`
}

func (*ReqCreateAuthor) TableName() string {
	return common.TableAuthor
}
