package main

import (
	"algo/structs"
	"fmt"
)

func binary_three() {
	t := structs.BinaryThree[int]{}
	t.Add(10)
	t.Add(19)
	t.Add(5)
	t.Add(1)
	t.Add(6)
	t.Add(17)
	t.Add(21)
	t.Dump()
	fmt.Println()
	fmt.Println(t.FindMax())
	fmt.Println(t.FindMin())
}
