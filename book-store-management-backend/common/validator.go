package common

func ValidateEmptyString(s string) bool {
	return s == ""
}

func ValidatePositiveNumber(number interface{}) bool {
	switch v := number.(type) {
	case int:
		return v > 0
	case int8:
		return v > 0
	case int16:
		return v > 0
	case int64:
		return v > 0
	case float32:
		return v > 0
	case float64:
		return v > 0
	default:
		return false
	}
}
