package categorybiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type createListCategoryBiz struct {
	gen       generator.IdGenerator
	repo      CreateCategoryRepo
	requester middleware.Requester
}

func NewCreateListCategoryBiz(
	gen generator.IdGenerator,
	repo CreateCategoryRepo,
	requester middleware.Requester) *createListCategoryBiz {
	return &createListCategoryBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createListCategoryBiz) CreateListCategory(
	ctx context.Context, data []categorymodel.Category) error {
	if !biz.requester.IsHasFeature(common.CategoryCreateFeatureCode) {
		return categorymodel.ErrCategoryCreateNoPermission
	}

	for _, v := range data {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	for i, _ := range data {
		if err := handleCategoryId(biz.gen, &data[i]); err != nil {
			return err
		}
	}

	for _, v := range data {
		if err := biz.repo.CreateCategory(ctx, &v); err != nil {
			return err
		}
	}

	return nil
}
