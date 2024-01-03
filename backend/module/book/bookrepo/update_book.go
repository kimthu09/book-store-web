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
	store UpdateBookRepo
}

func NewUpdateBookRepo(store UpdateBookRepo) *updateBookRepo {
	return &updateBookRepo{store: store}
}

func (repo *updateBookRepo) UpdateBook(ctx context.Context, data *bookmodel.Book) error {
	dbData := bookstore.BookDBModel{
		ID:          data.ID,
		Name:        data.Name,
		BookTitleID: data.BookTitleID,
		Image:       data.Image,
		PublisherID: data.PublisherID,
		Edition:     data.Edition,
		Quantity:    nil,
		ListedPrice: data.ListedPrice,
		SellPrice:   data.SellPrice,
		ImportPrice: data.ListedPrice,
	}

	err := repo.store.UpdateBook(ctx, *dbData.ID, &dbData)
	if err != nil {
		return err
	}

	return nil
}
