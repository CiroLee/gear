package gear

import gear "github.com/CiroLee/gear/slice"

// Pick return parts of a map depend on keys
func Pick[K comparable, V any](m map[K]V, keys []K) map[K]V {
	var r = make(map[K]V, len(keys))
	for _, k := range keys {
		r[k] = m[k]
	}
	return r
}

// PickBy return parts of a map depend on custom function
func PickBy[K comparable, V any](m map[K]V, fn func(k K, v V) bool) map[K]V {
	var r = make(map[K]V, 0)
	for k, v := range m {
		if fn(k, v) {
			r[k] = v
		}
	}
	return r
}

// Omit remove parts of a map depend on keys
func Omit[K comparable, V any](m map[K]V, keys []K) map[K]V {
	var r = make(map[K]V, len(m)-len(keys))
	for k, v := range m {
		if gear.IndexOf(keys, k) == -1 {
			r[k] = v
		}
	}
	return r
}

// Values return values of a map
func Values[K comparable, V any](m map[K]V) []V {
	var values = make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Keys return keys of a map
func Keys[K comparable, V any](m map[K]V) []K {
	var keys = make([]K, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
