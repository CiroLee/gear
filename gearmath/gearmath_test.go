package gearmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	is := assert.New(t)

	s := []int{1, 2, 3, 4, 5}
	r := Sum(s)

	is.Equal(r, 15)
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
