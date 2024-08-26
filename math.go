package main

import "golang.org/x/exp/constraints"

func AbsoluteValue[T constraints.Integer](integerValue T) T {
	if integerValue < 0 {
		return -integerValue
	}
	return integerValue
}
