package structs

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type threeNode[T constraints.Ordered] struct {
	value T
	left  *threeNode[T]
	right *threeNode[T]
}

func (self *threeNode[T]) Add(node *threeNode[T]) {
	if self.value == node.value {
		return
	}
	if self.value > node.value {
		if self.left == nil {
			self.left = node
		} else {
			self.left.Add(node)
		}
	} else {
		if self.right == nil {
			self.right = node
		} else {
			self.right.Add(node)
		}
	}
}

func (self *threeNode[T]) FindMax() (bool, T) {
	if self.right == nil {
		return true, self.value
	}
	return self.right.FindMax()
}

func (self *threeNode[T]) FindMin() (bool, T) {
	if self.left == nil {
		return true, self.value
	}
	return self.left.FindMin()
}

func (self *threeNode[T]) SubThree(acc []T, value T) []T {
	if self.value > value && self.left != nil {
		return self.left.SubThree(acc, value)
	}
	if self.value < value && self.right != nil {
		return self.right.SubThree(acc, value)
	}
	if self.value == value {
		fmt.Println("HUI")
		acc = append(acc, self.value)
	}
	if self.left != nil {
		return self.left.SubThree(acc, value)
	}
	if self.right != nil {
		return self.right.SubThree(acc, value)
	}
	return acc
}

func (self *threeNode[T]) Dump() {
	fmt.Println(self.value)
	if self.left != nil {
		self.left.Dump()
	}
	if self.right != nil {
		self.right.Dump()
	}
}

type BinaryThree[T constraints.Ordered] struct {
	root *threeNode[T]
}

func (self *BinaryThree[T]) Add(value T) {
	node := &threeNode[T]{value: value}
	if self.root == nil {
		self.root = node
		return
	}
	self.root.Add(node)
}

func (self *BinaryThree[T]) FindMax() (bool, T) {
	var zero T
	if self.root == nil {
		return false, zero
	}
	return self.root.FindMax()
}

func (self *BinaryThree[T]) FindMin() (bool, T) {
	var zero T
	if self.root == nil {
		return false, zero
	}
	return self.root.FindMin()
}

func (self *BinaryThree[T]) Dump() {
	if self.root == nil {
		return
	}
	self.root.Dump()
}

//           10
//         /   \
//        5     19
//       / \   / \
//      1   6 17  21
