package userrepo

import (
	"book-store-management-backend/module/user/usermodel"
	"context"
)

type FindUserStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}

type seeUserDetailRepo struct {
	userStore FindUserStore
}

func NewSeeUserDetailRepo(
	userStore FindUserStore) *seeUserDetailRepo {
	return &seeUserDetailRepo{
		userStore: userStore,
	}
}

func (biz *seeUserDetailRepo) SeeUserDetail(
	ctx context.Context,
	userId string) (*usermodel.User, error) {
	user, errUser := biz.userStore.FindUser(
		ctx, map[string]interface{}{"id": userId}, "Role.RoleFeatures")
	if errUser != nil {
		return nil, errUser
	}

	return user, nil
}
