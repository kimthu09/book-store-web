package invoicemodel

type ResGetNearestInvoice struct {
	// Data contains list of invoice.
	Data []Invoice `json:"data"`
}
