package salereportetailmodel

import "book-store-management-backend/module/book/bookmodel"

type SaleReportDetail struct {
	BookId     string               `json:"-"`
	Book       bookmodel.SimpleBook `json:"book"`
	Amount     int                  `json:"amount" example:"10"`
	TotalSales int                  `json:"totalSales" example:"100000"`
}
