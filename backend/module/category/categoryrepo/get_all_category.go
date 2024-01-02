package categoryrepo

import (
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type GetAllCategoryStore interface {
	GetAllCategory(ctx context.Context) ([]categorymodel.SimpleCategory, error)
}

type getAllCategoryStore struct {
	store GetAllCategoryStore
}

func NewGetAllCategoryRepo(store GetAllCategoryStore) *getAllCategoryStore {
	return &getAllCategoryStore{store: store}
}

func (repo *getAllCategoryStore) GetAllCategory(
	ctx context.Context) ([]categorymodel.SimpleCategory, error) {
	result, err := repo.store.GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
