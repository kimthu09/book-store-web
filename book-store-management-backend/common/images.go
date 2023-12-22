package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Images []Image

func (Images) TableName() string { return TableImage }

func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal data from DB:", value))
	}

	var imgs []Image
	if err := json.Unmarshal(bytes, &imgs); err != nil {
		return err
	}

	*j = imgs
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
