package containers

// Container is base interface that all data structures implement.
type Container[V any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []V
}
