package data

type Queue struct {
	items []interface{}
	max   int
}

func NewQueue(max int) *Queue {
	return &Queue{items: make([]interface{}, 0, max), max: max}
}

func (q *Queue) Push(item interface{}) bool {
	if q.max != 0 && len(q.items) >= q.max {
		return false
	}
	q.items = append([]interface{}{item}, q.items...)
	return true
}

func (q *Queue) Pop() interface{} {
	if len(q.items) == 0 {
		return nil
	}

	item := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return item
}

func (q *Queue) Peek() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[len(q.items)-1]
}

func (q *Queue) Size() int {
	return len(q.items)
}
