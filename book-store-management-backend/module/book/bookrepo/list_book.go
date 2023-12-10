package bookrepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookstore"
	"context"
	"strings"
)

type ListBookStore interface {
	ListBook(ctx context.Context, filter *bookmodel.Filter, propertiesContainSearchKey []string, paging *common.Paging) ([]bookstore.BookDBModel, error)
}

type listBookRepo struct {
	store ListBookStore
}

func NewListBookRepo(store ListBookStore) *listBookRepo {
	return &listBookRepo{store: store}
}

func (repo *listBookRepo) ListBook(ctx context.Context, filter *bookmodel.Filter, paging *common.Paging) ([]bookmodel.Book, error) {
	//return []bookstore.BookDBModel{
	//	{
	//		ID:        nil,
	//		Name:      "Sách 1",
	//		AuthorIDs: "1|2",
	//	},
	//	{
	//		ID:        nil,
	//		Name:      "Sách 2",
	//		AuthorIDs: "2|5",
	//	},
	//}, nil

	resultDbModel, err := repo.store.ListBook(ctx, filter, []string{"name"}, paging)

	if err != nil {
		return nil, err
	}

	result := make([]bookmodel.Book, len(resultDbModel))
	for i, v := range resultDbModel {
		result[i] = bookmodel.Book{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Edition:     v.Edition,
			Quantity:    v.Quantity,
			ListedPrice: v.ListedPrice,
			SellPrice:   v.SellPrice,
			PublisherID: v.PublisherID,
			AuthorIDs:   strings.Split(v.AuthorIDs, "|"),
			CategoryIDs: strings.Split(v.CategoryIDs, "|"),
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			DeletedAt:   v.DeletedAt,
			IsActive:    v.IsActive,
		}
	}
	return result, nil
}
