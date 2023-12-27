package categorymodel

import "book-store-management-backend/common"

type ResListCategory struct {
	Data   []Category    `json:"data"`
	Paging common.Paging `json:"paging"`
	Filter Filter        `json:"filter"`
}
