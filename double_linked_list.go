package main

import (
	"algo/structs"
)

func doulbe_linked_list() {
	list := structs.DoubleLinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Prepend(-1)
	list.Prepend(-2)
	list.Prepend(-3)
	list.Dump()
	list.DumpReverse()
}
