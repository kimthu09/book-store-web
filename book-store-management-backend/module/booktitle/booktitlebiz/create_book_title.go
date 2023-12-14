package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

type CreateBookTitleRepo interface {
	CreateBookTitle(ctx context.Context, data *booktitlemodel.BookTitle) error
}

type authorRepo interface {
	IsExistAuthorId(ctx context.Context, authorId string) bool
}

type publisherRepo interface {
	IsExistPublisherId(ctx context.Context, publisherId string) bool
}

type categoryRepo interface {
	IsExistCategoryId(ctx context.Context, categoryId string) bool
}

type createBookTitleBiz struct {
	gen          generator.IdGenerator
	repo         CreateBookTitleRepo
	authorRepo   authorRepo
	categoryRepo categoryRepo
}

func NewCreateBookTitleBiz(
	gen generator.IdGenerator,
	repo CreateBookTitleRepo,
	authorRepo authorRepo,
	categoryRepo categoryRepo) *createBookTitleBiz {
	return &createBookTitleBiz{
		gen:          gen,
		repo:         repo,
		authorRepo:   authorRepo,
		categoryRepo: categoryRepo,
	}
}

func (biz *createBookTitleBiz) CreateBookTitle(ctx context.Context, reqData *booktitlemodel.ReqCreateBookTitle, resData *booktitlemodel.ResCreateBookTitle) error {
	data := &booktitlemodel.BookTitle{
		ID:          nil,
		Name:        reqData.Name,
		Description: reqData.Description,
		AuthorIDs:   reqData.AuthorIDs,
		CategoryIDs: reqData.CategoryIDs,
	}
	if reqData.Id != "" {
		data.ID = &reqData.Id
	}

	if err := data.Validate(); err != nil {
		return err
	}
	data.AuthorIDs = common.RemoveDuplicateStringValues(data.AuthorIDs)
	if err := validateAuthors(ctx, biz.authorRepo, data.AuthorIDs); err != nil {
		return err
	}
	data.CategoryIDs = common.RemoveDuplicateStringValues(data.CategoryIDs)
	if err := validateCategories(ctx, biz.categoryRepo, data.CategoryIDs); err != nil {
		return err
	}
	if err := handleBookTitleId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CreateBookTitle(ctx, data); err != nil {
		return err
	}
	resData.Id = *data.ID
	return nil
}

func handleBookTitleId(gen generator.IdGenerator, data *booktitlemodel.BookTitle) error {
	if data.ID != nil && *data.ID != "" {
		return nil
	}

	id, err := gen.IdProcess(data.ID)
	if err != nil {
		return err
	}
	data.ID = id
	return nil
}
