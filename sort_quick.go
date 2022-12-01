package main

import (
	"algo/sort"
	"fmt"
)

func sort_quick() {
	array := []int{11, 44, 22, 8, 18, 6, 43, 19, 5}
	fmt.Println(array)
	result := sort.QuickSort(array)
	fmt.Println(result)
}
