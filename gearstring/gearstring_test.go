package gearstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestContact(t *testing.T) {
	is := assert.New(t)
	r := Contact("hello ", "world")
	is.Equal(r, "hello world")
}

func TestUppercase(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := "hello world"
	s2 := "你好world"

	r1 := ToUpperCase(s1)
	r2 := ToUpperCase(s2)

	is.Equal(r1, "Hello world")
	is.Equal(r2, s2)
}

func TestToLowerCase(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := "HELLO WORLD"
	s2 := "你好world"

	r1 := ToLowerCase(s1)
	r2 := ToLowerCase(s2)

	is.Equal(r1, "hELLO WORLD")
	is.Equal(r2, s2)
}
func TestEncryptedPhone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	phone1 := "1329988998" // 10 digits
	phone2 := "13299889988"

	r1, err1 := EncryptedPhone(phone1)
	r2, err2 := EncryptedPhone(phone2)

	is.Equal(r1, "")
	is.Error(err1)
	is.Equal(r2, "132****9988")
	is.Nil(err2)
}
