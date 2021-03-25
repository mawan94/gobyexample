package sort

import (
	"fmt"
	"testing"
)

func mergeSort(arr []int) {
	sub(arr, 0, len(arr)-1, make([]int, len(arr)))
}

func sub(arr []int, b, e int, tmp []int) {
	if b >= e {
		return
	}
	m := int(uint(b+e) >> 1)
	sub(arr, b, m, tmp)
	sub(arr, m+1, e, tmp)
	merge(arr, b, m, e, tmp)
}

func merge(arr []int, b, m, e int, tmp []int) {
	left, right := b, m+1
	tmpIndex := 0

	for left <= m && right <= e {
		if arr[left] <= arr[right] {
			tmp[tmpIndex] = arr[left]
			left++
		} else {
			tmp[tmpIndex] = arr[right]
			right++
		}
		tmpIndex++
	}

	for left <= m || right <= e {
		if left <= m {
			tmp[tmpIndex] = arr[left]
			left++
		} else {
			tmp[tmpIndex] = arr[right]
			right++
		}
		tmpIndex++
	}

	for i := 0; i < tmpIndex; i++ {
		arr[b+i] = tmp[i]
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4, -5, 3}
	mergeSort(arr)
	fmt.Println(arr)
}

