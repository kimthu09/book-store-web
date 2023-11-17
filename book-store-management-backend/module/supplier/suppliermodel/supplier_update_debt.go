package suppliermodel

import "book-store-management-backend/common"

type SupplierUpdateDebt struct {
	Amount   *float32 `json:"amount" gorm:"-"`
	CreateBy string   `json:"-" gorm:"-"`
}

func (*SupplierUpdateDebt) TableName() string {
	return common.TableSupplier
}

func (data *SupplierUpdateDebt) Validate() *common.AppError {
	if data.Amount == nil {
		return ErrSupplierDebtPayNotExist
	}
	if *data.Amount == 0 {
		return ErrSupplierDebtPayIsInvalid
	}
	return nil
}
