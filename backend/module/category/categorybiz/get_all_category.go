package categorybiz

import (
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type GetAllCategoryRepo interface {
	GetAllCategory(ctx context.Context) ([]categorymodel.SimpleCategory, error)
}

type getAllCategoryBiz struct {
	repo GetAllCategoryRepo
}

func NewGetAllCategoryBiz(repo GetAllCategoryRepo) *getAllCategoryBiz {
	return &getAllCategoryBiz{repo: repo}
}

func (biz *getAllCategoryBiz) GetAllCategory(
	ctx context.Context) ([]categorymodel.SimpleCategory, error) {
	result, err := biz.repo.GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
