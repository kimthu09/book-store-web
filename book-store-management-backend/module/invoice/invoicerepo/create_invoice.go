package invoicerepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorymodel"
	"context"
)

type InvoiceStore interface {
	CreateInvoice(
		ctx context.Context,
		data *invoicemodel.ReqCreateInvoice,
	) error
}

type InvoiceDetailStore interface {
	CreateListImportNoteDetail(
		ctx context.Context,
		data []invoicedetailmodel.ReqCreateInvoiceDetail,
	) error
}

type BookStore interface {
	FindBook(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*bookmodel.Book, error)
	UpdateQuantityBook(
		ctx context.Context,
		id string,
		data *bookmodel.BookUpdateQuantity) error
}

type StockChangeHistoryStore interface {
	CreateLisStockChangeHistory(
		ctx context.Context,
		data []stockchangehistorymodel.StockChangeHistory) error
}

type createInvoiceRepo struct {
	invoiceStore            InvoiceStore
	invoiceDetailStore      InvoiceDetailStore
	bookStore               BookStore
	stockChangeHistoryStore StockChangeHistoryStore
}

func NewCreateInvoiceRepo(
	invoiceStore InvoiceStore,
	invoiceDetailStore InvoiceDetailStore,
	bookStore BookStore,
	stockChangeHistoryStore StockChangeHistoryStore) *createInvoiceRepo {
	return &createInvoiceRepo{
		invoiceStore:            invoiceStore,
		invoiceDetailStore:      invoiceDetailStore,
		bookStore:               bookStore,
		stockChangeHistoryStore: stockChangeHistoryStore,
	}
}

func (repo *createInvoiceRepo) HandleData(
	ctx context.Context,
	data *invoicemodel.ReqCreateInvoice) error {
	totalPrice := 0
	var history []stockchangehistorymodel.StockChangeHistory
	for i, invoiceDetail := range data.InvoiceDetails {
		book, errGetActiveBook := repo.getBook(ctx, invoiceDetail.BookId)
		if errGetActiveBook != nil {
			return errGetActiveBook
		}

		qtyLeft := *book.Quantity - invoiceDetail.Quantity
		if qtyLeft < 0 {
			return invoicemodel.ErrInvoiceBookIsNotEnough
		}

		data.InvoiceDetails[i].BookName = *book.Name
		data.InvoiceDetails[i].UnitPrice = *book.SellPrice
		totalPrice += data.InvoiceDetails[i].UnitPrice * invoiceDetail.Quantity

		typeChange := stockchangehistorymodel.Sell
		stockChangeHistory := stockchangehistorymodel.StockChangeHistory{
			Id:           data.Id,
			BookId:       invoiceDetail.BookId,
			Quantity:     -invoiceDetail.Quantity,
			QuantityLeft: qtyLeft,
			Type:         &typeChange,
		}
		history = append(history, stockChangeHistory)

		bookUpdateQuantity := bookmodel.BookUpdateQuantity{
			QuantityUpdate: -invoiceDetail.Quantity,
		}
		if err := repo.bookStore.UpdateQuantityBook(
			ctx, invoiceDetail.BookId, &bookUpdateQuantity); err != nil {
			return err
		}
	}

	data.TotalPrice = totalPrice

	if err := repo.stockChangeHistoryStore.CreateLisStockChangeHistory(
		ctx, history); err != nil {
		return err
	}

	return nil
}

func (repo *createInvoiceRepo) getBook(
	ctx context.Context,
	bookId string) (*bookmodel.Book, error) {
	book, err := repo.bookStore.FindBook(
		ctx,
		map[string]interface{}{
			"Id": bookId,
		},
	)
	if err != nil {
		return nil, err
	}

	if !*book.IsActive {
		return nil, invoicedetailmodel.ErrInvoiceDetailBookIsInactive
	}

	return book, nil
}

func (repo *createInvoiceRepo) HandleInvoice(
	ctx context.Context,
	data *invoicemodel.ReqCreateInvoice) error {
	if err := repo.createInvoice(ctx, data); err != nil {
		return err
	}

	if err := repo.createInvoiceDetails(ctx, data.InvoiceDetails); err != nil {
		return err
	}
	return nil
}

func (repo *createInvoiceRepo) createInvoice(
	ctx context.Context,
	data *invoicemodel.ReqCreateInvoice) error {
	if err := repo.invoiceStore.CreateInvoice(ctx, data); err != nil {
		return err
	}
	return nil
}

func (repo *createInvoiceRepo) createInvoiceDetails(
	ctx context.Context,
	data []invoicedetailmodel.ReqCreateInvoiceDetail) error {
	if err := repo.invoiceDetailStore.CreateListImportNoteDetail(
		ctx, data,
	); err != nil {
		return err
	}
	return nil
}
