package sort

import (
	"golang.org/x/exp/constraints"
)

func BubbleSort1[T constraints.Ordered](slice []T) int {
	swaps := true
	iters := 0

	for swaps {
		swaps = false
		for i := 0; i < len(slice)-1; i++ {
			iters++
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swaps = true
			}
		}
	}

	return iters
}

func BubbleSort2[T constraints.Ordered](slice []T) int {
	iters := 0
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			iters++
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return iters
}
