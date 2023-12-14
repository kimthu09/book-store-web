package authorbiz

import (
	"book-store-management-backend/component/generator"
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type CreateAuthorRepo interface {
	CreateAuthor(ctx context.Context, data *authormodel.Author) error
}

type createAuthorBiz struct {
	gen  generator.IdGenerator
	repo CreateAuthorRepo
}

func NewCreateAuthorBiz(gen generator.IdGenerator, repo CreateAuthorRepo) *createAuthorBiz {
	return &createAuthorBiz{
		gen:  gen,
		repo: repo,
	}
}

func (biz *createAuthorBiz) CreateAuthor(ctx context.Context, data *authormodel.Author) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleAuthorId(biz.gen, data); err != nil {
		return err
	}
	if err := biz.repo.CreateAuthor(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleAuthorId(gen generator.IdGenerator, data *authormodel.Author) error {
	id, err := gen.IdProcess(&data.Id)
	if err != nil {
		return err
	}
	data.Id = *id
	return nil
}
