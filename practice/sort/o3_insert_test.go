package sort

import (
	"fmt"
	"testing"
)

/*
	插入排序	（具备稳定性）
	核心思想：把容器看成两截，左边是排好序的部分，右边是待排序的部分。从右半部分第一个元素开始，依次将元素放到左边部分的合适位置、
*/
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		prev := i - 1
		// 满足条件就把当前元素一层一层往前推
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
