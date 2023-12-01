package publishermodel

import "book-store-management-backend/common"

type ResListPublisher struct {
	Data   []Publisher   `json:"data"`
	Paging common.Paging `json:"paging"`
	Filter Filter        `json:"filter"`
}
