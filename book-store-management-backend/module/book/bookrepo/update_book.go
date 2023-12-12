package bookrepo

import (
	"book-store-management-backend/module/book/bookstore"
	"context"
)

type updateBookStore interface {
	UpdateBook(ctx context.Context, id string, data *bookstore.BookDBModel) error
}

type updateBookRepo struct {
	store updateBookStore
}

func NewUpdateBookRepo(store updateBookStore) *updateBookRepo {
	return &updateBookRepo{store: store}
}

func (repo *updateBookRepo) UpdateBook(ctx context.Context, id string, data *bookstore.BookDBModel) error {
	err := repo.store.UpdateBook(ctx, id, data)
	return err
}
