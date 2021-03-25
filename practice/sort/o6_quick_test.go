package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func quickSort(arr []int) {
	s(arr, 0, len(arr)-1)
}

func s(arr []int, b, e int) {
	if e-b >= 1 {
		//prev, next := partitionV1(arr, b, e)
		prev, next := partitionV2(arr, b, e)
		s(arr, next, e)
		s(arr, b, prev)
	}
}

func partitionV1(arr []int, b, e int) (prev, next int) {
	headPoint := b
	tailPoint := e
	headValue := arr[b]

	for headPoint != tailPoint {
		for headPoint != tailPoint && arr[tailPoint] >= headValue {
			tailPoint--
		}
		if headPoint != tailPoint {
			arr[headPoint] = arr[tailPoint]
			headPoint++
		}
		for headPoint != tailPoint && arr[headPoint] <= headValue {
			headPoint++
		}
		if headPoint != tailPoint {
			arr[tailPoint] = arr[headPoint]
			tailPoint--
		}
	}
	arr[headPoint] = headValue
	prev = headPoint - 1
	next = headPoint + 1
	return
}

func partitionV2(arr []int, b, e int) (prev, next int) {
	l := b - 1
	r := e + 1
	pivotVal := arr[b + rand.Intn(e-b)]

	for i := b; i != r; {
		if arr[i] < pivotVal {
			l++
			Swap(arr, i, l)
			i++
		} else if arr[i] == pivotVal {
			i++
		} else {
			r--
			Swap(arr, i, r)
		}
	}
	return l, r
}

func TestQuickSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4, -5, 3}
	quickSort(arr)
	fmt.Println(arr)
}
