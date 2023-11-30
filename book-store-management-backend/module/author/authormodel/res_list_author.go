package authormodel

import "book-store-management-backend/common"

type ResListAuthor struct {
	Data   []Author      `json:"data"`
	Paging common.Paging `json:"paging"`
	Filter Filter        `json:"filter"`
}
