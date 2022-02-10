package utils

import "sort"

// Sort sorts values (in-place) with respect to the given comparator.
//
// Uses Go's sort (hybrid of quicksort for large and then insertion sort for smaller slices).
func Sort[V comparable](values []V, comparator Comparator) {
	sort.Sort(sortable[V]{values, comparator})
}

type sortable[V comparable] struct {
	values     []V
	comparator Comparator
}

func (s sortable[V]) Len() int {
	return len(s.values)
}
func (s sortable[V]) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s sortable[V]) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}
