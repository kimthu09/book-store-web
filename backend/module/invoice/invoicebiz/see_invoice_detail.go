package invoicebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/shopgeneral/shopgeneralmodel"
	"context"
)

type SeeShopGeneralRepo interface {
	FindShopGeneral(
		ctx context.Context) (*shopgeneralmodel.ShopGeneral, error)
}

type SeeInvoiceDetailRepo interface {
	SeeInvoiceDetail(
		ctx context.Context,
		invoiceId string) (*invoicemodel.Invoice, error)
}

type seeDetailInvoiceBiz struct {
	invoiceRepo SeeInvoiceDetailRepo
	shopRepo    SeeShopGeneralRepo
	requester   middleware.Requester
}

func NewSeeDetailInvoiceBiz(
	invoiceRepo SeeInvoiceDetailRepo,
	shopRepo SeeShopGeneralRepo,
	requester middleware.Requester) *seeDetailInvoiceBiz {
	return &seeDetailInvoiceBiz{
		invoiceRepo: invoiceRepo,
		shopRepo:    shopRepo,
		requester:   requester,
	}
}

func (biz *seeDetailInvoiceBiz) SeeDetailInvoice(
	ctx context.Context,
	invoiceId string) (*invoicemodel.ResDetailInvoice, error) {
	if !biz.requester.IsHasFeature(common.InvoiceViewFeatureCode) {
		return nil, invoicemodel.ErrInvoiceViewNoPermission
	}

	invoice, errInvoice := biz.invoiceRepo.SeeInvoiceDetail(
		ctx, invoiceId)
	if errInvoice != nil {
		return nil, errInvoice
	}

	general, errGetGeneral := biz.shopRepo.FindShopGeneral(ctx)
	if errGetGeneral != nil {
		return nil, errGetGeneral
	}

	invoiceDetail := invoicemodel.ResDetailInvoice{
		Invoice:     *invoice,
		ShopGeneral: *general,
	}

	return &invoiceDetail, nil
}
