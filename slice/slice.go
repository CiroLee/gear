package gear

// IndexOf return the index of the element in the slice
func IndexOf[T comparable](s []T, el T) int {
	for k, v := range s {
		if v == el {
			return k
		}
	}
	return -1
}

// Includes weather the slice contains a certain element
func Includes[T comparable](s []T, el T) bool {
	for _, v := range s {
		if v == el {
			return true
		}
	}
	return false
}
