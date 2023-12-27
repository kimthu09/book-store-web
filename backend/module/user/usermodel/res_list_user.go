package usermodel

import (
	"book-store-management-backend/common"
)

type ResListUser struct {
	// Data contains list of user.
	Data []ResUser `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve user.
	Filter Filter `json:"filter,omitempty"`
}
