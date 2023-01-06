package data

type Stack struct {
	items []interface{}
	max   int
}

func NewStack(max int) *Stack {
	return &Stack{items: make([]interface{}, 0, max), max: max}
}

func (s *Stack) Push(item interface{}) bool {
	if s.max != 0 && len(s.items) >= s.max {
		return false
	}
	s.items = append([]interface{}{item}, s.items...)
	return true
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[0]
}

func (s *Stack) Size() int {
	return len(s.items)
}
