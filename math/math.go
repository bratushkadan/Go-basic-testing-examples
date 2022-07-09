package math

import "errors"

var ZeroArgError = errors.New("arg is zero")
var NegativeArgError = errors.New("arg is negative")

var ZeroDivisionError = errors.New("zero division")

const (
	EstimateValueNegative = "negative"
	EstimatedValueSmall   = "small"
	EstimatedValueMedium  = "medium"
	EstimatedValueLarge   = "large"
)

func Add(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, ZeroArgError
	}

	if a < 0 || b < 0 {
		return 0, NegativeArgError
	}
	return a + b, nil //
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 1, ZeroDivisionError
		// return 0, ZeroDivisionError
	}
	return a / b, nil
}

func EstimateValue(value int) string {
	switch {
	case value < 0:
		return EstimateValueNegative
	case value < 10:
		return EstimatedValueSmall

	case value < 100:
		return EstimatedValueMedium

	default:
		return EstimatedValueLarge
	}
}
