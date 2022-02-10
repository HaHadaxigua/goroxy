package containers

import "roxy/utils"

// Container is base interface that all data structures implement.
type Container[V any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []V
}

// GetSortedValues returns sorted container's elements with respect to the passed comparator.
// Does not effect the ordering of elements within the container.
func GetSortedValues[V comparable](container Container[V], comparator utils.Comparator[V]) []V {
	values := container.Values()
	utils.Sort(values, comparator)
	return values
}
