package invoicebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/shopgeneral/shopgeneralmodel"
	"context"
)

type CreateInvoiceRepo interface {
	GetShopGeneral(
		ctx context.Context,
	) (*shopgeneralmodel.ShopGeneral, error)
	HandleData(
		ctx context.Context,
		data *invoicemodel.ReqCreateInvoice,
	) error
	HandleInvoice(
		ctx context.Context,
		data *invoicemodel.ReqCreateInvoice,
	) error
}

type createInvoiceBiz struct {
	gen       generator.IdGenerator
	repo      CreateInvoiceRepo
	requester middleware.Requester
}

func NewCreateInvoiceBiz(
	gen generator.IdGenerator,
	repo CreateInvoiceRepo,
	requester middleware.Requester) *createInvoiceBiz {
	return &createInvoiceBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createInvoiceBiz) CreateInvoice(
	ctx context.Context,
	data *invoicemodel.ReqCreateInvoice) error {
	if !biz.requester.IsHasFeature(common.InvoiceCreateFeatureCode) {
		return invoicemodel.ErrInvoiceCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleInvoiceId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.HandleData(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleInvoice(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleInvoiceId(gen generator.IdGenerator, data *invoicemodel.ReqCreateInvoice) error {
	idInvoice, err := gen.GenerateId()
	if err != nil {
		return err
	}
	data.Id = idInvoice

	for i := range data.InvoiceDetails {
		data.InvoiceDetails[i].InvoiceId = idInvoice
	}

	return err
}
