package gearmath

import (
	"math"

	"github.com/CiroLee/gear/gearslice"
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

// SumBy summarize the values in the slice using the given return value from the function
func SumBy[T any, V constraints.Integer | constraints.Float](s []T, fn func(el T, index int) V) V {
	var r V
	for i, v := range s {
		r += fn(v, i)
	}
	return r
}

// Min return the minimum value of the slice
// return zero value if the slice is empty
func Min[T constraints.Integer | constraints.Float](s []T) T {
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
func Max[T constraints.Integer | constraints.Float](s []T) T {
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

// IsSubset return true if the slice contains all the elements in the subset
func IsSubset[T comparable](s, subset []T) bool {
	var b = true
	if len(subset) > len(s) {
		b = false
	} else {
		for _, v := range subset {
			if !gearslice.Includes(s, v) {
				b = false
				break
			}
		}
	}

	return b
}

// Union return the union values of slices
func Union[T constraints.Ordered | constraints.Complex](args ...[]T) []T {
	return gearslice.Uniq(gearslice.Contact(args...))
}
