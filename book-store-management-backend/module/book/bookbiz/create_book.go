package bookbiz

import (
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

func (biz *createBookBiz) CreateBook(ctx context.Context, data *bookmodel.Book) error {
	//if !biz.requester.IsHasFeature(common.BookCreateFeatureCode) {
	//	return bookmodel.ErrBookCreateNoPermission
	//}
	//
	//if err := data.Validate(); err != nil {
	//	return err
	//}
	//
	//if err := handleBookId(biz.gen, data); err != nil {
	//	return err
	//}
	//if err := biz.repo.CreateBook(ctx, data); err != nil {
	//	return err
	//}
	//
	return nil
}

func handleBookId(gen generator.IdGenerator, data *bookmodel.Book) error {
	id, err := gen.IdProcess(&data.ID)
	if err != nil {
		return err
	}
	data.ID = *id
	return nil
}
