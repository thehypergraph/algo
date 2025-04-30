package list

import "core:fmt"

Node :: struct {
	data: int,
	next: ^Node,
}

List :: struct {
	head: ^Node,
}

create :: proc(data: int) -> List {
	head := Node{data = data, next = nil}
	return List{head = &head}
}

insert :: proc(list: ^List, data: int) -> bool {
	if list == nil {
		return false
	}

	new_node := new(Node)
	new_node.data = data
	new_node.next = nil

	if list.head == nil {
		list.head = new_node
		return true
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = new_node
	return true
}

print :: proc(list: ^List) {
	current := list.head
	for current != nil {
		fmt.println(current.data)
		current = current.next
	}
}
