package main

import (
	"algo/structs"
	"fmt"
)

func queue_list() {
	q := structs.ListQueue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Dump()

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	q.Dump()

	q.Enqueue(4)
	q.Dump()

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	q.Dump()
}
