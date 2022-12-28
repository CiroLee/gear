package gear

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
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
