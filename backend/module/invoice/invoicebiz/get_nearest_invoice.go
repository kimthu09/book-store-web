package invoicebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
)

type GetNearestInvoiceRepo interface {
	GetNearestInvoice(
		ctx context.Context,
		amountNeed int) ([]invoicemodel.Invoice, error)
}

type getNearestInvoiceBiz struct {
	repo      GetNearestInvoiceRepo
	requester middleware.Requester
}

func NewGetNearestBiz(
	repo GetNearestInvoiceRepo,
	requester middleware.Requester) *getNearestInvoiceBiz {
	return &getNearestInvoiceBiz{repo: repo, requester: requester}
}

func (biz *getNearestInvoiceBiz) GetNearestInvoice(
	ctx context.Context,
	amountNeed int) ([]invoicemodel.Invoice, error) {
	if !biz.requester.IsHasFeature(common.InvoiceViewFeatureCode) {
		return nil, invoicemodel.ErrInvoiceViewNoPermission
	}

	result, err := biz.repo.GetNearestInvoice(ctx, amountNeed)
	if err != nil {
		return nil, err
	}
	return result, nil
}
