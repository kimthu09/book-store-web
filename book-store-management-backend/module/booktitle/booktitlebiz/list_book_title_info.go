package booktitlebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

type ListBookTitleRepo interface {
	ListBookTitle(ctx context.Context, filter *booktitlemodel.Filter, paging *common.Paging) ([]booktitlemodel.BookTitle, error)
}

type listBookTitleBiz struct {
	repo ListBookTitleRepo
}

func NewListBookTitleBiz(repo ListBookTitleRepo) *listBookTitleBiz {
	return &listBookTitleBiz{repo: repo}
}

func (biz *listBookTitleBiz) ListBookTitle(ctx context.Context, filter *booktitlemodel.Filter, paging *common.Paging) ([]booktitlemodel.BookTitle, error) {
	result, err := biz.repo.ListBookTitle(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
