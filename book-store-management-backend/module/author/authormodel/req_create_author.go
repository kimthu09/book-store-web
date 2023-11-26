package authormodel

import "book-store-management-backend/common"

type CreateAuthorRequest struct {
	Name string `json:"name" json:"column:name;"`
}

func (*CreateAuthorRequest) TableName() string {
	return common.TableAuthor
}
