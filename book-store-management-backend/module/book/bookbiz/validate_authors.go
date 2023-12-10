package bookbiz

import (
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

func validateAuthors(ctx context.Context, repo AuthorRepo, authorIDs []string) error {
	for _, id := range authorIDs {
		if ok := repo.IsExistAuthorId(ctx, id); !ok {
			return bookmodel.ErrBookValidateAuthor
		}
	}
	return nil
}
