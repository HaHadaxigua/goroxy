package generic_tools

// ----------------------------------------------------------------------------- convert

func Slice2Map[K comparable, V any](keyFn func(v V) K, data []V) map[K]V {
	out := make(map[K]V)
	for _, e := range data {
		out[keyFn(e)] = e
	}
	return out
}

// ----------------------------------------------------------------------------- check

func SliceContainsItem[V comparable](target V, data []V) bool {
	for _, e := range data {
		if e == target {
			return true
		}
	}
	return false
}
