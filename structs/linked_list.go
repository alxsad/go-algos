package structs

import "fmt"

type linkedNode[T comparable] struct {
	value T
	next  *linkedNode[T]
}

type LinkedList[T comparable] struct {
	head *linkedNode[T]
}

func (self *LinkedList[T]) Add(value T) {
	node := &linkedNode[T]{value: value}
	if self.head == nil {
		self.head = node
		return
	}
	cursor := self.head
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.next = node
}

func (self *LinkedList[T]) Find(value T) int {
	if self.head == nil {
		return -1
	}
	i := 0
	cursor := self.head
	for cursor != nil {
		if cursor.value == value {
			return i
		}
		i++
		cursor = cursor.next
	}
	return -1
}

func (self *LinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 {
		return fmt.Errorf("invalid index")
	}
	if self.head == nil && index != 0 {
		return fmt.Errorf("invalid index")
	}
	node := &linkedNode[T]{value: value}
	i := 0
	cursor := self.head
	prev := self.head

	for prev != nil {
		if i == index {
			if cursor != self.head {
				prev.next = node
			} else {
				self.head = node
			}
			node.next = cursor
			return nil
		}
		i++
		prev = cursor
		if cursor == nil {
			return fmt.Errorf("out of index")
		}
		cursor = cursor.next
	}
	return nil
}

func (self *LinkedList[T]) Dump() {
	if self.head == nil {
		return
	}
	cursor := self.head
	for cursor != nil {
		fmt.Println(cursor.value)
		cursor = cursor.next
	}
}
