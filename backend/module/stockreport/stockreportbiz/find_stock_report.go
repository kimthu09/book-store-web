package stockreportbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorymodel"
	"book-store-management-backend/module/stockreport/stockreportmodel"
	"book-store-management-backend/module/stockreportdetail/stockreportdetailmodel"
	"context"
	"errors"
	"time"
)

type FindBookStore interface {
	ListAllBook(
		ctx context.Context) ([]bookmodel.SimpleBook, error)
}

type FindStockChangeHistoryStore interface {
	ListAllStockChangeForReport(
		ctx context.Context,
		bookId string,
		timeFrom time.Time,
		timeTo time.Time) ([]stockchangehistorymodel.StockChangeHistory, error)
	GetNearlyStockChangeHistory(
		ctx context.Context,
		bookId string,
		timeFrom time.Time) (*stockchangehistorymodel.StockChangeHistory, error)
}

type FindStockReportStore interface {
	FindStockReport(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*stockreportmodel.StockReport, error)
	CreateStockReport(
		ctx context.Context,
		data *stockreportmodel.ReqFindStockReport) error
}

type findStockReportBiz struct {
	gen                     generator.IdGenerator
	bookStore               FindBookStore
	stockChangeHistoryStore FindStockChangeHistoryStore
	inventoryReportStore    FindStockReportStore
	requester               middleware.Requester
}

func NewFindStockReportBiz(
	gen generator.IdGenerator,
	bookStore FindBookStore,
	stockChangeHistoryStore FindStockChangeHistoryStore,
	inventoryReportStore FindStockReportStore,
	requester middleware.Requester) *findStockReportBiz {
	return &findStockReportBiz{
		gen:                     gen,
		bookStore:               bookStore,
		stockChangeHistoryStore: stockChangeHistoryStore,
		inventoryReportStore:    inventoryReportStore,
		requester:               requester,
	}
}

func (biz *findStockReportBiz) FindStockReport(
	ctx context.Context,
	data *stockreportmodel.ReqFindStockReport,
) (*stockreportmodel.StockReport, error) {
	if !biz.requester.IsHasFeature(common.StockReportViewFeatureCode) {
		return nil, stockreportmodel.ErrStockReportViewNoPermission
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)
	data.TimeFromTime = timeFrom
	data.TimeToTime = timeTo

	report, err := biz.inventoryReportStore.FindStockReport(
		ctx, map[string]interface{}{
			"timeFrom": data.TimeFromTime, "timeTo": data.TimeToTime,
		}, "Details.Book",
	)
	if err == nil {
		return report, nil
	} else {
		var appErr *common.AppError
		if errors.As(err, &appErr) {
			if appErr.Key != common.ErrRecordNotFound().Key {
				return nil, err
			}
		}
	}

	allBook, err := biz.bookStore.ListAllBook(ctx)
	if err != nil {
		return nil, err
	}

	reportId := ""

	now := time.Now()

	if now.Before(timeFrom) {
		return nil, stockreportmodel.ErrStockReportFutureDateInvalid
	}

	if timeTo.Before(now) {
		id, err := biz.gen.GenerateId()
		if err != nil {
			return nil, err
		}
		reportId = id
	}

	allDetailCreates := make([]stockreportdetailmodel.StockReportDetailCreate, 0)
	allDetails := make([]stockreportdetailmodel.StockReportDetail, 0)

	qtyInit := 0
	qtyImport := 0
	qtyModify := 0
	qtySell := 0
	qtyFinal := 0

	for _, book := range allBook {
		stockChange, err :=
			biz.stockChangeHistoryStore.ListAllStockChangeForReport(
				ctx, book.ID, timeFrom, timeTo)
		if err != nil {
			return nil, err
		}

		sellAmount := 0
		importAmount := 0
		modifyAmount := 0
		for _, change := range stockChange {
			if *change.Type == stockchangehistorymodel.Sell {
				sellAmount += change.Quantity
			} else if *change.Type == stockchangehistorymodel.Import {
				importAmount += change.Quantity
			} else if *change.Type == stockchangehistorymodel.Modify {
				modifyAmount += change.Quantity
			}
		}

		initial := 0
		if nearly, err :=
			biz.stockChangeHistoryStore.GetNearlyStockChangeHistory(
				ctx, book.ID, timeFrom,
			); err != nil {
			var appErr *common.AppError
			if errors.As(err, &appErr) {
				if appErr.Key != common.ErrRecordNotFound().Key {
					return nil, err
				}
			}
		} else {
			initial = nearly.QuantityLeft
		}

		final := initial
		if len(stockChange) != 0 {
			final = stockChange[len(stockChange)-1].QuantityLeft
		}

		if initial == 0 {
			initial = final - importAmount - sellAmount - modifyAmount
		}

		if initial != 0 || sellAmount != 0 ||
			importAmount != 0 || modifyAmount != 0 {
			detailCreate := stockreportdetailmodel.StockReportDetailCreate{
				ReportId: reportId,
				BookId:   book.ID,
				Initial:  initial,
				Sell:     sellAmount,
				Import:   importAmount,
				Modify:   modifyAmount,
				Final:    final,
			}
			allDetailCreates = append(allDetailCreates, detailCreate)

			detail := stockreportdetailmodel.StockReportDetail{
				ReportId: reportId,
				BookId:   book.ID,
				Book: bookmodel.SimpleBook{
					ID:   book.ID,
					Name: book.Name,
				},
				Initial: initial,
				Sell:    sellAmount,
				Import:  importAmount,
				Modify:  modifyAmount,
				Final:   final,
			}
			allDetails = append(allDetails, detail)
		}

		qtyInit += initial
		qtyImport += importAmount
		qtyModify += modifyAmount
		qtySell += sellAmount
		qtyFinal += final
	}

	data.Id = reportId
	data.Details = allDetailCreates
	data.Initial = qtyInit
	data.Import = qtyImport
	data.Modify = qtyModify
	data.Sell = qtySell
	data.Final = qtyFinal
	if reportId != "" {
		if err := biz.inventoryReportStore.CreateStockReport(
			ctx, data,
		); err != nil {
			return nil, err
		}
	}

	stockReport := stockreportmodel.StockReport{
		Id:       reportId,
		TimeFrom: timeFrom,
		TimeTo:   timeTo,
		Initial:  qtyInit,
		Sell:     qtySell,
		Import:   qtyImport,
		Modify:   qtyModify,
		Final:    qtyFinal,
		Details:  allDetails,
	}

	return &stockReport, nil
}
