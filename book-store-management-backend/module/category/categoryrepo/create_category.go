package categoryrepo

import (
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type CreateCategory interface {
	CreateCategory(ctx context.Context, data *categorymodel.Category) error
}

type createCategoryRepo struct {
	store CreateCategory
}

func NewCreateCategoryRepo(store CreateCategory) *createCategoryRepo {
	return &createCategoryRepo{store: store}
}

func (repo *createCategoryRepo) CreateCategory(ctx context.Context, data *categorymodel.Category) error {
	if err := repo.store.CreateCategory(ctx, data); err != nil {
		return err
	}
	return nil
}
