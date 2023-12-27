package categorybiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type ListCategoryRepo interface {
	ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	repo      ListCategoryRepo
	requester middleware.Requester
}

func NewListCategoryRepo(repo ListCategoryRepo, requester middleware.Requester) *listCategoryBiz {
	return &listCategoryBiz{repo: repo, requester: requester}
}

func (biz *listCategoryBiz) ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]categorymodel.Category, error) {
	if !biz.requester.IsHasFeature(common.CategoryViewFeatureCode) {
		return nil, categorymodel.ErrCategoryViewNoPermission
	}

	result, err := biz.repo.ListCategory(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
