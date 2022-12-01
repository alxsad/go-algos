package main

import (
	"algo/structs"
	"fmt"
)

func hashmap() {
	m := structs.HashMap[int]{}
	m.Set("k1", 1)
	m.Set("k2", 2)
	m.Set("k3", 3)
	m.Dump()
	fmt.Println(m.Get("fake"))
	fmt.Println(m.Get("k3"))
}
