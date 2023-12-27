package common

import "math"

const maxDecimalNumber = 2

func roundToDecimal(num float32, decimalPlaces int) float32 {
	shift := math.Pow10(decimalPlaces)
	return float32(math.Round(float64(num)*shift)) / float32(shift)
}

func CustomRound(num *float32) {
	roundedNum := roundToDecimal(*num, maxDecimalNumber)
	*num = roundedNum
}
