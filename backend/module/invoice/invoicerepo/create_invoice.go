package invoicerepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/customer/customermodel"
	"book-store-management-backend/module/invoice/invoicemodel"
	"book-store-management-backend/module/invoicedetail/invoicedetailmodel"
	"book-store-management-backend/module/shopgeneral/shopgeneralmodel"
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

type CustomerStore interface {
	FindCustomer(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*customermodel.Customer, error)
	UpdateCustomerPoint(
		ctx context.Context,
		id string,
		data *customermodel.CustomerUpdatePoint,
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

type ShopGeneralStore interface {
	FindShopGeneral(
		ctx context.Context,
	) (*shopgeneralmodel.ShopGeneral, error)
}

type createInvoiceRepo struct {
	invoiceStore            InvoiceStore
	invoiceDetailStore      InvoiceDetailStore
	customerStore           CustomerStore
	bookStore               BookStore
	stockChangeHistoryStore StockChangeHistoryStore
	shopGeneralStore        ShopGeneralStore
}

func NewCreateInvoiceRepo(
	invoiceStore InvoiceStore,
	invoiceDetailStore InvoiceDetailStore,
	customerStore CustomerStore,
	bookStore BookStore,
	stockChangeHistoryStore StockChangeHistoryStore,
	shopGeneralStore ShopGeneralStore) *createInvoiceRepo {
	return &createInvoiceRepo{
		invoiceStore:            invoiceStore,
		invoiceDetailStore:      invoiceDetailStore,
		customerStore:           customerStore,
		bookStore:               bookStore,
		stockChangeHistoryStore: stockChangeHistoryStore,
		shopGeneralStore:        shopGeneralStore,
	}
}

func (repo *createInvoiceRepo) GetShopGeneral(
	ctx context.Context) (*shopgeneralmodel.ShopGeneral, error) {
	if data, err := repo.shopGeneralStore.FindShopGeneral(ctx); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (repo *createInvoiceRepo) HandleData(
	ctx context.Context,
	data *invoicemodel.ReqCreateInvoice) error {
	totalPrice := 0
	totalImportPrice := 0

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
		totalImportPrice += *book.ImportPrice * invoiceDetail.Quantity

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
	data.TotalImportPrice = totalImportPrice

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

func (repo *createInvoiceRepo) FindCustomer(
	ctx context.Context,
	customerId string) (*customermodel.Customer, error) {
	customer, err := repo.customerStore.FindCustomer(
		ctx, map[string]interface{}{"id": customerId},
	)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (repo *createInvoiceRepo) UpdateCustomerPoint(
	ctx context.Context,
	customerId string,
	data customermodel.CustomerUpdatePoint) error {
	if err := repo.customerStore.UpdateCustomerPoint(
		ctx, customerId, &data,
	); err != nil {
		return err
	}
	return nil
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
