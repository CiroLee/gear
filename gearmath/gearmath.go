package gearmath

import "golang.org/x/exp/constraints"

// Sum return a sum of the slice
func Sum[T constraints.Integer | constraints.Float | constraints.Complex](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// Min return the minimum value of the slice
// return zero value if the slice is tempy
func Min[T constraints.Ordered](s []T) T {
	var min T
	if len(s) == 0 {
		return min
	}
	min = s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

// Max return the minimum value of the slice
// return zero value if the slice is tempy
func Max[T constraints.Ordered](s []T) T {
	var max T
	if len(s) == 0 {
		return max
	}
	max = s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}
