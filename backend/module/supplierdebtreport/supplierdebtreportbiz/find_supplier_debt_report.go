package supplierdebtreportbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/common/enum"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/importnote/importnotemodel"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
	"book-store-management-backend/module/supplierdebtreport/supplierdebtreportmodel"
	"book-store-management-backend/module/supplierdebtreportdetail/supplierdebtreportdetailmodel"
	"context"
	"errors"
	"fmt"
	"time"
)

type FindSupplierStore interface {
	ListAllSupplier(
		ctx context.Context) ([]importnotemodel.SimpleSupplier, error)
}

type FindSupplierDebtStore interface {
	ListAllSupplierDebtForReport(
		ctx context.Context,
		supplierId string,
		timeFrom time.Time,
		timeTo time.Time) ([]supplierdebtmodel.SupplierDebt, error)
	GetNearlySupplierDebt(
		ctx context.Context,
		supplierId string,
		timeFrom time.Time) (*supplierdebtmodel.SupplierDebt, error)
}

type FindSupplierDebtReportStore interface {
	FindSupplierDebtReport(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*supplierdebtreportmodel.SupplierDebtReport, error)
	CreateSupplierDebtReport(
		ctx context.Context,
		data *supplierdebtreportmodel.ReqFindSupplierDebtReport) error
}

type findSupplierDebtReportBiz struct {
	gen                     generator.IdGenerator
	supplierStore           FindSupplierStore
	supplierDebtStore       FindSupplierDebtStore
	supplierDebtReportStore FindSupplierDebtReportStore
	requester               middleware.Requester
}

func NewFindSupplierDebtReportBiz(
	gen generator.IdGenerator,
	supplierStore FindSupplierStore,
	supplierDebtStore FindSupplierDebtStore,
	supplierDebtReportStore FindSupplierDebtReportStore,
	requester middleware.Requester) *findSupplierDebtReportBiz {
	return &findSupplierDebtReportBiz{
		gen:                     gen,
		supplierStore:           supplierStore,
		supplierDebtStore:       supplierDebtStore,
		supplierDebtReportStore: supplierDebtReportStore,
		requester:               requester,
	}
}

func (biz *findSupplierDebtReportBiz) FindSupplierDebtReport(
	ctx context.Context,
	data *supplierdebtreportmodel.ReqFindSupplierDebtReport) (*supplierdebtreportmodel.SupplierDebtReport, error) {
	if !biz.requester.IsHasFeature(common.SupplierDebtReportViewFeatureCode) {
		return nil, supplierdebtreportmodel.ErrSupplierDebtReportViewNoPermission
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)
	data.TimeFromTime = timeFrom
	data.TimeToTime = timeTo

	report, err := biz.supplierDebtReportStore.FindSupplierDebtReport(
		ctx, map[string]interface{}{
			"timeFrom": data.TimeFromTime, "timeTo": data.TimeToTime,
		}, "Details.Supplier",
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

	allSupplier, err := biz.supplierStore.ListAllSupplier(ctx)
	if err != nil {
		return nil, err
	}

	reportId := ""

	now := time.Now()
	if now.Before(timeFrom) {
		return nil, supplierdebtreportmodel.ErrSupplierDebtReportDateIsInFuture
	}

	if timeTo.Before(now) {
		id, err := biz.gen.GenerateId()
		if err != nil {
			return nil, err
		}
		reportId = id
	}

	allDetails := make([]supplierdebtreportdetailmodel.SupplierDebtReportDetail, 0)
	allDetailCreates := make([]supplierdebtreportdetailmodel.ReqCreateSupplierDebtReportDetail, 0)
	totalInitial := 0
	totalDebt := 0
	totalPay := 0
	totalFinal := 0
	for _, supplier := range allSupplier {
		supplierDebts, err := biz.supplierDebtStore.ListAllSupplierDebtForReport(
			ctx, supplier.Id, timeFrom, timeTo)
		if err != nil {
			return nil, err
		}

		debtAmount := 0
		payAmount := 0
		for _, supplierDebt := range supplierDebts {
			if *supplierDebt.DebtType == enum.Debt {
				debtAmount += supplierDebt.Quantity
			} else if *supplierDebt.DebtType == enum.Pay {
				payAmount += supplierDebt.Quantity
			}
		}
		fmt.Println(debtAmount)

		initial := 0
		if nearly, err := biz.supplierDebtStore.GetNearlySupplierDebt(
			ctx, supplier.Id, timeFrom); err != nil {
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
		if len(supplierDebts) != 0 {
			final = supplierDebts[len(supplierDebts)-1].QuantityLeft
		}

		if initial == 0 {
			initial = final - debtAmount - payAmount
		}

		if initial != 0 || debtAmount != 0 || payAmount != 0 {
			detailCreate := supplierdebtreportdetailmodel.ReqCreateSupplierDebtReportDetail{
				ReportId:   reportId,
				SupplierId: supplier.Id,
				Initial:    initial,
				Debt:       debtAmount,
				Pay:        payAmount,
				Final:      final,
			}
			allDetailCreates = append(allDetailCreates, detailCreate)

			detail := supplierdebtreportdetailmodel.SupplierDebtReportDetail{
				ReportId:   reportId,
				SupplierId: supplier.Id,
				Supplier: importnotemodel.SimpleSupplier{
					Id:    supplier.Id,
					Name:  supplier.Name,
					Phone: supplier.Phone,
				},
				Initial: initial,
				Debt:    debtAmount,
				Pay:     payAmount,
				Final:   final,
			}
			allDetails = append(allDetails, detail)
		}

		totalInitial += initial
		totalDebt += debtAmount
		totalPay += payAmount
		totalFinal += final
	}

	data.Id = reportId
	data.Details = allDetailCreates
	data.Initial = totalInitial
	data.Debt = totalDebt
	data.Pay = totalPay
	data.Final = totalFinal
	if reportId != "" {
		if err := biz.supplierDebtReportStore.CreateSupplierDebtReport(
			ctx, data,
		); err != nil {
			return nil, err
		}
	}

	reportDebt := supplierdebtreportmodel.SupplierDebtReport{
		Id:       reportId,
		TimeFrom: timeFrom,
		TimeTo:   timeTo,
		Initial:  totalInitial,
		Debt:     totalDebt,
		Pay:      totalPay,
		Final:    totalFinal,
		Details:  allDetails,
	}

	return &reportDebt, nil
}
