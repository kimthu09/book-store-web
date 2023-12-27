package categorybiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type CreateCategoryRepo interface {
	CreateCategory(ctx context.Context, data *categorymodel.Category) error
}

type createCategoryBiz struct {
	gen       generator.IdGenerator
	repo      CreateCategoryRepo
	requester middleware.Requester
}

func NewCreateCategoryBiz(gen generator.IdGenerator, repo CreateCategoryRepo, requester middleware.Requester) *createCategoryBiz {
	return &createCategoryBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *categorymodel.Category) error {
	if !biz.requester.IsHasFeature(common.CategoryCreateFeatureCode) {
		return categorymodel.ErrCategoryCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleCategoryId(biz.gen, data); err != nil {
		return err
	}
	if err := biz.repo.CreateCategory(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleCategoryId(gen generator.IdGenerator, data *categorymodel.Category) error {
	id, err := gen.IdProcess(&data.Id)
	if err != nil {
		return err
	}
	data.Id = *id
	return nil
}
