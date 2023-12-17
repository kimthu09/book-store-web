package userbiz

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type GetAllUserRepo interface {
	GetAllUser(
		ctx context.Context) ([]usermodel.SimpleUser, error)
}

type getAllUserBiz struct {
	repo GetAllUserRepo
}

func NewGetAllUserBiz(
	repo GetAllUserRepo) *getAllUserBiz {
	return &getAllUserBiz{repo: repo}
}

func (biz *getAllUserBiz) GetAllUser(
	ctx context.Context) ([]usermodel.SimpleUser, error) {
	result, err := biz.repo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
