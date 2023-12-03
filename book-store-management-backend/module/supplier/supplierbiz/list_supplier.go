package supplierbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/supplier/suppliermodel"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"context"
)

type ListSupplierRepo interface {
	ListSupplier(
		ctx context.Context,
		filter *filter.Filter,
		paging *common.Paging,
	) ([]suppliermodel.Supplier, error)
}

type listSupplierBiz struct {
	repo      ListSupplierRepo
	requester middleware.Requester
}

func NewListSupplierRepo(
	repo ListSupplierRepo,
	requester middleware.Requester) *listSupplierBiz {
	return &listSupplierBiz{repo: repo, requester: requester}
}

func (biz *listSupplierBiz) ListSupplier(
	ctx context.Context,
	filter *filter.Filter,
	paging *common.Paging) ([]suppliermodel.Supplier, error) {
	if !biz.requester.IsHasFeature(common.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	result, err := biz.repo.ListSupplier(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
