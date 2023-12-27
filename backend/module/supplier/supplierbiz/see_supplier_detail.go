package supplierbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/supplier/suppliermodel"
	"context"
)

type SeeSupplierDetailRepo interface {
	SeeSupplierDetail(
		ctx context.Context,
		supplierId string) (*suppliermodel.Supplier, error)
}

type seeSupplierDetailBiz struct {
	repo      SeeSupplierDetailRepo
	requester middleware.Requester
}

func NewSeeSupplierDetailBiz(
	repo SeeSupplierDetailRepo,
	requester middleware.Requester) *seeSupplierDetailBiz {
	return &seeSupplierDetailBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeSupplierDetailBiz) SeeSupplierDetail(
	ctx context.Context,
	supplierId string) (*suppliermodel.Supplier, error) {
	if !biz.requester.IsHasFeature(common.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	supplier, err := biz.repo.SeeSupplierDetail(
		ctx, supplierId)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}
