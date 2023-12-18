package authormodel

import "book-store-management-backend/common"

type SimpleAuthor struct {
	Id   string `json:"id" json:"column:id;" example:"author id"`
	Name string `json:"name" json:"column:name;" example:"Nguyễn Văn A"`
}

func (*SimpleAuthor) TableName() string {
	return common.TableAuthor
}
