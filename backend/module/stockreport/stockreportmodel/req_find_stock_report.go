package stockreportmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/stockreportdetail/stockreportdetailmodel"
	"time"
)

type ReqFindStockReport struct {
	Id           string                                           `json:"-" gorm:"column:id;"`
	TimeFrom     int64                                            `json:"timeFrom" gorm:"-" example:"1701388800"`
	TimeTo       int64                                            `json:"timeTo" gorm:"-" example:"1704067199"`
	TimeFromTime time.Time                                        `json:"-" gorm:"column:timeFrom"`
	TimeToTime   time.Time                                        `json:"-" gorm:"column:timeTo"`
	Details      []stockreportdetailmodel.StockReportDetailCreate `json:"-" gorm:"foreignkey:ReportId;association_foreignkey:id"`
}

func (*ReqFindStockReport) TableName() string {
	return common.TableStockReport
}

func (data *ReqFindStockReport) Validate() error {
	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)

	if timeFrom.After(timeTo) {
		return ErrStockReportDateInvalid
	}
	return nil
}
