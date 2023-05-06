package geardate

import (
	"testing"
	"time"

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

func TestTimeOffset(t *testing.T) {
	date := time.Date(2023, time.May, 6, 12, 0, 0, 0, time.UTC)

	t.Run("Valid offset", func(t *testing.T) {
		t.Parallel()

		offset := "1h30m"
		offset_after := "-1h30m"
		expectedDate := time.Date(2023, time.May, 6, 13, 30, 0, 0, time.UTC)
		expectedDateAfter := time.Date(2023, time.May, 6, 10, 30, 0, 0, time.UTC)
		result, err := TimeOffset(date, offset)
		result2, err2 := TimeOffset(date, offset_after)
		assert.NoError(t, err)
		assert.Equal(t, expectedDate, result)
		assert.NoError(t, err2)
		assert.Equal(t, expectedDateAfter, result2)
	})

	t.Run("Invalid offset", func(t *testing.T) {
		offset := "invalid"
		expectedDate := time.Time{}
		result, err := TimeOffset(date, offset)
		assert.Error(t, err)
		assert.Equal(t, expectedDate, result)
	})
}
