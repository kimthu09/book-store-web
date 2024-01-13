package invoicerepo

import (
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"context"
)

type SeeInvoiceDetailStore interface {
	ListInvoiceDetail(
		ctx context.Context,
		invoiceId string) ([]invoicedetailmodel.InvoiceDetail, error)
}

type FindInvoiceStore interface {
	FindInvoice(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*invoicemodel.Invoice, error)
}

type seeInvoiceDetailRepo struct {
	invoiceDetailStore SeeInvoiceDetailStore
	invoiceStore       FindInvoiceStore
}

func NewSeeInvoiceDetailRepo(
	invoiceDetailStore SeeInvoiceDetailStore,
	invoiceStore FindInvoiceStore) *seeInvoiceDetailRepo {
	return &seeInvoiceDetailRepo{
		invoiceDetailStore: invoiceDetailStore,
		invoiceStore:       invoiceStore,
	}
}

func (biz *seeInvoiceDetailRepo) SeeInvoiceDetail(
	ctx context.Context,
	invoiceId string) (*invoicemodel.Invoice, error) {
	invoice, errInvoice := biz.invoiceStore.FindInvoice(
		ctx,
		map[string]interface{}{
			"id": invoiceId,
		},
		"Customer", "CreatedByUser")
	if errInvoice != nil {
		return nil, errInvoice
	}

	details, errInvoiceDetail := biz.invoiceDetailStore.ListInvoiceDetail(
		ctx,
		invoiceId,
	)
	if errInvoiceDetail != nil {
		return nil, errInvoiceDetail
	}

	invoice.Details = details

	return invoice, nil
}
