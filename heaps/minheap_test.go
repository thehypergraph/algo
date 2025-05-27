package heaps

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()
	heap.PushItem(1, "1")
	heap.PushItem(2, "2")
	heap.PushItem(3, "3")

	if peek := heap.Peek(); peek == nil || peek.Priority != 1 {
		t.Errorf("Expected Peek() to return item with priority 1, got %v", peek)
	}

	if pop := heap.PopItem(); pop == nil || pop.Priority != 1 {
		t.Errorf("Expected Pop() to return item with priority 1, got %v", pop)
	}

	if pop := heap.PopItem(); pop == nil || pop.Priority != 2 {
		t.Errorf("Expected Pop() to return item with priority 2, got %v", pop)
	}
}
