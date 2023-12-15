package categoryrepo

import (
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type CategoryStore interface {
	CheckExistByID(ctx context.Context, id string) (bool, error)
	GetByListId(ctx context.Context, idList []string) ([]categorymodel.Category, error)
}

type CategoryPublicRepo interface {
	GetByListId(ctx context.Context, ids []string) ([]categorymodel.Category, error)
	IsExistCategoryId(ctx context.Context, categoryId string) bool
}

type categoryPublicRepo struct {
	store CategoryStore
}

func NewCategoryPublicRepo(store CategoryStore) *categoryPublicRepo {
	return &categoryPublicRepo{store: store}
}

func (repo *categoryPublicRepo) GetByListId(ctx context.Context, ids []string) ([]categorymodel.Category, error) {
	result, err := repo.store.GetByListId(ctx, ids)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *categoryPublicRepo) IsExistCategoryId(ctx context.Context, categoryId string) bool {
	isExist, err := repo.store.CheckExistByID(ctx, categoryId)

	if err != nil {
		return false
	}

	return isExist
}
