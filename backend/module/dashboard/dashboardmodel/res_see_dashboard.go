package dashboardmodel

import (
	"book-store-management-backend/module/book/bookmodel"
	"time"
)

type ResSeeDashboard struct {
	TimeFrom              time.Time                    `json:"timeFrom" gorm:"-"`
	TimeTo                time.Time                    `json:"timeTo" gorm:"-"`
	TotalSale             int                          `json:"totalSale" gorm:"-"`
	TotalSold             int                          `json:"totalSold" gorm:"-"`
	TotalCustomer         int                          `json:"totalCustomer" gorm:"-"`
	TotalPoint            int                          `json:"totalPoint" gorm:"-"`
	TopSoldBooks          []bookmodel.BookForDashboard `json:"topSoldBooks" gorm:"-"`
	ChartPriceComponents  []ChartComponent             `json:"chartPriceComponents" gorm:"-"`
	ChartProfitComponents []ChartComponent             `json:"chartProfitComponents" gorm:"-"`
}

type ChartComponent struct {
	Time  time.Time `json:"time" gorm:"-"`
	Value int       `json:"value" gorm:"-"`
}
