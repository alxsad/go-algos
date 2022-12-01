package main

import (
	"algo/sort"
	"fmt"
)

func sort_bubble() {
	array1 := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Println(array1)
	iters1 := sort.BubbleSort1(array1)
	fmt.Println(array1, iters1)

	array2 := []int{11, 14, 3, 8, 18, 17, 43}
	iters2 := sort.BubbleSort2(array2)
	fmt.Println(array2, iters2)
}
