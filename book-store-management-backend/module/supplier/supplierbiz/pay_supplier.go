package supplierbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/common/enum"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/supplier/suppliermodel"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"fmt"
)

type PaySupplierStoreRepo interface {
	GetDebtSupplier(
		ctx context.Context,
		supplierId string,
	) (*float32, error)
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate,
	) error
	UpdateDebtSupplier(
		ctx context.Context,
		supplierId string,
		data *suppliermodel.ReqUpdateDebtSupplier,
	) error
}

type paySupplierBiz struct {
	gen       generator.IdGenerator
	repo      PaySupplierStoreRepo
	requester middleware.Requester
}

func NewUpdatePayBiz(
	gen generator.IdGenerator,
	repo PaySupplierStoreRepo,
	requester middleware.Requester) *paySupplierBiz {
	return &paySupplierBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *paySupplierBiz) PaySupplier(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.ReqUpdateDebtSupplier) (*string, error) {
	if !biz.requester.IsHasFeature(common.SupplierPayFeatureCode) {
		return nil, suppliermodel.ErrSupplierPayNoPermission
	}

	if err := validateSupplierUpdateDebt(data); err != nil {
		return nil, err
	}

	data.Round()

	debtCurrent, errGetDebt := biz.repo.GetDebtSupplier(ctx, supplierId)
	if errGetDebt != nil {
		return nil, errGetDebt
	}
	if *debtCurrent+*data.QuantityUpdate > 0 {
		fmt.Println(*debtCurrent)
		fmt.Println(*data.QuantityUpdate)
		return nil, suppliermodel.ErrSupplierDebtPayIsInvalid
	}

	supplierDebtCreate, errGetSupplierDebtCreate := getSupplierDebtCreate(
		biz.gen, supplierId, *debtCurrent, data,
	)
	if errGetSupplierDebtCreate != nil {
		return nil, errGetSupplierDebtCreate
	}

	if err := biz.repo.UpdateDebtSupplier(ctx, supplierId, data); err != nil {
		return nil, err
	}

	if err := biz.repo.CreateSupplierDebt(ctx, supplierDebtCreate); err != nil {
		return nil, err
	}

	return &supplierDebtCreate.Id, nil
}

func validateSupplierUpdateDebt(data *suppliermodel.ReqUpdateDebtSupplier) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if *data.QuantityUpdate <= 0 {
		return suppliermodel.ErrSupplierDebtPayIsInvalid
	}

	return nil
}

func getSupplierDebtCreate(
	gen generator.IdGenerator,
	supplierId string,
	currentDebt float32,
	data *suppliermodel.ReqUpdateDebtSupplier,
) (*supplierdebtmodel.SupplierDebtCreate, error) {
	qtyPay := *data.QuantityUpdate
	qtyLeft := currentDebt + qtyPay

	id, errGenerateId := gen.GenerateId()
	if errGenerateId != nil {
		return nil, errGenerateId
	}

	debtType := enum.Pay
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:           id,
		SupplierId:   supplierId,
		Quantity:     qtyPay,
		QuantityLeft: qtyLeft,
		DebtType:     &debtType,
		CreateBy:     data.CreateBy,
	}

	return &supplierDebtCreate, nil
}
