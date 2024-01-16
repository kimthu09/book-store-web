package dashboardrepo

import (
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/dashboard/dashboardmodel"
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"book-store-management-backend/module/salereport/salereportmodel"
	"book-store-management-backend/module/salereportdetail/salereportetailmodel"
	"context"
	"sort"
	"time"
)

type FindInvoiceForReportStore interface {
	ListAllInvoiceForReport(
		ctx context.Context,
		startTime time.Time,
		endTime time.Time,
		moreKeys ...string) ([]invoicemodel.Invoice, error)
}

type ListInvoiceDetailForReportStore interface {
	ListInvoiceDetail(
		ctx context.Context,
		invoiceId string) ([]invoicedetailmodel.InvoiceDetail, error)
}

type seeDashboardRepo struct {
	invoiceStore       FindInvoiceForReportStore
	invoiceDetailStore ListInvoiceDetailForReportStore
	requester          middleware.Requester
}

func NewSeeDashboardBiz(
	invoiceStore FindInvoiceForReportStore,
	invoiceDetailStore ListInvoiceDetailForReportStore) *seeDashboardRepo {
	return &seeDashboardRepo{
		invoiceStore:       invoiceStore,
		invoiceDetailStore: invoiceDetailStore,
	}
}

func (repo *seeDashboardRepo) SeeDashboard(
	ctx context.Context,
	data *dashboardmodel.ReqSeeDashboard) (*dashboardmodel.ResSeeDashboard, error) {
	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)

	resDashBoard := dashboardmodel.ResSeeDashboard{
		TimeFrom: timeFrom,
		TimeTo:   timeTo,
	}

	allInvoices, errInvoices := repo.invoiceStore.ListAllInvoiceForReport(
		ctx, timeFrom, timeTo, "Details.Book")
	if errInvoices != nil {
		return nil, errInvoices
	}

	total := 0
	totalQty := 0
	totalCustomer := 0
	totalPoint := 0
	mapBookQty := make(map[string]int)
	mapBookName := make(map[string]string)
	mapBookSales := make(map[string]int)
	listPrice := make([]dashboardmodel.ChartComponent, 0)
	listProfit := make([]dashboardmodel.ChartComponent, 0)
	for _, invoice := range allInvoices {
		chartSale := dashboardmodel.ChartComponent{
			Time:  *invoice.CreatedAt,
			Value: invoice.AmountReceived - invoice.TotalImportPrice,
		}
		listProfit = append(listProfit, chartSale)

		chartQtyReceive := dashboardmodel.ChartComponent{
			Time:  *invoice.CreatedAt,
			Value: invoice.TotalPrice,
		}
		listPrice = append(listPrice, chartQtyReceive)

		totalPoint += invoice.PointReceive

		details, err := repo.invoiceDetailStore.ListInvoiceDetail(
			ctx, invoice.Id)
		if err != nil {
			return nil, err
		}

		for _, detail := range details {
			mapBookQty[detail.BookId] += detail.Quantity
			mapBookName[detail.BookId] = detail.Book.Name
			totalInvoiceDetail := detail.UnitPrice * detail.Quantity

			mapBookSales[detail.BookId] += totalInvoiceDetail

			total += totalInvoiceDetail
			totalQty += detail.Quantity
		}

		if invoice.CustomerId != "" {
			totalCustomer++
		}
	}

	details := make(salereportmodel.Details, 0)
	for key, value := range mapBookName {
		if mapBookQty[key] != 0 {
			detail := salereportetailmodel.SaleReportDetail{
				Book: bookmodel.SimpleBook{
					ID:   key,
					Name: value,
				},
				Amount:     mapBookQty[key],
				TotalSales: mapBookSales[key],
			}
			details = append(details, detail)
		}
	}

	sort.Sort(details)

	listBook := make([]bookmodel.BookForDashboard, 0)
	for i, v := range details {
		if i >= dashboardmodel.NumberNearestInvoice {
			break
		}
		book := bookmodel.BookForDashboard{
			Id:   v.Book.ID,
			Name: v.Book.Name,
			Qty:  v.Amount,
			Sale: v.TotalSales,
		}
		listBook = append(listBook, book)
	}
	resDashBoard.TopSoldBooks = listBook

	resDashBoard.ChartProfitComponents = listProfit
	resDashBoard.ChartPriceComponents = listPrice
	resDashBoard.TotalSale = total
	resDashBoard.TotalCustomer = totalCustomer
	resDashBoard.TotalSold = totalQty
	resDashBoard.TotalPoint = totalPoint

	return &resDashBoard, nil
}
