package lists

import (
	"github.com/HaHadaxigua/goroxy/utils"
)

// List interface that all lists implement
type List[V any] interface {
	Get(index int) (V, bool)
	Remove(index int)
	Add(values ...V)
	Contains(values ...V) bool
	Sort(comparator utils.Comparator[V])
	Swap(index1, index2 int)
	Insert(index int, values ...V)
	Set(index int, value V)

	//containers.Container[V]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
