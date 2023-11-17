package importnoterepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/importnote/importnotemodel"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"book-store-management-backend/module/supplier/suppliermodel"
	"context"
)

type CreateImportNoteStore interface {
	CreateImportNote(
		ctx context.Context,
		data *importnotemodel.ImportNoteCreate,
	) error
}

type CreateImportNoteDetailStore interface {
	CreateListImportNoteDetail(
		ctx context.Context,
		data []importnotedetailmodel.ImportNoteDetailCreate,
	) error
}

type UpdatePriceBookStore interface {
	FindBook(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*bookmodel.Book, error)
	UpdatePriceBook(
		ctx context.Context,
		id string,
		data *bookmodel.BookUpdatePrice,
	) error
}

type CheckSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*suppliermodel.Supplier, error)
}

type createImportNoteRepo struct {
	importNoteStore       CreateImportNoteStore
	importNoteDetailStore CreateImportNoteDetailStore
	bookStore             UpdatePriceBookStore
	supplierStore         CheckSupplierStore
}

func NewCreateImportNoteRepo(
	importNoteStore CreateImportNoteStore,
	importNoteDetailStore CreateImportNoteDetailStore,
	bookStore UpdatePriceBookStore,
	supplierStore CheckSupplierStore) *createImportNoteRepo {
	return &createImportNoteRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		bookStore:             bookStore,
		supplierStore:         supplierStore,
	}
}

func (repo *createImportNoteRepo) CheckBook(
	ctx context.Context,
	bookId string) error {
	if _, err := repo.bookStore.FindBook(
		ctx, map[string]interface{}{"id": bookId},
	); err != nil {
		return err
	}
	return nil
}

func (repo *createImportNoteRepo) CheckSupplier(
	ctx context.Context,
	supplierId string) error {
	if _, err := repo.supplierStore.FindSupplier(
		ctx, map[string]interface{}{"id": supplierId},
	); err != nil {
		return err
	}
	return nil
}

func (repo *createImportNoteRepo) HandleCreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	if err := repo.importNoteStore.CreateImportNote(ctx, data); err != nil {
		return err
	}
	if err := repo.importNoteDetailStore.CreateListImportNoteDetail(
		ctx,
		data.ImportNoteDetails); err != nil {
		return err
	}
	return nil
}

func (repo *createImportNoteRepo) UpdatePriceBook(
	ctx context.Context,
	bookId string,
	price float32) error {
	bookUpdatePrice := bookmodel.BookUpdatePrice{
		Price: &price,
	}

	if err := repo.bookStore.UpdatePriceBook(
		ctx, bookId, &bookUpdatePrice,
	); err != nil {
		return err
	}
	return nil
}
