package invoicemodel

type ResCreateInvoice struct {
	// Data contains the detailed information about invoice details.
	Data ResDetailInvoice `json:"data"`
}
