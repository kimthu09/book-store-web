package enum

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type DebtType int

const (
	Pay DebtType = iota
	Debt
)

var allDebtType = [2]string{"Pay", "Debt"}

func (debtType *DebtType) String() string {
	return allDebtType[*debtType]
}

func parseStrDebtType(s string) (DebtType, error) {
	for i := range allDebtType {
		if allDebtType[i] == s {
			return DebtType(i), nil
		}
	}
	return DebtType(0), errors.New("invalid debt type string")
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

func (debtType *DebtType) MarshalJSON() ([]byte, error) {
	if debtType == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", debtType.String())), nil
}

func (debtType *DebtType) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	debtTypeValue, err := parseStrDebtType(str)

	if err != nil {
		return err
	}

	*debtType = debtTypeValue

	return nil
}
