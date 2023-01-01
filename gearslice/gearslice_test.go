package gear

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s := []string{"a", "b", "c"}
	r1 := IndexOf(s, "d")
	r2 := IndexOf(s, "b")

	is.Equal(r1, -1)
	is.Equal(r2, 1)
}

func TestIncludes(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s := []string{"a", "b", "c"}
	r1 := Includes(s, "d")
	r2 := Includes(s, "a")

	is.False(r1)
	is.True(r2)
}

func TestEvery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s := []int{1, 2, 3, -1}
	r1 := Every(s, func(el int, _ int) bool {
		return el > 0
	})
	r2 := Every(s, func(el int, _ int) bool {
		return el > -2
	})

	is.False(r1)
	is.True(r2)
}

func TestSome(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s := []int{1, 2, 3, -1}

	r1 := Some(s, func(el int, _ int) bool {
		return el > 0
	})
	r2 := Some(s, func(el int, _ int) bool {
		return el > 10
	})

	is.True(r1)
	is.False(r2)
}

func TestDeduplicate(t *testing.T) {
	is := assert.New(t)

	s := []int{1, 2, 3, 4, 4, 5}
	r := Deduplicate(s)

	is.Equal(r, []int{1, 2, 3, 4, 5})
}

func TestMap(t *testing.T) {
	is := assert.New(t)

	s := []int{1, 2, 3, 4, 5}
	r := Map(s, func(el int, _ int) int {
		return el * 2
	})

	is.Equal(r, []int{2, 4, 6, 8, 10})
}

func TestForEach(t *testing.T) {
	is := assert.New(t)

	s := []int{1, 2, 3, 4, 5}
	var r = make([]string, 0, len(s))
	ForEach(s, func(el int, _ int) {
		r = append(r, fmt.Sprintf("%d", el))
	})

	is.Equal(s, []int{1, 2, 3, 4, 5})
	is.Equal(r, []string{"1", "2", "3", "4", "5"})
}

func TestFilter(t *testing.T) {
	is := assert.New(t)

	s := []int{1, 2, 3, 4, -1}
	r := Filter(s, func(el int, _ int) bool {
		return el > 0
	})

	is.Equal(r, []int{1, 2, 3, 4})
}
