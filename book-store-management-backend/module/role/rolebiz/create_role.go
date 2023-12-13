package rolebiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/role/rolemodel"
	"context"
)

type CreateRoleRepo interface {
	CreateRole(
		ctx context.Context,
		data *rolemodel.ReqCreateRole,
	) error
	CreateRoleFeatures(
		ctx context.Context,
		roleId string,
		featureIds []string,
	) error
}

type createRoleStore struct {
	gen       generator.IdGenerator
	repo      CreateRoleRepo
	requester middleware.Requester
}

func NewCreateRoleStore(
	gen generator.IdGenerator,
	repo CreateRoleRepo,
	requester middleware.Requester) *createRoleStore {
	return &createRoleStore{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createRoleStore) CreateRole(
	ctx context.Context,
	data *rolemodel.ReqCreateRole) error {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return rolemodel.ErrRoleCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleRoleId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CreateRole(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.CreateRoleFeatures(ctx, data.Id, data.Features); err != nil {
		return err
	}

	return nil
}

func handleRoleId(gen generator.IdGenerator, data *rolemodel.ReqCreateRole) error {
	id, err := gen.GenerateId()
	if err != nil {
		return err
	}

	data.Id = id
	return nil
}
