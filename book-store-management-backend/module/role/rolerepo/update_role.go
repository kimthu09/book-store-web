package rolerepo

import (
	"book-store-management-backend/module/role/rolemodel"
	"book-store-management-backend/module/rolefeature/rolefeaturemodel"
	"context"
)

type UpdateRoleStore interface {
	UpdateRole(
		ctx context.Context,
		id string,
		data *rolemodel.ReqUpdateRole) error
}

type UpdateRoleFeature interface {
	CreateListRoleFeature(
		ctx context.Context,
		data []rolefeaturemodel.RoleFeature,
	) error
	DeleteRoleFeature(
		ctx context.Context,
		conditions map[string]interface{},
	) error
	FindListFeatures(
		ctx context.Context,
		roleId string,
	) ([]rolefeaturemodel.RoleFeature, error)
}

type updateRoleRepo struct {
	roleStore        UpdateRoleStore
	roleFeatureStore UpdateRoleFeature
}

func NewUpdateRoleRepo(
	roleStore UpdateRoleStore,
	roleFeatureStore UpdateRoleFeature) *updateRoleRepo {
	return &updateRoleRepo{
		roleStore:        roleStore,
		roleFeatureStore: roleFeatureStore,
	}
}

func (repo *updateRoleRepo) GetListRoleFeatures(
	ctx context.Context,
	roleId string) ([]string, error) {
	features, err := repo.roleFeatureStore.FindListFeatures(ctx, roleId)
	if err != nil {
		return nil, err
	}

	var featureListStr []string
	for _, feature := range features {
		featureListStr = append(featureListStr, feature.FeatureId)
	}
	return featureListStr, nil
}

func (repo *updateRoleRepo) UpdateRole(
	ctx context.Context,
	roleId string,
	data *rolemodel.ReqUpdateRole) error {
	if err := repo.roleStore.UpdateRole(ctx, roleId, data); err != nil {
		return err
	}
	return nil
}

func (repo *updateRoleRepo) UpdateRoleFeatures(
	ctx context.Context,
	deletedRoleFeatures []rolefeaturemodel.RoleFeature,
	createdRoleFeatures []rolefeaturemodel.RoleFeature) error {
	if err := repo.deleteRoleFeatures(ctx, deletedRoleFeatures); err != nil {
		return err
	}
	if err := repo.roleFeatureStore.CreateListRoleFeature(
		ctx, createdRoleFeatures,
	); err != nil {
		return err
	}
	return nil
}

func (repo *updateRoleRepo) deleteRoleFeatures(
	ctx context.Context,
	deletedRoleFeatures []rolefeaturemodel.RoleFeature) error {
	for _, v := range deletedRoleFeatures {
		if err := repo.roleFeatureStore.DeleteRoleFeature(
			ctx,
			map[string]interface{}{
				"roleId":    v.RoleId,
				"featureId": v.FeatureId,
			},
		); err != nil {
			return err
		}
	}

	return nil
}
