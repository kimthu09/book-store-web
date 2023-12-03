package enum

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type DebtType string

const (
	Pay  DebtType = "Pay"
	Debt DebtType = "Debt"
)

var allDebtType = [2]DebtType{"Pay", "Debt"}

func (debtType *DebtType) String() string {
	return string(*debtType)
}

func parseStrDebtType(s string) (DebtType, error) {
	for _, dt := range allDebtType {
		if dt.String() == s {
			return dt, nil
		}
	}
	return "", errors.New("invalid debt type string")
}

func (debtType *DebtType) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	v, err := parseStrDebtType(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*debtType = v

	return nil
}

func (debtType *DebtType) Value() (driver.Value, error) {
	if debtType == nil {
		return nil, nil
	}

	return debtType.String(), nil
}
