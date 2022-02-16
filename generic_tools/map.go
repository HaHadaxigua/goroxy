package generic_tools

func Map2Slice[K comparable, V any](data map[K]V) []V {
	var out []V
	for _, v := range data {
		out = append(out, v)
	}
	return out
}

func MergeMap[K comparable, V any](dst, src map[K]V) {
	if dst == nil {
		dst = make(map[K]V)
	}
	for k, v := range src {
		dst[k] = v
	}
}

func MapFilter[K comparable, V any](data map[K]V, fn func(k K, v V) bool) map[K]V {
	out := make(map[K]V)
	for k, v := range data {
		if fn(k, v) {
			out[k] = v
		}
	}
	return out
}

func MapCollection[K comparable, V any, R any](data map[K]V, fn func(K, V, R) R) R {
	var out R
	for k, v := range data {
		out = fn(k, v, out)
	}
	return out
}
