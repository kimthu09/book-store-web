package invoicemodel

type ReqGetNearestInvoice struct {
	AmountNeed int `json:"amountNeed" form:"amountNeed"`
}
