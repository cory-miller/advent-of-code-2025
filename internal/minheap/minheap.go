package minheap

import (
	"container/heap"
)

var _ heap.Interface = new(MinHeap[any])

// A MinHeap implements heap.Interface and holds Items.
type MinHeap[T any] struct {
	elements []T
	less     func(i, j T) bool
}

func (h MinHeap[T]) Len() int {
	return len(h.elements)
}

func (h MinHeap[T]) Less(i, j int) bool {
	return h.less(h.elements[i], h.elements[j])
}

func (h MinHeap[T]) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *MinHeap[T]) Push(x any) {
	h.elements = append(h.elements, x.(T))
}

func (h *MinHeap[T]) Pop() any {
	old := h.elements
	n := len(old)
	x := old[n-1]
	h.elements = old[0 : n-1]
	return x
}

// New creates a new generic min-heap with a custom comparison function.
func New[T any](less func(a, b T) bool) *MinHeap[T] {
	return &MinHeap[T]{
		elements: make([]T, 0),
		less:     less,
	}
}
