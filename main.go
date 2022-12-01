package main

import (
	"fmt"
	"os"
)

func main() {
	example := os.Args[1]

	switch example {
	case "stack":
		stack()
	case "queue_slice":
		queue_slice()
	case "queue_list":
		queue_list()
	case "linked_list":
		linked_list()
	case "double_linked_list":
		doulbe_linked_list()
	case "sort_bubble":
		sort_bubble()
	case "sort_quick":
		sort_quick()
	case "binary_three":
		binary_three()
	case "hashmap":
		hashmap()
	default:
		fmt.Println("please select an example")
	}
}
