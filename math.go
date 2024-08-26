package main

import "golang.org/x/exp/constraints"

func Abs[T constraints.Integer](value T) T {
	if value < 0 {
		return -value
	}
	return value
}
