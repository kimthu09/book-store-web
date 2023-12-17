package booktitlerepo

import (
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"context"
	"strings"
)

type UpdateBookTitleRepo interface {
	UpdateBookTitle(ctx context.Context, id string, data *booktitlemodel.BookTitle) error
}

type updateBookStore interface {
	UpdateBookTitle(ctx context.Context, id string, data *booktitlestore.BookTitleDBModel) error
}

type updateBookRepo struct {
	store updateBookStore
}

func NewUpdateBookRepo(store updateBookStore) *updateBookRepo {
	return &updateBookRepo{store: store}
}

func (repo *updateBookRepo) UpdateBookTitle(ctx context.Context, id string, data *booktitlemodel.BookTitle) error {
	var dbAuthorIDs *string = nil
	if data.AuthorIDs != nil {
		tmp := strings.Join(*data.AuthorIDs, "|")
		dbAuthorIDs = &tmp
	}

	var dbCategoryIDs *string = nil
	if data.CategoryIDs != nil {
		tmp := strings.Join(*data.CategoryIDs, "|")
		dbCategoryIDs = &tmp
	}

	dbData := booktitlestore.BookTitleDBModel{
		ID:          nil,
		Name:        data.Name,
		Description: data.Description,
		AuthorIDs:   dbAuthorIDs,
		CategoryIDs: dbCategoryIDs,
	}
	err := repo.store.UpdateBookTitle(ctx, id, &dbData)
	return err
}
