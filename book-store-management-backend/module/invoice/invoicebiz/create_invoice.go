package invoicebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type CreateInvoiceRepo interface {
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
	data *invoicemodel.ReqCreateInvoice) (*invoicemodel.ResCreateInvoice, error) {
	if !biz.requester.IsHasFeature(common.InvoiceCreateFeatureCode) {
		return nil, invoicemodel.ErrInvoiceCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	if err := handleInvoiceId(biz.gen, data); err != nil {
		return nil, err
	}

	if err := biz.repo.HandleData(ctx, data); err != nil {
		return nil, err
	}

	if err := biz.repo.HandleInvoice(ctx, data); err != nil {
		return nil, err
	}

	var details []invoicedetailmodel.ReqCreateInvoiceDetail
	for _, v := range data.InvoiceDetails {
		reqCreateInvoiceDetail := invoicedetailmodel.ReqCreateInvoiceDetail{
			InvoiceId: v.InvoiceId,
			BookId:    v.BookId,
			BookName:  v.BookName,
			Quantity:  v.Quantity,
			UnitPrice: v.UnitPrice,
		}
		details = append(details, reqCreateInvoiceDetail)
	}

	resCreateInvoiceData := invoicemodel.ResCreateInvoiceData{
		Id:      data.Id,
		Details: details,
		Total:   data.TotalPrice,
		CreatedBy: usermodel.SimpleUser{
			Id:   biz.requester.GetUserId(),
			Name: biz.requester.GetName(),
		},
	}

	var resCreateInvoice invoicemodel.ResCreateInvoice
	resCreateInvoice.ResCreateInvoiceData = resCreateInvoiceData

	return &resCreateInvoice, nil
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
