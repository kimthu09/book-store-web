package suppliermodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
	"errors"
)

type Supplier struct {
	Id          string                           `json:"id" gorm:"column:id;"`
	Name        string                           `json:"name" gorm:"column:name;"`
	Email       string                           `json:"email" gorm:"column:email;"`
	Phone       string                           `json:"phone" gorm:"column:phone;"`
	Debt        float32                          `json:"debt" gorm:"column:debt;"`
	DebtHistory []supplierdebtmodel.SupplierDebt `json:"debtHistory"`
}

func (*Supplier) TableName() string {
	return common.TableSupplier
}

func (s *Supplier) Len() int {
	return len(s.DebtHistory)
}

func (s *Supplier) Less(i, j int) bool {
	return s.DebtHistory[i].CreateAt.After(*s.DebtHistory[j].CreateAt)
}

func (s *Supplier) Swap(i, j int) {
	s.DebtHistory[i], s.DebtHistory[j] = s.DebtHistory[j], s.DebtHistory[i]
}

var (
	ErrSupplierIdInvalid = common.NewCustomError(
		errors.New("id of supplier is invalid"),
		"id of supplier is invalid",
		"ErrSupplierIdInvalid",
	)
	ErrSupplierNameEmpty = common.NewCustomError(
		errors.New("name of supplier is empty"),
		"name of supplier is empty",
		"ErrSupplierNameEmpty",
	)
	ErrSupplierPhoneInvalid = common.NewCustomError(
		errors.New("phone of supplier is invalid"),
		"phone of supplier is invalid",
		"ErrSupplierPhoneInvalid",
	)
	ErrSupplierEmailInvalid = common.NewCustomError(
		errors.New("email of supplier is invalid"),
		"email of supplier is invalid",
		"ErrSupplierEmailInvalid",
	)
	ErrSupplierDebtPayNotExist = common.NewCustomError(
		errors.New("debt pay is not exist"),
		"debt pay is not exist",
		"ErrSupplierDebtPayNotExist",
	)
	ErrSupplierDebtPayIsInvalid = common.NewCustomError(
		errors.New("debt pay is invalid"),
		"debt pay is invalid",
		"ErrSupplierDebtPayIsInvalid",
	)
	ErrSupplierDebtPayIsOverCurrentDebt = common.NewCustomError(
		errors.New("debt pay is over"),
		"debt pay is over",
		"ErrSupplierDebtPayIsOverCurrentDebt",
	)
	ErrSupplierIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of supplier is duplicate"),
	)
	ErrSupplierPhoneDuplicate = common.ErrDuplicateKey(
		errors.New("phone of supplier is duplicate"),
	)
	ErrSupplierCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create supplier"),
	)
	ErrSupplierPayNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to pay supplier"),
	)
	ErrSupplierUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update info supplier"),
	)
	ErrSupplierViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view supplier"),
	)
)
