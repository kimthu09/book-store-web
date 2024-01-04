package bookrepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookstore"
	"context"
)

type UpdateBookRepo interface {
	UpdateBook(ctx context.Context, id string, data *bookstore.BookDBModel) error
}

type updateBookRepo struct {
	bookStore      UpdateBookRepo
	bookTitleStore GetBookTitleStore
}

func NewUpdateBookRepo(
	bookStore UpdateBookRepo,
	bookTitleStore GetBookTitleStore) *updateBookRepo {
	return &updateBookRepo{
		bookStore:      bookStore,
		bookTitleStore: bookTitleStore,
	}
}

func (repo *updateBookRepo) UpdateBook(ctx context.Context, data *bookmodel.Book) error {
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

	if err := repo.bookStore.UpdateBook(ctx, *dbData.ID, &dbData); err != nil {
		return err
	}

	return nil
}
