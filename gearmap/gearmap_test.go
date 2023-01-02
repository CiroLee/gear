package gearmap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPick(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	r1 := Pick(m, []string{"a", "b"})
	r2 := Pick(m, []string{})

	is.Equal(r1, map[string]int{"a": 1, "b": 2})
	is.Equal(r2, map[string]int{})
}

func TestPickBy(t *testing.T) {
	is := assert.New(t)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	r := PickBy(m, func(_ string, v int) bool {
		return v > 2
	})

	is.Equal(r, map[string]int{"c": 3})
}

func TestValues(t *testing.T) {
	is := assert.New(t)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	r := Values(m)
	sort.Ints(r)

	is.Equal(r, []int{1, 2, 3})
}

func TestKeys(t *testing.T) {
	is := assert.New(t)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	r := Keys(m)
	sort.Strings(r)

	is.Equal(r, []string{"a", "b", "c"})
}

func TestOmit(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	r1 := Omit(m, []string{"a", "b"})
	r2 := Omit(m, []string{})

	is.Equal(r1, map[string]int{"c": 3})
	is.Equal(r2, m)
}

func TestOmitBy(t *testing.T) {
	is := assert.New(t)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	r := OmitBy(m, func(_ string, v int) bool {
		return v < 2
	})

	is.Equal(r, map[string]int{"b": 2, "c": 3})
}

func TestAssign(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"d": 4}
	m3 := map[string]int{"c": 5, "d": 6}

	r1 := Assign(m1, m2)
	r2 := Assign(m1, m3)

	is.Equal(r1, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	is.Equal(r2, map[string]int{"a": 1, "b": 2, "c": 5, "d": 6})
}
