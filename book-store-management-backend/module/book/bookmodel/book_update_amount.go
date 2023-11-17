package bookmodel

import (
	"book-store-management-backend/common"
)

type BookUpdateAmount struct {
	AmountUpdate float32 `json:"amountUpdate" gorm:"-"`
}

func (*BookUpdateAmount) TableName() string {
	return common.TableBook
}

func (data *BookUpdateAmount) Validate() *common.AppError {
	if data.AmountUpdate == 0 {
		return ErrBookAmountUpdateInvalid
	}
	return nil
}
