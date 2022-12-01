package structs

import "fmt"

type Stack[T any] struct {
	values []T
}

func (self *Stack[T]) Push(value T) {
	self.values = append(self.values, value)
}

func (self *Stack[T]) Pop() (T, error) {
	if len(self.values) == 0 {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	top := self.values[len(self.values)-1]
	self.values = self.values[:len(self.values)-1]
	return top, nil
}
