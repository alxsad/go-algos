package structs

import "fmt"

type doubleLinkedNode[T any] struct {
	value T
	next  *doubleLinkedNode[T]
	prev  *doubleLinkedNode[T]
}

type DoubleLinkedList[T any] struct {
	head *doubleLinkedNode[T]
	tail *doubleLinkedNode[T]
}

func (self *DoubleLinkedList[T]) Append(value T) {
	node := &doubleLinkedNode[T]{value: value}
	if self.head == nil {
		self.head = node
		self.tail = node
		return
	}
	self.tail.next = node
	node.prev = self.tail
	self.tail = node
}

func (self *DoubleLinkedList[T]) Prepend(value T) {
	node := &doubleLinkedNode[T]{value: value}
	if self.head == nil {
		self.head = node
		self.tail = node
		return
	}

	self.head.prev = node
	node.next = self.head
	self.head = node
}

func (self *DoubleLinkedList[T]) DropHead() *doubleLinkedNode[T] {
	if self.head == nil {
		return nil
	}
	head := self.head
	if self.head == self.tail {
		self.head = nil
		self.tail = nil
		return head
	}
	self.head.next.prev = nil
	self.head = self.head.next
	head.next = nil
	return head
}

func (self *DoubleLinkedList[T]) Dump() {
	if self.head == nil {
		return
	}
	cursor := self.head
	for cursor != nil {
		fmt.Printf("%v ", cursor.value)
		cursor = cursor.next
	}
	fmt.Println("")
}

func (self *DoubleLinkedList[T]) DumpReverse() {
	if self.tail == nil {
		return
	}
	cursor := self.tail
	for cursor != nil {
		fmt.Printf("%v ", cursor.value)
		cursor = cursor.prev
	}
	fmt.Println("")
}
