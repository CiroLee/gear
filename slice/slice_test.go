package gear

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
