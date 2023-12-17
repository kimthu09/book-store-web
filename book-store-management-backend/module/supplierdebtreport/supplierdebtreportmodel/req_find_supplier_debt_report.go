package supplierdebtreportmodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplierdebtreportdetail/supplierdebtreportdetailmodel"
	"time"
)

type ReqFindSupplierDebtReport struct {
	Id           string                                                            `json:"-" gorm:"column:id;"`
	TimeFrom     int64                                                             `json:"timeFrom" gorm:"-" example:"1701388800"`
	TimeTo       int64                                                             `json:"timeTo" gorm:"-" example:"1704067199"`
	Initial      int                                                               `json:"-" gorm:"column:initial" example:"-100000"`
	Debt         int                                                               `json:"-" gorm:"column:debt" example:"-40000"`
	Pay          int                                                               `json:"-" gorm:"column:pay" example:"20000"`
	Final        int                                                               `json:"-" gorm:"column:final" example:"-120000"`
	TimeFromTime time.Time                                                         `json:"-" gorm:"column:timeFrom"`
	TimeToTime   time.Time                                                         `json:"-" gorm:"column:timeTo"`
	Details      []supplierdebtreportdetailmodel.ReqCreateSupplierDebtReportDetail `json:"-" gorm:"foreignkey:ReportId;association_foreignkey:id"`
}

func (*ReqFindSupplierDebtReport) TableName() string {
	return common.TableSupplierDebtReport
}

func (data *ReqFindSupplierDebtReport) Validate() error {
	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)

	if timeFrom.After(timeTo) {
		return ErrSupplierDebtReportDateInvalid
	}
	return nil
}
