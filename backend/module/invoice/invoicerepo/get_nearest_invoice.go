package invoicerepo

import (
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
)

type GetNearestInvoiceStore interface {
	GetNearest(
		ctx context.Context,
		amountNeed int,
		moreKeys ...string) ([]invoicemodel.Invoice, error)
}

type getNearestInvoiceRepo struct {
	store GetNearestInvoiceStore
}

func NewGetNearestInvoiceRepo(store GetNearestInvoiceStore) *getNearestInvoiceRepo {
	return &getNearestInvoiceRepo{store: store}
}

func (repo *getNearestInvoiceRepo) GetNearestInvoice(
	ctx context.Context,
	amountNeed int) ([]invoicemodel.Invoice, error) {
	result, err := repo.store.GetNearest(
		ctx,
		amountNeed,
		"CreatedByUser")

	if err != nil {
		return nil, err
	}

	return result, nil
}
