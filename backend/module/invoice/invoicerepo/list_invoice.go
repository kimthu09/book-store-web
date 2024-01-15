package invoicerepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
)

type ListInvoiceStore interface {
	ListInvoice(
		ctx context.Context,
		filter *invoicemodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
		moreKeys ...string) ([]invoicemodel.Invoice, error)
}

type listInvoiceRepo struct {
	store ListInvoiceStore
}

func NewListImportNoteRepo(store ListInvoiceStore) *listInvoiceRepo {
	return &listInvoiceRepo{store: store}
}

func (repo *listInvoiceRepo) ListInvoice(
	ctx context.Context,
	filter *invoicemodel.Filter,
	paging *common.Paging) ([]invoicemodel.Invoice, error) {
	result, err := repo.store.ListInvoice(
		ctx,
		filter,
		[]string{"Invoice.id"},
		paging,
		"CreatedByUser", "Customer")

	if err != nil {
		return nil, err
	}

	return result, nil
}
