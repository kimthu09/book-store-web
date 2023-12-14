package booktitlebiz

import (
	"book-store-management-backend/module/author/authormodel"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

type authorPublicRepo interface {
	GetByListId(ctx context.Context, ids []string) ([]authormodel.Author, error)
}

type categoryPublicRepo interface {
	GetByListId(ctx context.Context, ids []string) ([]categorymodel.Category, error)
}
