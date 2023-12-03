package importnotemodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type ImportNoteStatus string

const (
	InProgress ImportNoteStatus = "InProgress"
	Done       ImportNoteStatus = "Done"
	Cancel     ImportNoteStatus = "Cancel"
)

var allImportNoteStatus = [3]ImportNoteStatus{"InProgress", "Done", "Cancel"}

func (importNoteStatus *ImportNoteStatus) String() string {
	return string(*importNoteStatus)
}

func parseStrImportNoteStatus(s string) (ImportNoteStatus, error) {
	for _, value := range allImportNoteStatus {
		if value.String() == s {
			return value, nil
		}
	}
	return InProgress, errors.New("invalid status string")
}

func (importNoteStatus *ImportNoteStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	v, err := parseStrImportNoteStatus(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*importNoteStatus = v

	return nil
}

func (importNoteStatus *ImportNoteStatus) Value() (driver.Value, error) {
	if importNoteStatus == nil {
		return nil, nil
	}

	return importNoteStatus.String(), nil
}
