package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

type CreateBookTitleRepo interface {
	CreateBook(ctx context.Context, data *booktitlemodel.BookTitle) error
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
	gen           generator.IdGenerator
	repo          CreateBookTitleRepo
	authorRepo    authorRepo
	publisherRepo publisherRepo
	categoryRepo  categoryRepo
	requester     middleware.Requester
}

func NewCreateBookTitleBiz(
	gen generator.IdGenerator,
	repo CreateBookTitleRepo,
	authorRepo authorRepo,
	publisherRepo publisherRepo,
	categoryRepo categoryRepo,
	requester middleware.Requester) *createBookTitleBiz {
	return &createBookTitleBiz{
		gen:           gen,
		repo:          repo,
		authorRepo:    authorRepo,
		publisherRepo: publisherRepo,
		categoryRepo:  categoryRepo,
		requester:     requester,
	}
}

func (biz *createBookTitleBiz) CreateBookTitle(ctx context.Context, reqData *booktitlemodel.ReqCreateBookTitle, resData *booktitlemodel.ResCreateBookTitle) error {
	if !biz.requester.IsHasFeature(common.BookTitleCreateFeatureCode) {
		return booktitlemodel.ErrBookTitleCreateNoPermission
	}

	data := &booktitlemodel.BookTitle{
		ID:          nil,
		Name:        reqData.Name,
		Description: reqData.Description,
		AuthorIDs:   reqData.AuthorIDs,
		CategoryIDs: reqData.CategoryIDs,
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
	if err := handleBookId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CreateBook(ctx, data); err != nil {
		return err
	}
	resData.Id = *data.ID
	return nil
}

func handleBookId(gen generator.IdGenerator, data *booktitlemodel.BookTitle) error {
	id, err := gen.IdProcess(data.ID)
	if err != nil {
		return err
	}
	data.ID = id
	return nil
}
