package queue

// Queue represents a double-ended queue.
// The zero value is an Empty queue ready to use.
type Queue[V any] struct {
	// Push writes to rep[back] then increments back; PushFront
	// decrements front then writes to rep[front]; len(rep) is a power
	// of two; unused slots are nil and not garbage.
	rep    []V
	front  int
	back   int
	length int
}

// New returns an initialized Empty queue.
func New[V any]() *Queue[V] {
	return new(Queue[V]).Init()
}

// Init initializes or clears queue q.
func (q *Queue[V]) Init() *Queue[V] {
	q.rep = make([]V, 1)
	q.front, q.back, q.length = 0, 0, 0
	return q
}

// lazyInit lazily initializes a zero Queue value.
//
// I am mostly doing this because container/list does the same thing.
// Personally I think it's a little wasteful because every single
// PushFront/Push is going to pay the overhead of calling this.
// But that's the price for making zero values useful immediately.
func (q *Queue[V]) lazyInit() {
	if q.rep == nil {
		q.Init()
	}
}

// Len returns the number of elements of queue q.
func (q *Queue[V]) Len() int {
	return q.length
}

// Empty returns true if the queue q has no elements.
func (q *Queue[V]) Empty() bool {
	return q.length == 0
}

// full returns true if the queue q is at capacity.
func (q *Queue[V]) full() bool {
	return q.length == len(q.rep)
}

// sparse returns true if the queue q has excess capacity.
func (q *Queue[V]) sparse() bool {
	return 1 < q.length && q.length < len(q.rep)/4
}

// resize adjusts the size of queue q's underlying slice.
func (q *Queue[V]) resize(size int) {
	adjusted := make([]V, size)
	if q.front < q.back {
		// rep not "wrapped" around, one copy suffices
		copy(adjusted, q.rep[q.front:q.back])
	} else {
		// rep is "wrapped" around, need two copies
		n := copy(adjusted, q.rep[q.front:])
		copy(adjusted[n:], q.rep[:q.back])
	}
	q.rep = adjusted
	q.front = 0
	q.back = q.length
}

// lazyGrow grows the underlying slice if necessary.
func (q *Queue[V]) lazyGrow() {
	if q.full() {
		q.resize(len(q.rep) * 2)
	}
}

// lazyShrink shrinks the underlying slice if advisable.
func (q *Queue[V]) lazyShrink() {
	if q.sparse() {
		q.resize(len(q.rep) / 2)
	}
}

// inc returns the next integer position wrapping around queue q.
func (q *Queue[V]) inc(i int) int {
	return (i + 1) & (len(q.rep) - 1) // requires l = 2^n
}

// dec returns the previous integer position wrapping around queue q.
func (q *Queue[V]) dec(i int) int {
	return (i - 1) & (len(q.rep) - 1) // requires l = 2^n
}

// Front returns the first element of queue q or nil.
func (q *Queue[V]) Front() V {
	// no need to check q.Empty(), unused slots are nil
	return q.rep[q.front]
}

// Back returns the last element of queue q or nil.
func (q *Queue[V]) Back() V {
	// no need to check q.Empty(), unused slots are nil
	return q.rep[q.dec(q.back)]
}

// Push inserts a new value v at the back of queue q.
func (q *Queue[V]) Push(v V) {
	q.lazyInit()
	q.lazyGrow()
	q.rep[q.back] = v
	q.back = q.inc(q.back)
	q.length++
}

// Pop removes and returns the first element of queue q or nil.
func (q *Queue[V]) Pop() V {
	if q.Empty() {
		return q.nil()
	}
	v := q.rep[q.front]
	q.rep[q.front] = q.nil() // unused slots must be nil
	q.front = q.inc(q.front)
	q.length--
	q.lazyShrink()
	return v
}

func (q *Queue[V]) nil() (_ V) {
	return
}
