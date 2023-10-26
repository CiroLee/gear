package gearmath

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	is := assert.New(t)

	s := []int{1, 2, 3, 4, 5}
	r := Sum(s)

	is.Equal(r, 15)
}

func TestSumBy(t *testing.T) {
	is := assert.New(t)
	s := []string{"hello", "world"}
	r := SumBy(s, func(el string, _ int) int {
		return len(el)
	})

	is.Equal(r, 10)
}

func TestMin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := []int{1, 2, 3, 4, -4, 5, 6}
	s2 := []int{}

	r1 := Min(s1)
	r2 := Min(s2)

	is.Equal(r1, -4)
	is.Equal(r2, 0)
}

func TestMax(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := []int{1, 2, 3, 4, -4, 5, 6}
	s2 := []int{}

	r1 := Max(s1)
	r2 := Max(s2)

	is.Equal(r1, 6)
	is.Equal(r2, 0)
}

func TestIsPrime(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.False(IsPrime(4))
	is.True(IsPrime(3))
	is.True(IsPrime(2))
	is.False(IsPrime(1))
	is.False(IsPrime(0))
	is.False(IsPrime(-1))
}

func TestMean(t *testing.T) {
	is := assert.New(t)

	s := []float64{2, 4, 6, 8}
	r := Mean(s)

	is.Equal(r, 5.0)
}

func TestIsSubset(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 3}
	s3 := []int{2, 3, 4, 5}

	ss1 := []string{"foo", "bar"}
	ss2 := []string{"foo"}
	r1 := IsSubset(s1, s2)
	r2 := IsSubset(s2, s1)
	r3 := IsSubset(s1, s3)

	rr1 := IsSubset(ss1, ss2)

	is.True(r1)
	is.False(r2)
	is.False(r3)
	is.True(rr1)
}

func TestUnion(t *testing.T) {
	is := assert.New(t)

	s1, s2, s3 := []int{1, 2, 3, 4}, []int{2, 5, 7}, []int{-1, 0, 0}
	r := Union(s1, s2, s3)
	sort.Ints(r)

	is.Equal(r, []int{-1, 0, 1, 2, 3, 4, 5, 7})
}

func TestIntersection(t *testing.T) {
	is := assert.New(t)
	// 测试两个有序的切片的交集
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}
	result := Intersection(slice1, slice2)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	is.Equal(result, []int{3, 4, 5})
}
