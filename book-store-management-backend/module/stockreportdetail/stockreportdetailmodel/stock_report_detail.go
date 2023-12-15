package stockreportdetailmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
)

type StockReportDetail struct {
	ReportId string               `json:"-" gorm:"column:reportId;"`
	BookId   string               `json:"-" gorm:"column:bookId;"`
	Book     bookmodel.SimpleBook `json:"book"`
	Initial  int                  `json:"initial" gorm:"column:initial;" example:"0"`
	Sell     int                  `json:"sell" gorm:"column:sell" example:"-10"`
	Import   int                  `json:"import" gorm:"column:import;" example:"100"`
	Modify   int                  `json:"modify" gorm:"column:modify;" example:"-60"`
	Final    int                  `json:"final" gorm:"column:final;" example:"30"`
}

func (*StockReportDetail) TableName() string {
	return common.TableStockReportDetail
}
