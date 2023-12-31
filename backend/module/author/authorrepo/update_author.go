package authorrepo

import (
	"book-store-management-backend/module/author/authormodel"
	"context"
)

type UpdateAuthorStore interface {
	UpdateAuthor(
		ctx context.Context,
		id string,
		data *authormodel.ReqUpdateAuthor) error
}

type updateAuthorRepo struct {
	store UpdateAuthorStore
}

func NewUpdateAuthorRepo(store UpdateAuthorStore) *updateAuthorRepo {
	return &updateAuthorRepo{store: store}
}

func (repo *updateAuthorRepo) UpdateAuthorInfo(
	ctx context.Context,
	supplierId string,
	data *authormodel.ReqUpdateAuthor) error {
	if err := repo.store.UpdateAuthor(ctx, supplierId, data); err != nil {
		return err
	}
	return nil
}
