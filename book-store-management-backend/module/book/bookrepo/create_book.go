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

func (repo *createBookRepo) CreateBook(ctx context.Context, data *bookmodel.Book) error {
	strAuthorIDs := strings.Join(data.AuthorIDs, "|")
	strCategoryIDs := strings.Join(data.CategoryIDs, "|")
	dbData := bookstore.BookDBModel{
		ID:          data.ID,
		Name:        &data.Name,
		Description: &data.Description,
		Edition:     &data.Edition,
		Quantity:    &data.Quantity,
		ListedPrice: &data.ListedPrice,
		SellPrice:   &data.SellPrice,
		PublisherID: &data.PublisherID,
		AuthorIDs:   &strAuthorIDs,
		CategoryIDs: &strCategoryIDs,
	}
	if err := repo.store.CreateBook(ctx, &dbData); err != nil {
		return err
	}
	return nil
}
