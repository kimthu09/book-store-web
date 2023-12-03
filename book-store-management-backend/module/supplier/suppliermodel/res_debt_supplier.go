package suppliermodel

import (
	"book-store-management-backend/module/supplierdebt/supplierdebtmodel"
)

type ResDebtSupplier struct {
	Id          string                           `json:"id" gorm:"column:id;" example:"123"`
	Name        string                           `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
	Email       string                           `json:"email" gorm:"column:email;" example:"a@gmail.com"`
	Phone       string                           `json:"phone" gorm:"column:phone;" example:"0123456789"`
	Debt        float32                          `json:"debt" gorm:"column:debt;" example:"-100000"`
	DebtHistory []supplierdebtmodel.SupplierDebt `json:"debt_history"`
}

func GetResSeeDebtSupplierFromSupplier(supplier *Supplier) *ResDebtSupplier {
	var src ResDebtSupplier
	src.Id = supplier.Id
	src.Name = supplier.Name
	src.Email = supplier.Email
	src.Phone = supplier.Phone
	src.Debt = supplier.Debt
	return &src
}
