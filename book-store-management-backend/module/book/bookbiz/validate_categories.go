package bookbiz

import (
	"book-store-management-backend/module/book/bookmodel"
	"context"
)

func validateCategories(ctx context.Context, repo categoryRepo, categoryIDs []string) error {
	for _, id := range categoryIDs {
		if ok := repo.IsExistCategoryId(ctx, id); !ok {
			return bookmodel.ErrBookValidateCategory
		}
	}
	return nil

}
