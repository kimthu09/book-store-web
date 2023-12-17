package salereportbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"book-store-management-backend/module/salereport/salereportmodel"
	"book-store-management-backend/module/salereportdetail/salereportetailmodel"
	"context"
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

type findSaleReportBiz struct {
	invoiceStore       FindInvoiceForReportStore
	invoiceDetailStore ListInvoiceDetailForReportStore
	requester          middleware.Requester
}

func NewFindSaleReportBiz(
	invoiceStore FindInvoiceForReportStore,
	invoiceDetailStore ListInvoiceDetailForReportStore,
	requester middleware.Requester) *findSaleReportBiz {
	return &findSaleReportBiz{
		invoiceStore:       invoiceStore,
		invoiceDetailStore: invoiceDetailStore,
		requester:          requester,
	}
}

func (biz *findSaleReportBiz) FindSaleReport(
	ctx context.Context,
	data *salereportmodel.ReqFindSaleReport) (*salereportmodel.SaleReport, error) {
	if !biz.requester.IsHasFeature(common.SaleReportViewFeatureCode) {
		return nil, salereportmodel.ErrSaleReportViewNoPermission
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)
	data.TimeFromTime = timeFrom
	data.TimeToTime = timeTo

	allInvoices, errInvoices := biz.invoiceStore.ListAllInvoiceForReport(
		ctx, timeFrom, timeTo, "Details.Book")
	if errInvoices != nil {
		return nil, errInvoices
	}

	total := 0
	totalAmount := 0
	mapBookAmount := make(map[string]int)
	mapBookName := make(map[string]string)
	mapBookSales := make(map[string]int)
	for _, invoice := range allInvoices {
		details, err := biz.invoiceDetailStore.ListInvoiceDetail(
			ctx, invoice.Id)
		if err != nil {
			return nil, err
		}

		for _, detail := range details {
			mapBookAmount[detail.BookId] += detail.Quantity
			mapBookName[detail.BookId] = detail.Book.Name
			totalInvoiceDetail := detail.UnitPrice * detail.Quantity
			mapBookSales[detail.BookId] += totalInvoiceDetail

			total += totalInvoiceDetail
			totalAmount += detail.Quantity
		}
	}

	details := make([]salereportetailmodel.SaleReportDetail, 0)
	for key, value := range mapBookName {
		if mapBookAmount[key] != 0 {
			detail := salereportetailmodel.SaleReportDetail{
				Book: bookmodel.SimpleBook{
					ID:   key,
					Name: value,
				},
				Amount:     mapBookAmount[key],
				TotalSales: mapBookSales[key],
			}
			details = append(details, detail)
		}
	}

	saleReport := salereportmodel.SaleReport{
		TimeFrom: timeFrom,
		TimeTo:   timeTo,
		Total:    total,
		Amount:   totalAmount,
		Details:  details,
	}

	return &saleReport, nil
}
