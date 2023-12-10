package bookrepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookstore"
	"context"
	"strings"
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

func (biz *createBookRepo) CreateBook(ctx context.Context, data *bookmodel.Book) error {
	dbData := bookstore.BookDBModel{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Edition:     data.Edition,
		Quantity:    data.Quantity,
		ListedPrice: data.ListedPrice,
		SellPrice:   data.SellPrice,
		PublisherID: data.PublisherID,
		AuthorIDs:   strings.Join(data.AuthorIDs, "|"),
		CategoryIDs: strings.Join(data.CategoryIDs, "|"),
	}
	if err := biz.store.CreateBook(ctx, &dbData); err != nil {
		return err
	}
	return nil
}
