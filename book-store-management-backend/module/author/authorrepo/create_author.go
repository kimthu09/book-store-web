package authorrepo

import (
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type CreateAuthor interface {
	CreateAuthor(ctx context.Context, data *authormodel.Author) error
}

type createAuthorRepo struct {
	store CreateAuthor
}

func NewCreateAuthorRepo(store CreateAuthor) *createAuthorRepo {
	return &createAuthorRepo{store: store}
}

func (biz *createAuthorRepo) CreateAuthor(ctx context.Context, data *authormodel.Author) error {
	if err := biz.store.CreateAuthor(ctx, data); err != nil {
		return err
	}

	return nil
}
