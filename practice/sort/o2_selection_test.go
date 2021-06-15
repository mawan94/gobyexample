package sort

import (
	"fmt"
	"testing"
)

/*
	选择排序 (不具备稳定性)
	核心思想：从容器第一个元素开始，从后面剩余的元素中挑选一个最小的元素进行交换。
	一轮后就能确定当前剩余元素中最小的元素。 len(arr) - 1 轮后即可完成排序
*/
func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		Swap(arr, i, minIndex)
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4, -5, 3}
	selectionSort(arr)
	fmt.Println(arr)
}
