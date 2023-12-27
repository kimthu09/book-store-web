package importnotemodel

import (
	"book-store-management-backend/common"
)

type ResListImportNote struct {
	// Data contains list of import note.
	Data []ImportNote `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve import note.
	Filter Filter `json:"filter,omitempty"`
}
