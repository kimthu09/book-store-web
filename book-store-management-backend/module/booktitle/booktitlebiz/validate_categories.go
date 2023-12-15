package booktitlebiz

import (
	"book-store-management-backend/module/booktitle/booktitlemodel"
	"book-store-management-backend/module/category/categoryrepo"
	"context"
)

func validateCategories(ctx context.Context, repo categoryrepo.CategoryPublicRepo, categoryIDs []string) error {
	for _, id := range categoryIDs {
		if ok := repo.IsExistCategoryId(ctx, id); !ok {
			return booktitlemodel.ErrBookTitleValidateCategory
		}
	}
	return nil

}
