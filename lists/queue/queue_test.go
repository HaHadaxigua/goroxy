package queue

import (
	"fmt"
	"testing"
)

func TestPop(t *testing.T) {
	queue := New[string]()
	queue.Push("hl")
	queue.Push("x")

	for !queue.Empty() {
		fmt.Println(queue.Pop())
	}

	queue.Pop()
}
