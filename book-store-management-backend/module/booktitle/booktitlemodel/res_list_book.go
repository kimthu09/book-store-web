package booktitlemodel

import "book-store-management-backend/common"

type ResListBook struct {
	Data   []Book        `json:"data"`
	Paging common.Paging `json:"paging,omitempty"`
	Filter Filter        `json:"filter,omitempty"`
}
