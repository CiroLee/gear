package gearstring

import "strings"

/*
* like javascript String methods
 */

// SubString returns the part of the string from the start and excluding the end, or to the end of the string if no end index is supplied
// ps: not included the end index element
func SubString(str string, start, end int) string {
	var runeStr = []rune(str)
	_start, _end := start, end
	if start < 0 {
		_start = 0
	}
	if end < 0 {
		_end = 0
	}

	if _start >= len(runeStr) {
		_start = len(runeStr)
	}

	if _end >= len(runeStr) {
		_end = len(runeStr)
	}

	if _start > _end {
		_start, _end = _end, _start
	}

	if _start == _end {
		return ""
	}

	return string(runeStr[_start:_end])
}

// CharAt returns a specified character from a string
func CharAt(str string, index int) string {
	var runeStr = []rune(str)
	if index > len(runeStr)-1 || index < 0 {
		return ""
	}

	return string(runeStr[index])
}

// Contact Concatenate multiple strings and return a new string
func Contact(args ...string) string {
	var builder strings.Builder
	for _, v := range args {
		builder.WriteString(v)
	}
	return builder.String()
}

func lowerOrUpper(s, t string) string {
	var r []rune
	var rs = []rune(s)
	for i, v := range rs {
		if i == 0 {
			var f string
			if t == "upper" {
				f = strings.ToUpper(string(v))
			} else if t == "lower" {
				f = strings.ToLower(string(v))
			}
			r = append(r, []rune(f)[0])
		} else {
			r = append(r, v)
		}
	}
	return string(r)
}

// ToUpperCase change the first letter of the gearstring to upper
func ToUpperCase(s string) string {
	return lowerOrUpper(s, "upper")
}

// ToLowerCase change the first letter of the gearstring to lower
func ToLowerCase(s string) string {
	return lowerOrUpper(s, "lower")
}
