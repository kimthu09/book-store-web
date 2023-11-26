package bookrepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

type CreateBookStore interface {
	CreateBook(ctx context.Context, data *bookmodel.BookCreate) error
}

type createBookRepo struct {
	store CreateBookStore
}

func NewCreateBookRepo(store CreateBookStore) *createBookRepo {
	return &createBookRepo{store: store}
}

func (biz *createBookRepo) CreateBook(ctx context.Context, data *bookmodel.BookCreate) error {
	if err := biz.store.CreateBook(ctx, data); err != nil {
		return err
	}

	return nil
}
