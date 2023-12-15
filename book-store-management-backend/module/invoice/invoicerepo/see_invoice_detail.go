package invoicerepo

import (
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"context"
)

type FindInvoiceStore interface {
	FindInvoice(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*invoicemodel.Invoice, error)
}

type ListInvoiceDetailStore interface {
	ListInvoiceDetail(
		ctx context.Context,
		invoiceId string,
	) ([]invoicedetailmodel.InvoiceDetail, error)
}

type seeInvoiceDetailRepo struct {
	invoiceStore       FindInvoiceStore
	invoiceDetailStore ListInvoiceDetailStore
}

func NewSeeInvoiceDetailRepo(
	invoiceStore FindInvoiceStore,
	invoiceDetailStore ListInvoiceDetailStore) *seeInvoiceDetailRepo {
	return &seeInvoiceDetailRepo{
		invoiceStore:       invoiceStore,
		invoiceDetailStore: invoiceDetailStore,
	}
}

func (biz *seeInvoiceDetailRepo) SeeInvoiceDetail(
	ctx context.Context,
	invoiceId string) (*invoicemodel.ResDetailInvoice, error) {
	invoice, errInvoice := biz.invoiceStore.FindInvoice(
		ctx,
		map[string]interface{}{
			"id": invoiceId,
		}, "CreatedByUser")
	if errInvoice != nil {
		return nil, errInvoice
	}

	resDetailInvoice := invoicemodel.GetResDetailInvoiceFromInvoice(invoice)

	details, errInvoiceDetail := biz.invoiceDetailStore.ListInvoiceDetail(
		ctx,
		invoiceId)
	if errInvoiceDetail != nil {
		return nil, errInvoiceDetail
	}
	resDetailInvoice.Details = details

	return resDetailInvoice, nil
}
