package sort

import (
	"fmt"
	"testing"
)

func selectionSort(arr []int) {
	var minIndex = 0
	for i := 0; i < len(arr)-1; i++ {
		minIndex = i
		for j := i; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		Swap(arr,i,minIndex)
	}
}


func TestSelectionSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4, -5, 3}
	selectionSort(arr)
	fmt.Println(arr)
}
