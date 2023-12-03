package suppliermodel

import "book-store-management-backend/common"

type ReqUpdateDebtSupplier struct {
	QuantityUpdate *float32 `json:"qty_update" gorm:"-" example:"10000"`
	CreateBy       string   `json:"-" gorm:"-"`
}

func (*ReqUpdateDebtSupplier) TableName() string {
	return common.TableSupplier
}

func (data *ReqUpdateDebtSupplier) Validate() *common.AppError {
	if data.QuantityUpdate == nil {
		return ErrSupplierDebtPayNotExist
	}
	if *data.QuantityUpdate == 0 {
		return ErrSupplierDebtPayIsInvalid
	}
	return nil
}

func (data *ReqUpdateDebtSupplier) Round() {
	common.CustomRound(data.QuantityUpdate)
}
