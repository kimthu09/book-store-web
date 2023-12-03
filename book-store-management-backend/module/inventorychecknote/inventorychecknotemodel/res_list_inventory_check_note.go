package inventorychecknotemodel

import (
	"book-store-management-backend/common"
)

type ResListInventoryCheckNote struct {
	// Data contains list of inventory check note.
	Data []InventoryCheckNote `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve inventory check note.
	Filter Filter `json:"filter,omitempty"`
}
