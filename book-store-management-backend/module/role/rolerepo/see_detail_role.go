package rolerepo

import (
	"book-store-management-backend/module/feature/featuremodel"
	"book-store-management-backend/module/role/rolemodel"
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

type FindRoleStore interface {
	FindRole(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*rolemodel.ResSeeDetailRole, error)
}

type seeRoleDetailRepo struct {
	roleStore        FindRoleStore
	roleFeatureStore ListRoleFeaturesStore
	featureStore     ListAllFeatures
}

func NewSeeRoleDetailRepo(
	roleStore FindRoleStore,
	roleFeatureStore ListRoleFeaturesStore,
	featureStore ListAllFeatures) *seeRoleDetailRepo {
	return &seeRoleDetailRepo{
		roleStore:        roleStore,
		roleFeatureStore: roleFeatureStore,
		featureStore:     featureStore,
	}
}

func (biz *seeRoleDetailRepo) SeeRoleDetail(
	ctx context.Context,
	roleId string) (*rolemodel.ResSeeDetailRole, error) {
	role, errRole := biz.roleStore.FindRole(
		ctx, map[string]interface{}{"id": roleId})
	if errRole != nil {
		return nil, errRole
	}

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
			GroupName:   v.GroupName,
			IsHas:       mapHasFeature[v.Id],
		}
		featureDetails = append(featureDetails, featureDetail)
	}

	var roleDetail rolemodel.ResSeeDetailRole
	roleDetail.Id = role.Id
	roleDetail.Name = role.Name
	roleDetail.Data = featureDetails

	return &roleDetail, nil
}
