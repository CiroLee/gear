package geardate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	is := assert.New(t)
	var timestamp int64 = 1673259412 // 2023-01-09 18:16:52
	r := Format(timestamp, DefaultLayout)

	is.Equal(r, "2023-01-09 18:16:52")
}

func TestIsLeap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	years := []int{1999, 2001, 2100, 2400}
	is.False(IsLeap(years[0]))
	is.False(IsLeap(years[1]))
	is.False(IsLeap(years[2]))
	is.True(IsLeap(years[3]))

}
