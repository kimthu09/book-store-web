package bookbiz

import (
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

type GetAllBookRepo interface {
	GetAllBook(
		ctx context.Context) ([]bookmodel.ResUnitBook, error)
}

type getAllBookBiz struct {
	repo GetAllBookRepo
}

func NewGetAllBookBiz(
	repo GetAllBookRepo) *getAllBookBiz {
	return &getAllBookBiz{
		repo: repo,
	}
}

func (biz *getAllBookBiz) GetAllBook(
	ctx context.Context) ([]bookmodel.ResUnitBook, error) {
	books, err := biz.repo.GetAllBook(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
