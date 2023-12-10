package bookrepo

import "context"

type DeleteBookStore interface {
	DeleteBook(ctx context.Context, id string) error
}

type deleteBookRepo struct {
	store DeleteBookStore
}

func NewDeleteBookRepo(store DeleteBookStore) *deleteBookRepo {
	return &deleteBookRepo{store: store}
}

func (repo *deleteBookRepo) DeleteBook(ctx context.Context, id string) error {
	err := repo.store.DeleteBook(ctx, id)
	return err
}
