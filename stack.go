package main

import (
	"algo/structs"
	"fmt"
)

func stack() {
	s := structs.Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
