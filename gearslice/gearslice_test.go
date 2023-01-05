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

func TestUniq(t *testing.T) {
	is := assert.New(t)

	s := []int{1, 2, 3, 4, 4, 5}
	r := Uniq(s)

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

func TestFindIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s := []int{0, 1, 2, 3, 4, 5}
	r1 := FindIndex(s, func(el int, _ int) bool {
		return el > 0
	})
	r2 := FindIndex(s, func(el int, _ int) bool {
		return el > 100
	})

	is.Equal(r1, 1)
	is.Equal(r2, -1)
}

func TestPop(t *testing.T) {
	is := assert.New(t)

	s := []int{1, 2, 3, 4}
	r := Pop(&s)

	is.Equal(r, 4)
	is.Equal(s, []int{1, 2, 3})
}

func TestShift(t *testing.T) {
	is := assert.New(t)

	s := []int{1, 2, 3, 4}
	r := Shift(&s)

	is.Equal(r, 1)
	is.Equal(s, []int{2, 3, 4})
}

func TestInsert(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1, s2 := []int{1, 2, 3}, []int{2, 3, 4}

	r1 := Insert(&s1, 1, 20)
	r2 := Insert(&s2, 3, 20)
	r3 := Insert(&s2, -1, 20)

	is.Nil(r1)
	is.Equal(s1, []int{1, 20, 2, 3})
	is.Error(r2)
	is.Error(r3)
	is.Equal(s2, s2)
}

func TestRemove(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1, s2 := []int{1, 2, 3}, []int{2, 3, 4}

	r1 := Remove(&s1, 1)
	r2 := Remove(&s2, 3)
	r3 := Remove(&s2, -1)

	is.Nil(r1)
	is.Equal(s1, []int{1, 3})
	is.Error(r2)
	is.Error(r3)
	is.Equal(s2, s2)
}

func TestDrop(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s := []int{1, 2, 3, 4, 5, 6, 7}
	r1, err1 := Drop(s, 2)
	r2, err2 := Drop(s, -2)
	r3, err3 := Drop(s, 10)

	is.Equal(r1, []int{3, 4, 5, 6, 7})
	is.Nil(err1)
	is.Equal(r2, []int{1, 2, 3, 4, 5})
	is.Nil(err2)
	is.Empty(r3)
	is.Error(err3)
}
