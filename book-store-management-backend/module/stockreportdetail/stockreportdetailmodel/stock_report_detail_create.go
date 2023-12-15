package stockreportdetailmodel

import (
	"book-store-management-backend/common"
)

type StockReportDetailCreate struct {
	ReportId string `json:"-" gorm:"column:reportId;"`
	BookId   string `json:"-" gorm:"column:bookId;"`
	Initial  int    `json:"-" gorm:"column:initial;"`
	Sell     int    `json:"-" gorm:"column:sell"`
	Import   int    `json:"-" gorm:"column:import;"`
	Modify   int    `json:"-" gorm:"column:modify;"`
	Final    int    `json:"-" gorm:"column:final;"`
}

func (*StockReportDetailCreate) TableName() string {
	return common.TableStockReportDetail
}
