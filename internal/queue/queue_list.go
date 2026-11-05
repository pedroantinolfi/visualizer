package queue

import "errors"

// node represents a single element in the linked list.
type node[T comparable] struct {
	value T
	next  *node[T]
}

// ListQueue implements a Queue[T] using a linked list.
type ListQueue[T comparable] struct {
	front *node[T]
	back  *node[T]
	size  int
}

// NewListQueue creates a new empty queue.
func NewListQueue[T comparable]() *ListQueue[T] {
	return &ListQueue[T]{} // front and back are implicitly nil, size is 0
}

// Enqueue adds an element to the back of the queue.
func (q *ListQueue[T]) Enqueue(val T) {
	newNode := &node[T]{value: val}
	if q.IsEmpty() {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}
	q.size++
}

// Dequeue removes and returns the element at the front.
// Returns an error if the queue is empty.
func (q *ListQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	
	val := q.front.value
	q.front = q.front.next
	
	// If the queue became empty after dequeue, reset the back pointer
	if q.front == nil {
		q.back = nil
	}
	
	q.size--
	return val, nil
}

// Front returns the element at the front without removing it.
// Returns an error if the queue is empty.
func (q *ListQueue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.front.value, nil
}

// IsEmpty returns true if the queue has no elements.
func (q *ListQueue[T]) IsEmpty() bool {
	return q.size == 0
}

// GetValues iterates through the queue and returns all elements as a []any slice.
// This is the bridge method required to pass the data structure to WebAssembly/JavaScript.
func (q *ListQueue[T]) GetValues() []any {
	var values []any
	current := q.front
	for current != nil {
		values = append(values, current.value)
		current = current.next
	}
	return values
}