package inventorychecknoterepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotemodel"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
)

type CreateInventoryCheckNoteStore interface {
	CreateInventoryCheckNote(
		ctx context.Context,
		data *inventorychecknotemodel.InventoryCheckNoteCreate,
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

type createInventoryCheckNoteRepo struct {
	inventoryCheckNoteStore       CreateInventoryCheckNoteStore
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore
	bookStore                     UpdateBookStore
}

func NewCreateInventoryCheckNoteRepo(
	inventoryCheckNoteStore CreateInventoryCheckNoteStore,
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore,
	bookStore UpdateBookStore) *createInventoryCheckNoteRepo {
	return &createInventoryCheckNoteRepo{
		inventoryCheckNoteStore:       inventoryCheckNoteStore,
		inventoryCheckNoteDetailStore: inventoryCheckNoteDetailStore,
		bookStore:                     bookStore,
	}
}

func (repo *createInventoryCheckNoteRepo) HandleInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
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
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
	qtyDiff := 0
	qtyAfter := 0
	for i, value := range data.Details {
		book, errGetBook := repo.bookStore.FindBook(
			ctx, map[string]interface{}{"id": value.BookId})
		if errGetBook != nil {
			return errGetBook
		}

		data.Details[i].Initial = book.Quantity
		data.Details[i].Final = book.Quantity + value.Difference
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
	}

	data.QuantityDifferent = qtyDiff
	data.QuantityAfterAdjust = qtyAfter
	return nil
}
