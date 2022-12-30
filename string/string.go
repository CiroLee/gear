package gear

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
