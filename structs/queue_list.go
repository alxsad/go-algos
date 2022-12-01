package structs

import (
	"fmt"
)

type ListQueue[T any] struct {
	list DoubleLinkedList[T]
}

func (self *ListQueue[T]) Enqueue(value T) {
	self.list.Append(value)
}

func (self *ListQueue[T]) Dequeue() (T, error) {
	head := self.list.DropHead()

	if head == nil {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}

	return head.value, nil
}

func (self *ListQueue[T]) Dump() {
	self.list.Dump()
}
