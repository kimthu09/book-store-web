package bookrepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookstore"
	"context"
)

type CreateBookStore interface {
	CreateBook(ctx context.Context, bookGeneral *bookstore.BookDBModel) error
}

type createBookRepo struct {
	store CreateBookStore
}

func NewCreateBookRepo(store CreateBookStore) *createBookRepo {
	return &createBookRepo{store: store}
}

func (repo *createBookRepo) CreateBook(ctx context.Context, data *bookmodel.Book) error {

	dbData := bookstore.BookDBModel{
		ID:          data.ID,
		Name:        data.Name,
		BookTitleID: data.BookTitleID,
		PublisherID: data.PublisherID,
		Edition:     data.Edition,
		Quantity:    nil,
		ListedPrice: data.ListedPrice,
		SellPrice:   data.SellPrice,
		ImportPrice: nil,
	}
	err := repo.store.CreateBook(ctx, &dbData)
	if err != nil {
		return err
	}

	return nil
}
