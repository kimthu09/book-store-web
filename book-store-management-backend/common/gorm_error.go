package common

import (
	"encoding/json"
	"strings"
)

const GormDuplicateErrorNumber = 1062

type GormErr struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}

func GetGormErr(appErr error) *GormErr {
	var byteErr, _ = json.Marshal(appErr)
	var newError GormErr
	err := json.Unmarshal(byteErr, &newError)
	if err != nil {
		return nil
	}
	return &newError
}

func (gErr *GormErr) GetDuplicateErrorKey(args ...string) string {
	if gErr == nil {
		return ""
	}
	if gErr.Number != GormDuplicateErrorNumber {
		return ""
	}
	for _, v := range args {
		if strings.Contains(gErr.Message, v) {
			return v
		}
	}
	return ""
}
