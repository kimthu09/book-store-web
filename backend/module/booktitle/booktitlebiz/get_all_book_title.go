package booktitlebiz

import (
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

type GetAllBookTitleRepo interface {
	GetAllBookTitle(ctx context.Context) ([]booktitlemodel.SimpleBookTitle, error)
}

type getAllBookTitleBiz struct {
	repo GetAllBookTitleRepo
}

func NewGetAllBookTitleBiz(repo GetAllBookTitleRepo) *getAllBookTitleBiz {
	return &getAllBookTitleBiz{repo: repo}
}

func (biz *getAllBookTitleBiz) GetAllBookTitle(
	ctx context.Context) ([]booktitlemodel.SimpleBookTitle, error) {
	result, err := biz.repo.GetAllBookTitle(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
