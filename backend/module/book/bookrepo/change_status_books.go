package bookrepo

import (
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

type ChangeStatusBookStorage interface {
	UpdateStatusBook(
		ctx context.Context,
		id string,
		data *bookmodel.ReqUpdateStatusBook) error
}

type changeStatusBookRepo struct {
	bookStore ChangeStatusBookStorage
}

func NewChangeStatusBookRepo(
	bookStore ChangeStatusBookStorage) *changeStatusBookRepo {
	return &changeStatusBookRepo{
		bookStore: bookStore,
	}
}

func (repo *changeStatusBookRepo) UpdateStatusBooks(
	ctx context.Context,
	data *bookmodel.ReqUpdateStatusBooks) error {
	for _, v := range data.BookIds {
		updateModel := bookmodel.ReqUpdateStatusBook{
			BookId:   v,
			IsActive: data.IsActive,
		}
		if err := repo.bookStore.UpdateStatusBook(ctx, v, &updateModel); err != nil {
			return err
		}
	}

	return nil
}
