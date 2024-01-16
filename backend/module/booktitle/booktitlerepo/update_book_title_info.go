package booktitlerepo

import (
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"context"
	"strings"
)

type updateBookTitleStore interface {
	UpdateBookTitle(ctx context.Context, id string, data *booktitlestore.BookTitleDBModel) error
}

type UpdateBookStore interface {
	UpdateName(ctx context.Context, bookTitleId string, name *string) error
}

type updateBookRepo struct {
	bookTitleStore updateBookTitleStore
	bookStore      UpdateBookStore
}

func NewUpdateBookRepo(
	bookStore UpdateBookStore, bookTitleStore updateBookTitleStore) *updateBookRepo {
	return &updateBookRepo{bookTitleStore: bookTitleStore, bookStore: bookStore}
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

	if data.Name != nil {
		errBook := repo.bookStore.UpdateName(ctx, id, data.Name)
		if errBook != nil {
			return errBook
		}
	}

	dbData := booktitlestore.BookTitleDBModel{
		ID:          nil,
		Name:        data.Name,
		Description: data.Description,
		AuthorIDs:   dbAuthorIDs,
		CategoryIDs: dbCategoryIDs,
	}
	err := repo.bookTitleStore.UpdateBookTitle(ctx, id, &dbData)
	return err
}
