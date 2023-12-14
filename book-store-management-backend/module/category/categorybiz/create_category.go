package categorybiz

import (
	"book-store-management-backend/component/generator"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type CreateCategoryRepo interface {
	CreateCategory(ctx context.Context, data *categorymodel.Category) error
}

type createCategoryBiz struct {
	gen  generator.IdGenerator
	repo CreateCategoryRepo
}

func NewCreateCategoryBiz(gen generator.IdGenerator, repo CreateCategoryRepo) *createCategoryBiz {
	return &createCategoryBiz{
		gen:  gen,
		repo: repo,
	}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *categorymodel.Category) error {
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
