package suppliermodel

import "book-store-management-backend/common"

type ReqUpdateDebtSupplier struct {
	Id             *string `json:"id" gorm:"-" example:"supplier debt id"`
	QuantityUpdate *int    `json:"qtyUpdate" gorm:"-" example:"10000"`
	CreatedBy      string  `json:"-" gorm:"-"`
}

func (*ReqUpdateDebtSupplier) TableName() string {
	return common.TableSupplier
}

func (data *ReqUpdateDebtSupplier) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrSupplierDebtIdInvalid
	}
	if data.QuantityUpdate == nil {
		return ErrSupplierDebtPayNotExist
	}
	if *data.QuantityUpdate == 0 {
		return ErrSupplierDebtPayIsInvalid
	}
	return nil
}
