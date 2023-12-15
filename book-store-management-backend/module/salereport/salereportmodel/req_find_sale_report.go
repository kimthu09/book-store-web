package salereportmodel

import (
	"book-store-management-backend/module/salereportdetail/salereportetailmodel"
	"log"
	"time"
)

type ReqFindSaleReport struct {
	TimeFrom     int64                                   `json:"timeFrom" gorm:"-" example:"1701388800"`
	TimeTo       int64                                   `json:"timeTo" gorm:"-" example:"1704067199"`
	TimeFromTime time.Time                               `json:"-" gorm:"column:timeFrom"`
	TimeToTime   time.Time                               `json:"-" gorm:"column:timeTo"`
	Total        int                                     `json:"-"`
	Details      []salereportetailmodel.SaleReportDetail `json:"-"`
}

func (data *ReqFindSaleReport) Validate() error {
	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)

	log.Println(timeFrom)
	log.Println(timeTo)

	if timeFrom.After(timeTo) {
		return ErrSaleReportDateInvalid
	}
	return nil
}
