package lists

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Create(data int) *LinkedList {
	head := new(Node)
	head.Data = data
	head.Next = nil
	list.Head = head

	return list
}

func (list *LinkedList) Add(data int) {
	node := new(Node)
	node.Data = data
	node.Next = nil

	if list.Head == nil {
		list.Head.Next = node
	} else {
		cur := list.Head

		for cur.Next != nil {
			cur = cur.Next
		}

		cur.Next = node
	}
}

func (list *LinkedList) Remove(data int) bool {
	cur := list.Head
	prev := new(Node)

	for cur != nil {
		if cur.Data == data {
			prev.Next = cur.Next
			return true
		}

		prev = cur
		cur = cur.Next
	}

	return false
}
