package bookbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

type CreateBookRepo interface {
	CreateBook(ctx context.Context, data *bookmodel.Book) error
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

type createBookBiz struct {
	gen           generator.IdGenerator
	repo          CreateBookRepo
	authorRepo    authorRepo
	publisherRepo publisherRepo
	categoryRepo  categoryRepo
	requester     middleware.Requester
}

func NewCreateBookBiz(
	gen generator.IdGenerator,
	repo CreateBookRepo,
	authorRepo authorRepo,
	publisherRepo publisherRepo,
	categoryRepo categoryRepo,
	requester middleware.Requester) *createBookBiz {
	return &createBookBiz{
		gen:           gen,
		repo:          repo,
		authorRepo:    authorRepo,
		publisherRepo: publisherRepo,
		categoryRepo:  categoryRepo,
		requester:     requester,
	}
}

func (biz *createBookBiz) CreateBook(ctx context.Context, reqData *bookmodel.ReqCreateBook, resData *bookmodel.ResCreateBook) error {
	if !biz.requester.IsHasFeature(common.BookCreateFeatureCode) {
		return bookmodel.ErrBookCreateNoPermission
	}

	data := &bookmodel.Book{
		ID:          nil,
		Name:        reqData.Name,
		Description: reqData.Description,
		Edition:     reqData.Edition,
		Quantity:    reqData.Quantity,
		ListedPrice: reqData.ListedPrice,
		SellPrice:   reqData.SellPrice,
		PublisherID: reqData.PublisherID,
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
	if err := validatePublisher(ctx, biz.publisherRepo, data.PublisherID); err != nil {
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

func handleBookId(gen generator.IdGenerator, data *bookmodel.Book) error {
	id, err := gen.IdProcess(data.ID)
	if err != nil {
		return err
	}
	data.ID = id
	return nil
}
