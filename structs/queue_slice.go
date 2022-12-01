package structs

import "fmt"

type SliceQueue[T any] struct {
	values []T
}

func (self *SliceQueue[T]) Enqueue(value T) {
	self.values = append(self.values, value)
}

func (self *SliceQueue[T]) Dequeue() (T, error) {
	if len(self.values) == 0 {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	top := self.values[0]
	self.values = self.values[1:len(self.values)]
	return top, nil
}

func (self *SliceQueue[T]) Dump() {
	fmt.Println(self.values)
}
