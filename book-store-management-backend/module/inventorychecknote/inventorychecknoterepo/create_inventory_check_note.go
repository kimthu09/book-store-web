package inventorychecknoterepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"book-store-management-backend/module/stockchangehistory/stockchangehistorymodel"
	"context"
)

type CreateInventoryCheckNoteStore interface {
	CreateInventoryCheckNote(
		ctx context.Context,
		data *inventorychecknotemodel.ReqCreateInventoryCheckNote,
	) error
}

type CreateInventoryCheckNoteDetailStore interface {
	CreateListInventoryCheckNoteDetail(
		ctx context.Context,
		data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate,
	) error
}

type UpdateBookStore interface {
	UpdateQuantityBook(
		ctx context.Context,
		id string,
		data *bookmodel.BookUpdateQuantity,
	) error
	FindBook(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*bookmodel.Book, error)
}

type StockChangeHistoryStore interface {
	CreateLisStockChangeHistory(
		ctx context.Context,
		data []stockchangehistorymodel.StockChangeHistory) error
}

type createInventoryCheckNoteRepo struct {
	inventoryCheckNoteStore       CreateInventoryCheckNoteStore
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore
	bookStore                     UpdateBookStore
	stockChangeHistoryStore       StockChangeHistoryStore
}

func NewCreateInventoryCheckNoteRepo(
	inventoryCheckNoteStore CreateInventoryCheckNoteStore,
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore,
	bookStore UpdateBookStore,
	stockChangeHistoryStore StockChangeHistoryStore) *createInventoryCheckNoteRepo {
	return &createInventoryCheckNoteRepo{
		inventoryCheckNoteStore:       inventoryCheckNoteStore,
		inventoryCheckNoteDetailStore: inventoryCheckNoteDetailStore,
		bookStore:                     bookStore,
		stockChangeHistoryStore:       stockChangeHistoryStore,
	}
}

func (repo *createInventoryCheckNoteRepo) HandleInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.ReqCreateInventoryCheckNote) error {
	if err := repo.inventoryCheckNoteStore.CreateInventoryCheckNote(ctx, data); err != nil {
		return err
	}

	if err := repo.inventoryCheckNoteDetailStore.CreateListInventoryCheckNoteDetail(
		ctx, data.Details,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createInventoryCheckNoteRepo) HandleBookQuantity(
	ctx context.Context,
	data *inventorychecknotemodel.ReqCreateInventoryCheckNote) error {
	qtyDiff := 0
	qtyAfter := 0

	var history []stockchangehistorymodel.StockChangeHistory
	for i, value := range data.Details {
		book, errGetBook := repo.bookStore.FindBook(
			ctx, map[string]interface{}{"id": value.BookId})
		if errGetBook != nil {
			return errGetBook
		}

		data.Details[i].Initial = *book.Quantity
		data.Details[i].Final = *book.Quantity + value.Difference
		qtyDiff += value.Difference
		qtyAfter += data.Details[i].Final

		if data.Details[i].Final < 0 {
			return inventorychecknotemodel.ErrInventoryCheckNoteModifyQuantityIsInvalid
		}

		bookUpdate := bookmodel.BookUpdateQuantity{QuantityUpdate: value.Difference}
		if err := repo.bookStore.UpdateQuantityBook(
			ctx, value.BookId, &bookUpdate,
		); err != nil {
			return err
		}

		typeChange := stockchangehistorymodel.Modify
		stockChangeHistory := stockchangehistorymodel.StockChangeHistory{
			Id:           *data.Id,
			BookId:       data.Details[i].BookId,
			Quantity:     value.Difference,
			QuantityLeft: value.Difference + *book.Quantity,
			Type:         &typeChange,
		}
		history = append(history, stockChangeHistory)
	}

	if err := repo.stockChangeHistoryStore.CreateLisStockChangeHistory(
		ctx, history); err != nil {
		return err
	}

	data.QuantityDifferent = qtyDiff
	data.QuantityAfterAdjust = qtyAfter
	return nil
}
