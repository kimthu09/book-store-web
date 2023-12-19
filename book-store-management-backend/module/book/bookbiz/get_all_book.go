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

	for i := 0; i < len(books); i++ {
		if books[i].ImgUrl == "" {
			books[i].ImgUrl = "https://img.freepik.com/premium-vector/book-vector-illustration-cartoon-book-books-hand-draw-isolated_648083-244.jpg?w=2000"
		}
	}
	return books, nil
}
