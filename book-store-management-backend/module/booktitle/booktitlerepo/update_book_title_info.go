package booktitlerepo

import (
	"book-store-management-backend/module/booktitle/booktitlestore"
	"context"
)

type updateBookStore interface {
	UpdateBook(ctx context.Context, id string, data *booktitlestore.BookTitleDBModel) error
}

type updateBookRepo struct {
	store updateBookStore
}

func NewUpdateBookRepo(store updateBookStore) *updateBookRepo {
	return &updateBookRepo{store: store}
}

func (repo *updateBookRepo) UpdateBook(ctx context.Context, id string, data *booktitlestore.BookTitleDBModel) error {
	err := repo.store.UpdateBook(ctx, id, data)
	return err
}
