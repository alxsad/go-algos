package sort

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](a []T) []T {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := len(a) / 2

	a[pivot], a[right] = a[right], a[pivot]
	fmt.Println(a)

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])

	return a
}
