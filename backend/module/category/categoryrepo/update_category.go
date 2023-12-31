package categoryrepo

import (
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type UpdateCategoryStore interface {
	UpdateCategory(
		ctx context.Context,
		id string,
		data *categorymodel.ReqUpdateCategory) error
}

type updateCategoryRepo struct {
	store UpdateCategoryStore
}

func NewUpdateCategoryRepo(store UpdateCategoryStore) *updateCategoryRepo {
	return &updateCategoryRepo{store: store}
}

func (repo *updateCategoryRepo) UpdateCategoryInfo(
	ctx context.Context,
	supplierId string,
	data *categorymodel.ReqUpdateCategory) error {
	if err := repo.store.UpdateCategory(ctx, supplierId, data); err != nil {
		return err
	}
	return nil
}
