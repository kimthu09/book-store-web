package categorybiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type UpdateCategoryRepo interface {
	UpdateCategoryInfo(
		ctx context.Context,
		supplierId string,
		data *categorymodel.ReqUpdateCategory) error
}

type updateCategoryBiz struct {
	repo      UpdateCategoryRepo
	requester middleware.Requester
}

func NewUpdateCategoryBiz(
	repo UpdateCategoryRepo,
	requester middleware.Requester) *updateCategoryBiz {
	return &updateCategoryBiz{repo: repo, requester: requester}
}

func (biz *updateCategoryBiz) UpdateCategory(
	ctx context.Context,
	id string,
	data *categorymodel.ReqUpdateCategory) error {
	if !biz.requester.IsHasFeature(common.CategoryUpdateFeatureCode) {
		return categorymodel.ErrCategoryUpdateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.UpdateCategoryInfo(ctx, id, data); err != nil {
		return err
	}

	return nil
}
