package bookrepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookstore"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

type CreateBookStore interface {
	CreateBook(ctx context.Context, bookGeneral *bookstore.BookDBModel) error
}

type GetBookTitleStore interface {
	FindBookTitle(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*booktitlemodel.SimpleBookTitle, error)
}

type createBookRepo struct {
	bookStore      CreateBookStore
	bookTitleStore GetBookTitleStore
}

func NewCreateBookRepo(
	bookStore CreateBookStore,
	bookTitleStore GetBookTitleStore) *createBookRepo {
	return &createBookRepo{
		bookStore:      bookStore,
		bookTitleStore: bookTitleStore,
	}
}

func (repo *createBookRepo) CreateBook(ctx context.Context, data *bookmodel.Book) error {
	bookTitle, err := repo.bookTitleStore.FindBookTitle(ctx, map[string]interface{}{"id": data.BookTitleID})
	if err != nil {
		return err
	}

	dbData := bookstore.BookDBModel{
		ID:          data.ID,
		Name:        &bookTitle.Name,
		BookTitleID: data.BookTitleID,
		Image:       data.Image,
		PublisherID: data.PublisherID,
		Edition:     data.Edition,
		Quantity:    nil,
		ListedPrice: data.ListedPrice,
		SellPrice:   data.SellPrice,
		ImportPrice: data.ListedPrice,
	}
	if err := repo.bookStore.CreateBook(ctx, &dbData); err != nil {
		return err
	}

	return nil
}
