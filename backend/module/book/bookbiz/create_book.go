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

type createBookBiz struct {
	gen       generator.IdGenerator
	repo      CreateBookRepo
	requester middleware.Requester
}

func NewCreateBookBiz(
	gen generator.IdGenerator,
	repo CreateBookRepo,
	requester middleware.Requester) *createBookBiz {
	return &createBookBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createBookBiz) CreateBook(ctx context.Context, reqData *bookmodel.ReqCreateBook, resData *bookmodel.ResCreateBook) error {
	if !biz.requester.IsHasFeature(common.BookCreateFeatureCode) {
		return bookmodel.ErrBookCreateNoPermission
	}

	if err := reqData.Validate(); err != nil {
		return err
	}
	data := &bookmodel.Book{
		ID:          reqData.ID,
		Image:       reqData.Image,
		BookTitleID: reqData.BookTitleID,
		PublisherID: reqData.PublisherID,
		Edition:     reqData.Edition,
		ListedPrice: reqData.ListedPrice,
		SellPrice:   reqData.SellPrice,
	}
	if err := handleBookId(biz.gen, data); err != nil {
		return bookmodel.ErrBookIdInvalid
	}

	if err := biz.repo.CreateBook(ctx, data); err != nil {
		return err
	}

	resData.Id = *data.ID
	return nil
}

func handleBookId(gen generator.IdGenerator, data *bookmodel.Book) error {
	if data.ID != nil {
		return nil
	}

	id, err := gen.IdProcess(data.ID)
	if err != nil {
		return err
	}
	data.ID = id
	return nil
}
