package importnotemodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ImportNoteStatus int

const (
	InProgress ImportNoteStatus = iota
	Done
	Cancel
)

var allImportNoteStatus = [3]string{"InProgress", "Done", "Cancel"}

func (importNoteStatus *ImportNoteStatus) String() string {
	return allImportNoteStatus[*importNoteStatus]
}

func parseStrImportNoteStatus(s string) (ImportNoteStatus, error) {
	for i := range allImportNoteStatus {
		if allImportNoteStatus[i] == s {
			return ImportNoteStatus(i), nil
		}
	}
	return ImportNoteStatus(0), errors.New("invalid status string")
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

func (importNoteStatus *ImportNoteStatus) MarshalJSON() ([]byte, error) {
	if importNoteStatus == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", importNoteStatus.String())), nil
}

func (importNoteStatus *ImportNoteStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	importNoteStatusValue, err := parseStrImportNoteStatus(str)

	if err != nil {
		return err
	}

	*importNoteStatus = importNoteStatusValue

	return nil
}
