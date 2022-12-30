package gear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubstring(t *testing.T) {
	is := assert.New(t)

	str := "hello world"
	r1 := SubString(str, 1, 5)
	r2 := SubString(str, -1, 5)
	r3 := SubString(str, 1, -1)
	r4 := SubString(str, 12, 6)
	r5 := SubString(str, 1, 12)
	r6 := SubString(str, 2, 2)

	is.Equal(r1, "ello")
	is.Equal(r2, "hello")
	is.Equal(r3, "h")
	is.Equal(r4, "world")
	is.Equal(r5, "ello world")
	is.Empty(r6)
}

func TestCharAt(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str := "hello world"

	r1 := CharAt(str, 2)
	r2 := CharAt(str, 13)
	r3 := CharAt(str, -1)

	is.Equal(r1, "l")
	is.Equal(r2, "")
	is.Equal(r3, "")
	
}
