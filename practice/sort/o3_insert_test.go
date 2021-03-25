package sort

import (
	"fmt"
	"testing"
)

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		prev := i - 1
		for prev >= 0 && arr[prev+1] < arr[prev] {
			Swap(arr, prev, prev+1)
			prev--
		}
	}
}

func TestInsertSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4, -5, 3}
	insertSort(arr)
	fmt.Println(arr)
}


