package customerbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/customer/customermodel"
	"book-store-management-backend/module/invoice/invoicemodel"
	"context"
)

type SeeCustomerInvoiceRepo interface {
	SeeCustomerInvoice(
		ctx context.Context,
		customerId string,
		filter *customermodel.FilterInvoice,
		paging *common.Paging) ([]invoicemodel.Invoice, error)
}

type seeCustomerInvoiceBiz struct {
	repo      SeeCustomerInvoiceRepo
	requester middleware.Requester
}

func NewSeeCustomerInvoiceBiz(
	repo SeeCustomerInvoiceRepo,
	requester middleware.Requester) *seeCustomerInvoiceBiz {
	return &seeCustomerInvoiceBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeCustomerInvoiceBiz) SeeCustomerInvoice(
	ctx context.Context,
	customerId string,
	filter *customermodel.FilterInvoice,
	paging *common.Paging) ([]invoicemodel.Invoice, error) {
	if !biz.requester.IsHasFeature(common.CustomerViewFeatureCode) {
		return nil, customermodel.ErrCustomerViewNoPermission
	}

	invoices, err := biz.repo.SeeCustomerInvoice(ctx, customerId, filter, paging)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}
