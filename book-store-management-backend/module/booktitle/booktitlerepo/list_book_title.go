package booktitlerepo

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/booktitle/booktitlestore"
	"context"
	"strings"
)

type ListBookStore interface {
	ListBook(ctx context.Context, filter *booktitlemodel.Filter, propertiesContainSearchKey []string, paging *common.Paging) ([]booktitlestore.BookTitleDBModel, error)
}

type listBookTitleRepo struct {
	store ListBookStore
}

func NewListBookTitleRepo(store ListBookStore) *listBookTitleRepo {
	return &listBookTitleRepo{store: store}
}

func (repo *listBookTitleRepo) ListBookTitle(ctx context.Context, filter *booktitlemodel.Filter, paging *common.Paging) ([]booktitlemodel.BookTitle, error) {
	resultDbModel, err := repo.store.ListBook(ctx, filter, []string{"name"}, paging)

	if err != nil {
		return nil, err
	}

	result := make([]booktitlemodel.BookTitle, len(resultDbModel))
	for i, v := range resultDbModel {
		result[i] = booktitlemodel.BookTitle{
			ID:          v.ID,
			Name:        *v.Name,
			Description: *v.Description,
			AuthorIDs:   strings.Split(*v.AuthorIDs, "|"),
			CategoryIDs: strings.Split(*v.CategoryIDs, "|"),
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			DeletedAt:   v.DeletedAt,
			IsActive:    v.IsActive,
		}
	}
	return result, nil
}