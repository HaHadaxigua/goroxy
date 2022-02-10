package doublylinkedlist

// Each calls the given function once for each element, passing that element's index and value.
func (list *List[V]) Each(f func(index int, value V)) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (list *List[V]) Map(f func(index int, value V) V) *List[V] {
	newList := &List[V]{}
	iterator := list.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (list *List[V]) Select(f func(index int, value V) bool) *List[V] {
	newList := &List[V]{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (list *List[V]) Any(f func(index int, value V) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (list *List[V]) All(f func(index int, value V) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Find passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (list *List[V]) Find(f func(index int, value V) bool) (index int, value V) {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, *new(V)
}
