package bookmodel

import "book-store-management-backend/common"

type ReqListBook struct {
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve book.
	Filter Filter `json:"filter,omitempty"`
}
