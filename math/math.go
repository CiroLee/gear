package gear

import "golang.org/x/exp/constraints"

// Sum return a sum of the slice
func Sum[T constraints.Integer | constraints.Float | constraints.Complex](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
