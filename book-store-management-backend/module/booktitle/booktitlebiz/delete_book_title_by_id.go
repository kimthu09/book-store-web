package booktitlebiz

import (
	"context"
)

type DeleteBookRepo interface {
	DeleteBookTitle(ctx context.Context, id string) error
}

type DeleteRepo interface {
	DeleteBookTitle(ctx context.Context, id string) error
}

type deleteBookBiz struct {
	repo DeleteRepo
}

func NewDeleteBookBiz(repo DeleteRepo) *deleteBookBiz {
	return &deleteBookBiz{
		repo: repo,
	}
}

func (biz *deleteBookBiz) DeleteBookTitle(ctx context.Context, id string) error {
	return biz.repo.DeleteBookTitle(ctx, id)
}
