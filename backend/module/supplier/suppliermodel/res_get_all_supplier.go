package suppliermodel

import "book-store-management-backend/module/importnote/importnotemodel"

type ResGetAllSupplier struct {
	// Data contains list of suppliers.
	Data []importnotemodel.SimpleSupplier `json:"data"`
}
