package bookbiz

import (
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

func validatePublisher(ctx context.Context, repo publisherRepo, publisherId string) error {
	if !repo.IsExistPublisherId(ctx, publisherId) {
		return bookmodel.ErrBookValidatePublisher
	}
	return nil
}
