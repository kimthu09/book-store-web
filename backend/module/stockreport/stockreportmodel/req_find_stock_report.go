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
	Initial      int                                              `json:"-" gorm:"column:initial;" example:"0"`
	Sell         int                                              `json:"-" gorm:"column:sell" example:"-10"`
	Import       int                                              `json:"-" gorm:"column:import;" example:"100"`
	Modify       int                                              `json:"-" gorm:"column:modify;" example:"-60"`
	Final        int                                              `json:"-" gorm:"column:final;" example:"30"`
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
