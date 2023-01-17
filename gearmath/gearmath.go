package gearmath

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Sum return a sum of the slice
func Sum[T constraints.Integer | constraints.Float | constraints.Complex](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// Min return the minimum value of the slice
// return zero value if the slice is empty
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
// return zero value if the slice is empty
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

// Mean return the mean value of the slice
func Mean(s []float64) float64 {
	sum := Sum(s)
	n := len(s)
	return sum / (float64(n))
}

// IsPrime weather a number is a prime
func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	sqrt := math.Sqrt(float64(num))
	for i := 2; float64(i) <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
