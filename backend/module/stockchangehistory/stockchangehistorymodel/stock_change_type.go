package stockchangehistorymodel

type StockChangeType string

const (
	Sell   StockChangeType = "Sell"
	Import StockChangeType = "Import"
	Modify StockChangeType = "Modify"
)
