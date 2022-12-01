package main

import (
	"algo/structs"
	"fmt"
)

func linked_list() {
	list := structs.LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	fmt.Println(list.InsertAt(1, 4))
	list.Dump()
	fmt.Println("index =", list.Find(3))
}
