package bookbiz

import "context"

type DeleteBookRepo interface {
	DeleteBook(ctx context.Context, id string) error
}

type DeleteRepo interface {
	DeleteBook(ctx context.Context, id string) error
}

type deleteBookBiz struct {
	repo DeleteRepo
}

func NewDeleteBookBiz(repo DeleteRepo) *deleteBookBiz {
	return &deleteBookBiz{repo: repo}
}

func (biz *deleteBookBiz) DeleteBook(ctx context.Context, id string) error {
	return biz.repo.DeleteBook(ctx, id)
}
