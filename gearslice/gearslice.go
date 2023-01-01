package gear

import "golang.org/x/exp/constraints"

/*
 * like javascript Array's methods
 */

// IndexOf return the index of the element in the gearslice
func IndexOf[T comparable](s []T, el T) int {
	for k, v := range s {
		if v == el {
			return k
		}
	}
	return -1
}

// Includes weather the gearslice contains a certain element
func Includes[T comparable](s []T, el T) bool {
	for _, v := range s {
		if v == el {
			return true
		}
	}
	return false
}

// Every weather all elements in the gearslice passed the test implemented by the provided function
func Every[T any](s []T, fn func(el T, index int) bool) bool {
	for i, v := range s {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

// Some weather at least one element in the gearslice passed the test implemented by the provide function
func Some[T any](s []T, fn func(el T, index int) bool) bool {
	for i, v := range s {
		if fn(v, i) {
			return true
		}
	}
	return false
}

// Deduplicate remove duplicate elements in the gearslice
func Deduplicate[T constraints.Ordered | string](s []T) []T {
	r := make([]T, 0)
	m := make(map[T]bool) // value maybe any type
	for _, v := range s {
		if _, ok := m[v]; !ok {
			r = append(r, v)
			m[v] = true
		}
	}
	return r
}

// Map create a new gearslice populated with the results of calling the provide function on every element in the calling gearslice
func Map[T, K any](s []T, cb func(el T, index int) K) []K {
	var r = make([]K, 0, len(s))
	for i, v := range s {
		r = append(r, cb(v, i))
	}
	return r
}

// ForEach execute a provided function once for each element in the gearslice
func ForEach[T any](s []T, cb func(el T, index int)) {
	for i, v := range s {
		cb(v, i)
	}
}

// Filter filtered down to just the elements from the given gearslice that passed the test implemented by the provided function
func Filter[T any](s []T, filter func(el T, index int) bool) []T {
	var r = make([]T, 0, len(s))
	for i, v := range s {
		if filter(v, i) {
			r = append(r, v)
		}
	}
	return r
}
