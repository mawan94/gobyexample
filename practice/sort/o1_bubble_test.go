package sort

import (
	"fmt"
	"testing"
)

func bubbleSort(array []int) {
	ok := true
	for i := 1; i < len(array); i++ {
		for j := 0; j < len(array)-i; j++ {
			if array[j] > array[j+1] {
				ok = false
				Swap(array, j, j+1)
			}
		}
		if ok {
			break
		}
		ok = true
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4, -5, 3}
	bubbleSort(arr)
	fmt.Println(arr)
}
