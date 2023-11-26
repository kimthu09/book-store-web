package categorymodel

import "book-store-management-backend/common"

type CreateCategoryRequest struct {
	Name string `json:"name" json:"column:name;"`
}

func (*CreateCategoryRequest) TableName() string {
	return common.TableCategory
}
