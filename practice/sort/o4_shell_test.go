package sort

import (
	"fmt"
	"testing"
)

/*
	希尔排序：（稳定性排序）可以理解成升级版的插入排序。解决了插入排序的痛点（降序转升序）
	核心思想：和插入排序一样 多了一个大方向调整相对次序
*/
func shellSort(arr []int) {
	gap := len(arr)
	for ; gap != 0; gap >>= 1 {
		for j := gap; j < len(arr); j += gap {
			prev := j - gap
			for prev >= 0 && arr[prev+gap] < arr[prev] {
				Swap(arr, prev, prev+gap)
				prev -= gap
			}
		}
	}
}

func TestShellSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4, -5, 3}
	shellSort(arr)
	fmt.Println(arr)
}
