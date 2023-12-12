package booktitlemodel

import "book-store-management-backend/common"

type ResListBookTitle struct {
	Data   []BookTitle   `json:"data"`
	Paging common.Paging `json:"paging,omitempty"`
	Filter Filter        `json:"filter,omitempty"`
}
