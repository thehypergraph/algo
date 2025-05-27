package heaps

import (
	"container/heap"
	"fmt"
	"sync"
)

// Item represents an element in the min heap with a priority and optional data.
type Item struct {
	Priority int         // The priority value (lower is higher priority)
	Data     interface{} // Optional data associated with the item
}

// MinHeap is a thread-safe min heap implementation.
type MinHeap struct {
	items []*Item      // Slice of Item pointers
	lock  sync.RWMutex // Read-write mutex for thread safety
}

// NewMinHeap creates and initializes a new MinHeap.
func NewMinHeap() *MinHeap {
	mh := &MinHeap{
		items: make([]*Item, 0),
	}
	heap.Init(mh) // Initialize the heap
	return mh
}

// Len returns the number of items in the heap (required for heap.Interface).
func (mh *MinHeap) Len() int {
	mh.lock.RLock()
	defer mh.lock.RUnlock()
	return len(mh.items)
}

// Less compares two items by priority (required for heap.Interface).
func (mh *MinHeap) Less(i, j int) bool {
	mh.lock.RLock()
	defer mh.lock.RUnlock()
	return mh.items[i].Priority < mh.items[j].Priority
}

// Swap exchanges two items in the heap (required for heap.Interface).
func (mh *MinHeap) Swap(i, j int) {
	mh.lock.Lock()
	defer mh.lock.Unlock()
	mh.items[i], mh.items[j] = mh.items[j], mh.items[i]
}

// Push adds a new item to the heap (required for heap.Interface).
func (mh *MinHeap) Push(x interface{}) {
	mh.lock.Lock()
	defer mh.lock.Unlock()
	item, ok := x.(*Item)
	if !ok {
		panic("Push: expected *Item")
	}
	fmt.Println("Pushing item", item.Priority, item.Data)
	mh.items = append(mh.items, item)
}

// Pop removes and returns the item with the smallest priority (required for heap.Interface).
func (mh *MinHeap) Pop() interface{} {
	mh.lock.Lock()
	defer mh.lock.Unlock()
	n := len(mh.items)
	if n == 0 {
		return nil
	}
	item := mh.items[n-1]
	mh.items = mh.items[:n-1]
	return item
}

// PushItem adds a new item with the given priority and data to the heap.
func (mh *MinHeap) PushItem(priority int, data interface{}) {
	item := &Item{
		Priority: priority,
		Data:     data,
	}
	fmt.Println("Pushing item", priority, data)
	heap.Push(mh, item)
}

// PopItem removes and returns the item with the smallest priority.
func (mh *MinHeap) PopItem() *Item {
	item := heap.Pop(mh)
	if item == nil {
		return nil
	}
	return item.(*Item)
}

// Peek returns the item with the smallest priority without removing it.
func (mh *MinHeap) Peek() *Item {
	mh.lock.RLock()
	defer mh.lock.RUnlock()
	if len(mh.items) == 0 {
		return nil
	}
	return mh.items[0]
}

// Size returns the number of items in the heap.
func (mh *MinHeap) Size() int {
	return mh.Len()
}

func (mh *MinHeap) Print() {
	for _, item := range mh.items {
		fmt.Println(item.Priority, item.Data)
	}
}
