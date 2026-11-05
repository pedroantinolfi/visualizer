package main

import (
	"fmt"
	"syscall/js"

	"github.com/pedroantinolfi/visualizer/internal/queue"
	// "github.com/pedroantinolfi/visualizer/internal/avl" // To be added later
)

// Global instances for each data structure.
// As we add more structures, we instantiate them here so their state persists.
var q = queue.NewListQueue[string]()

// var avlTree = avl.NewTree() // Coming soon

// queueEnqueueWrapper captures the value from JS and inserts it into the Go queue.
func queueEnqueueWrapper(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return nil
	}
	val := args[0].String()
	q.Enqueue(val)
	fmt.Printf("Go [Queue]: Enqueued '%s'\n", val)
	return nil
}

// queueDequeueWrapper extracts a value and returns it to JS.
func queueDequeueWrapper(this js.Value, args []js.Value) any {
	val, err := q.Dequeue()
	if err != nil {
		fmt.Println("Go [Queue] Error:", err)
		return nil
	}
	fmt.Printf("Go [Queue]: Dequeued '%s'\n", val)
	return val
}

// queueGetValuesWrapper returns the complete state of the queue.
func queueGetValuesWrapper(this js.Value, args []js.Value) any {
	return q.GetValues()
}

func main() {
	fmt.Println("Go WebAssembly initialized successfully.")

	// Registering Queue functions in the JavaScript global object (window).
	// We use the "wasmQueue..." prefix to avoid collisions with future data structures.
	js.Global().Set("wasmQueueEnqueue", js.FuncOf(queueEnqueueWrapper))
	js.Global().Set("wasmQueueDequeue", js.FuncOf(queueDequeueWrapper))
	js.Global().Set("wasmQueueGetValues", js.FuncOf(queueGetValuesWrapper))

	// Block the main thread with a channel to keep the Go program running.
	<-make(chan bool)
}
