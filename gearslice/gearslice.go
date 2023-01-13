package gearslice

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

/*
 * like javascript Array's methods
 */

// IndexOf return the index of the element in the slice, if the element is not in the slice, return -1
func IndexOf[T comparable](s []T, el T) int {
	for k, v := range s {
		if v == el {
			return k
		}
	}
	return -1
}

// Includes weather the slice contains a certain element
func Includes[T comparable](s []T, el T) bool {
	for _, v := range s {
		if v == el {
			return true
		}
	}
	return false
}

// Every weather all elements in the slice passed the test implemented by the provided function
func Every[T any](s []T, fn func(el T, index int) bool) bool {
	for i, v := range s {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

// Some weather at least one element in the slice passed the test implemented by the provide function
func Some[T any](s []T, fn func(el T, index int) bool) bool {
	for i, v := range s {
		if fn(v, i) {
			return true
		}
	}
	return false
}

// Uniq remove duplicate elements in the slice
func Uniq[T constraints.Ordered | string](s []T) []T {
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

// Map create a new slice populated with the results of calling the provide function on every element in the calling slice
func Map[T, K any](s []T, cb func(el T, index int) K) []K {
	var r = make([]K, 0, len(s))
	for i, v := range s {
		r = append(r, cb(v, i))
	}
	return r
}

// ForEach execute a provided function once for each element in the slice
func ForEach[T any](s []T, cb func(el T, index int)) {
	for i, v := range s {
		cb(v, i)
	}
}

// Filter filtered down to just the elements from the given slice that passed the test implemented by the provided function
func Filter[T any](s []T, filter func(el T, index int) bool) []T {
	var r = make([]T, 0, len(s))
	for i, v := range s {
		if filter(v, i) {
			r = append(r, v)
		}
	}
	return r
}

// Find Returns the value of the first element of the slice that passed the test function provided
func Find[T any](s []T, fn func(el T, index int) bool) (T, bool) {
	for i, v := range s {
		if fn(v, i) {
			return v, true
		}
	}
	var empty T
	return empty, false
}

// FindIndex returns the index of the first element in the slice that passed the test implemented by the provided function
// returns -1 if no corresponding element is found
func FindIndex[T any](s []T, fn func(el T, index int) bool) int {
	for i, v := range s {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

// Contact Concatenate multiple slices and return a new slice
func Contact[T any](args ...[]T) []T {
	var r = args[0]
	for i := 1; i < len(args); i++ {
		r = append(r, args[i]...)
	}
	return r
}

// Pop removes the last element from a slice and returns that element. it will change the length of the slice
func Pop[T any](s *[]T) T {
	var last = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

// Shift removes the first element from a slice and returns that removed element. This method changes the length of the slice
func Shift[T any](s *[]T) T {
	var first = (*s)[0]
	*s = (*s)[1:]
	return first
}

// Remove remove a value in the slice at a given index. it will modify the origin slice
func Remove[T any](s *[]T, index int) error {
	if index < 0 || index > len(*s)-1 {
		return errors.New("insert: invalid index. index must be less than the length of the slice, and greater than or equal than zero")
	}
	*s = append((*s)[:index], (*s)[index+1:]...)
	return nil
}

// Insert insert a value in the slice at a given index. it will modify the origin slice
func Insert[T any](s *[]T, index int, value T) error {
	if index < 0 || index > len(*s)-1 {
		return errors.New("insert: invalid index. index must be less than the length of the slice, and greater than or equal than zero")
	}
	*s = append((*s)[:index+1], (*s)[index:]...)
	(*s)[index] = value
	return nil
}

// Drop drops n elements from the beginning(if n greater than zero) or the end(if n less than zero) of the slice
func Drop[T any](s []T, n int) ([]T, error) {
	if len(s) <= int(math.Abs(float64(n))) {
		return nil, errors.New("index exceeds the limit of the slice")
	}
	result := make([]T, 0, len(s)-int(math.Abs(float64(n))))
	if n < 0 {
		return append(result, s[0:len(s)-int(math.Abs(float64(n)))]...), nil
	}
	return append(result, s[n:]...), nil

}

// Sample get a random element from the slice
func Sample[T any](s []T) T {
	lens := len(s)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(lens + 1)
	return s[random]
}

// Reverse reverse a slice, return a new slice
func Reverse[T any](s []T) []T {
	size := len(s)
	var r = make([]T, 0, size)
	for i := size - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
