package categoryrepo

import "context"

type ExistCategoryStore interface {
	CheckExistByID(ctx context.Context, id string) (bool, error)
}

type existCategoryRepo struct {
	store ExistCategoryStore
}

func NewExistCategoryRepo(store ExistCategoryStore) *existCategoryRepo {
	return &existCategoryRepo{store: store}
}

func (repo *existCategoryRepo) IsExistCategoryId(ctx context.Context, categoryId string) bool {
	isExist, err := repo.store.CheckExistByID(ctx, categoryId)
	if err != nil {
		return false
	}
	return isExist
}
