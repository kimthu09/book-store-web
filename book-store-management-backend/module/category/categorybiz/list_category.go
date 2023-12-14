package categorybiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type ListCategoryRepo interface {
	ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	repo ListCategoryRepo
}

func NewListCategoryRepo(repo ListCategoryRepo) *listCategoryBiz {
	return &listCategoryBiz{repo: repo}
}

func (biz *listCategoryBiz) ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]categorymodel.Category, error) {
	result, err := biz.repo.ListCategory(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
