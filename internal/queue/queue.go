package queue

import "errors"

// ErrQueueEmpty is returned when an operation is performed on an empty queue.
var ErrQueueEmpty = errors.New("the queue is empty")

// Queue is the interface that all queues must implement.
type Queue[T any] interface {
	// Enqueue adds an element to the back of the queue.
	Enqueue(val T)
    // Dequeue removes and returns the element at the front.
    // Returns an error if the queue is empty.
	Dequeue() (T, error)
	// Front returns the element at the front without removing it.
    // Returns an error if the queue is empty.
	Front() (T, error)
	// IsEmpty returns true if the queue has no elements.
	IsEmpty() bool
	// GetValues iterates through the queue and returns all elements as a []any slice.
    // This is the bridge method required to pass the data structure to WebAssembly/JavaScript.
	GetValues() []any
}