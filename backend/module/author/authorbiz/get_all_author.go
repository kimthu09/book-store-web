package authorbiz

import (
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type GetAllAuthorRepo interface {
	GetAllAuthor(ctx context.Context) ([]authormodel.SimpleAuthor, error)
}

type getAllAuthorBiz struct {
	repo GetAllAuthorRepo
}

func NewGetAllAuthorBiz(repo GetAllAuthorRepo) *getAllAuthorBiz {
	return &getAllAuthorBiz{repo: repo}
}

func (biz *getAllAuthorBiz) GetAllAuthor(
	ctx context.Context) ([]authormodel.SimpleAuthor, error) {
	result, err := biz.repo.GetAllAuthor(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
