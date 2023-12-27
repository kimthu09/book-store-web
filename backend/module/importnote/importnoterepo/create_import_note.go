package importnoterepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/importnote/importnotemodel"
	"book-store-management-backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

type CreateImportNoteStore interface {
	CreateImportNote(
		ctx context.Context,
		data *importnotemodel.ReqCreateImportNote,
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
	UpdateImportPriceBook(
		ctx context.Context,
		id string,
		data *bookmodel.BookUpdateImportPrice,
	) error
}

type createImportNoteRepo struct {
	importNoteStore       CreateImportNoteStore
	importNoteDetailStore CreateImportNoteDetailStore
	bookStore             UpdatePriceBookStore
}

func NewCreateImportNoteRepo(
	importNoteStore CreateImportNoteStore,
	importNoteDetailStore CreateImportNoteDetailStore,
	bookStore UpdatePriceBookStore) *createImportNoteRepo {
	return &createImportNoteRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		bookStore:             bookStore,
	}
}

func (repo *createImportNoteRepo) HandleCreateImportNote(
	ctx context.Context,
	data *importnotemodel.ReqCreateImportNote) error {
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

func (repo *createImportNoteRepo) UpdateImportPriceBook(
	ctx context.Context,
	bookId string,
	price int) error {
	bookUpdatePrice := bookmodel.BookUpdateImportPrice{
		ImportPrice: &price,
	}

	if err := repo.bookStore.UpdateImportPriceBook(
		ctx, bookId, &bookUpdatePrice,
	); err != nil {
		return err
	}
	return nil
}
