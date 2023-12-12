package booktitlebiz

import (
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"context"
)

func validateAuthors(ctx context.Context, repo authorRepo, authorIDs []string) error {
	for _, id := range authorIDs {
		if ok := repo.IsExistAuthorId(ctx, id); !ok {
			return booktitlemodel.ErrBookTitleValidateAuthor
		}
	}
	return nil
}
