package gearslice

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

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type Person struct {
		Name   string
		Age    int
		Gender int
		Grade  int
	}
	s := []Person{
		{Name: "Tom", Age: 12, Gender: 1, Grade: 2},
		{Name: "Jim", Age: 11, Gender: 1, Grade: 1},
		{Name: "Dave", Age: 13, Gender: 0, Grade: 3},
	}

	r1, ok1 := Find(s, func(el Person, _ int) bool {
		return el.Age > 11 && el.Gender == 0
	})
	r2, ok2 := Find(s, func(el Person, _ int) bool {
		return el.Age > 16 && el.Gender == 0
	})

	is.True(ok1)
	is.Equal(r1,
		Person{Name: "Dave", Age: 13, Gender: 0, Grade: 3},
	)
	is.False(ok2)
	is.Equal(r2, Person{})
}

func TestContact(t *testing.T) {
	is := assert.New(t)

	s1 := []int{1, 2, 3, 4}
	s2 := []int{5, 6}
	s3 := []int{6, 7, 8}

	r := Contact(s1, s2, s3)

	is.Equal(r, []int{1, 2, 3, 4, 5, 6, 6, 7, 8})
}
