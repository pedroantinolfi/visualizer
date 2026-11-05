package queue

import (
	"reflect"
	"testing"
)

func TestNewListQueueIsEmpty(t *testing.T) {
	q := NewListQueue[int]()
	if !q.IsEmpty() {
		t.Error("expected newly created queue to be empty")
	}
}

func TestListQueueEnqueueDequeue(t *testing.T) {
	q := NewListQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("expected 1, got %d", val)
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 2 {
		t.Errorf("expected 2, got %d", val)
	}
}

func TestListQueueDequeueWhenEmpty(t *testing.T) {
	q := NewListQueue[int]()
	_, err := q.Dequeue()
	if err == nil {
		t.Error("expected error when dequeuing from empty queue")
	}
}

func TestListQueueFront(t *testing.T) {
	q := NewListQueue[string]()
	q.Enqueue("a")
	q.Enqueue("b")

	val, err := q.Front()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "a" {
		t.Errorf("expected 'a', got %s", val)
	}

	// Front should not remove the element
	val, err = q.Front()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "a" {
		t.Errorf("expected 'a', got %s", val)
	}
}

func TestListQueueFrontWhenEmpty(t *testing.T) {
	q := NewListQueue[int]()
	_, err := q.Front()
	if err == nil {
		t.Error("expected error when calling Front on empty queue")
	}
}

// TestListQueueEmptyReset ensures the internal pointers (front and back)
// are properly handled when the queue is emptied and then reused.
func TestListQueueEmptyReset(t *testing.T) {
	q := NewListQueue[int]()
	q.Enqueue(10)
	_, _ = q.Dequeue() // Queue is empty again

	if !q.IsEmpty() {
		t.Error("expected queue to be empty after dequeuing all elements")
	}

	// Enqueue again to ensure the 'back' pointer was properly reset
	q.Enqueue(20)
	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("unexpected error after re-enqueuing: %v", err)
	}
	if val != 20 {
		t.Errorf("expected 20, got %d", val)
	}
}

// TestListQueueGetValues verifies that the bridge method for WebAssembly
// correctly exports the elements in order.
func TestListQueueGetValues(t *testing.T) {
	q := NewListQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	expected := []any{10, 20, 30}
	values := q.GetValues()

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("expected %v, got %v", expected, values)
	}
}