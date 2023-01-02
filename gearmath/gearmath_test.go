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
