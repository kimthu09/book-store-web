package supplierbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/supplier/suppliermodel"
	"book-store-management-backend/module/supplier/suppliermodel/filter"
	"context"
)

type SeeSupplierImportNoteRepo interface {
	SeeSupplierImportNote(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierImportFilter,
		paging *common.Paging) (*suppliermodel.ResImportNoteSupplier, error)
}

type seeSupplierImportNoteBiz struct {
	repo      SeeSupplierImportNoteRepo
	requester middleware.Requester
}

func NewSeeSupplierImportNoteBiz(
	repo SeeSupplierImportNoteRepo,
	requester middleware.Requester) *seeSupplierImportNoteBiz {
	return &seeSupplierImportNoteBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeSupplierImportNoteBiz) SeeSupplierImportNote(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) (*suppliermodel.ResImportNoteSupplier, error) {
	if !biz.requester.IsHasFeature(common.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	supplier, err := biz.repo.SeeSupplierImportNote(
		ctx, supplierId, filter, paging)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}
