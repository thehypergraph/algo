package main

import "series"
import "list"
import "core:fmt"

main :: proc() {

	fmt.println(series.fib(10))

	ll := list.create(1)
	list.insert(&ll, 2)
	// list.insert(&ll, 3)
	// list.print(&ll)
}
