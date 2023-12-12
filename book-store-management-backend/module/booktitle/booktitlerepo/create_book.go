package booktitlerepo

import (
	"book-store-management-backend/module/booktitle/booktitlemodel"
	booktitlestore "book-store-management-backend/module/booktitle/booktitlestore"
	"context"
	"strings"
)

type CreateBookStore interface {
	CreateBook(ctx context.Context, bookGeneral *booktitlestore.BookTitleDBModel) error
}

type createBookRepo struct {
	store CreateBookStore
}

func NewCreateBookRepo(store CreateBookStore) *createBookRepo {
	return &createBookRepo{store: store}
}

func (repo *createBookRepo) CreateBook(ctx context.Context, data *booktitlemodel.BookTitle) error {
	strAuthorIDs := strings.Join(data.AuthorIDs, "|")
	strCategoryIDs := strings.Join(data.CategoryIDs, "|")
	dbData := booktitlestore.BookTitleDBModel{
		ID:          data.ID,
		Name:        &data.Name,
		Description: &data.Description,
		AuthorIDs:   &strAuthorIDs,
		CategoryIDs: &strCategoryIDs,
	}
	if err := repo.store.CreateBook(ctx, &dbData); err != nil {
		return err
	}
	return nil

}
