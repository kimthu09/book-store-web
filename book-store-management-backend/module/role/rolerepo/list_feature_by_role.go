package rolerepo

import (
	"book-store-management-backend/module/feature/featuremodel"
	"book-store-management-backend/module/rolefeature/rolefeaturemodel"
	"context"
	"log"
)

type ListRoleFeaturesStore interface {
	FindListFeatures(
		ctx context.Context,
		roleId string,
	) ([]rolefeaturemodel.RoleFeature, error)
}

type ListAllFeatures interface {
	ListFeature(
		ctx context.Context,
	) ([]featuremodel.Feature, error)
}

type listFeatureByRoleRepo struct {
	roleFeatureStore ListRoleFeaturesStore
	featureStore     ListAllFeatures
}

func NewListFeatureByRoleRepo(
	roleFeatureStore ListRoleFeaturesStore,
	featureStore ListAllFeatures) *listFeatureByRoleRepo {
	return &listFeatureByRoleRepo{
		roleFeatureStore: roleFeatureStore,
		featureStore:     featureStore,
	}
}

func (biz *listFeatureByRoleRepo) ListFeatureByRole(
	ctx context.Context,
	roleId string) ([]featuremodel.ResFeatureDetail, error) {
	features, errFeature := biz.featureStore.ListFeature(ctx)
	if errFeature != nil {
		return nil, errFeature
	}
	log.Print(features)

	featuresRoleHas, errRoleFeature :=
		biz.roleFeatureStore.FindListFeatures(ctx, roleId)
	if errRoleFeature != nil {
		return nil, errRoleFeature
	}
	log.Print(featuresRoleHas)

	mapHasFeature := make(map[string]bool)
	for _, v := range features {
		mapHasFeature[v.Id] = false
	}
	for _, v := range featuresRoleHas {
		mapHasFeature[v.FeatureId] = true
	}

	var featureDetails []featuremodel.ResFeatureDetail
	for _, v := range features {
		featureDetail := featuremodel.ResFeatureDetail{
			Id:          v.Id,
			Description: v.Description,
			IsHas:       mapHasFeature[v.Id],
		}
		featureDetails = append(featureDetails, featureDetail)
	}
	log.Print(featureDetails)
	return featureDetails, nil
}
