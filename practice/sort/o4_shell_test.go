package sort

import (
	"fmt"
	"testing"
)

func shellSort(arr []int) {
	gap := len(arr) >> 1
	for ; gap > 0; gap >>= 1 {
		for j := gap; j < len(arr); j++ {
			prev := j - gap
			for prev >= 0 && arr[prev] > arr[prev+gap] {
				Swap(arr, prev, prev+gap)
				prev--
			}
		}
	}
}

func TestShellSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4, -5, 3}
	shellSort(arr)
	fmt.Println(arr)
}
