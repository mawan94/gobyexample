package sort

import (
	"fmt"
	"testing"
)

/*
	堆排序： 不具备稳定性
	核心思想：构建一个大顶堆 每次可以确认的是堆顶的元素是最大的
*/

func heapSort(arr []int) {
	if arr == nil || len(arr) == 1 {
		return
	}

	// 方式1
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}

	// 方式2
	//for i := 0; i < len(arr); i++ {
	//	heapify(arr, len(arr)-1-i, len(arr)-1)
	//}

	lastIndex := len(arr) - 1
	for lastIndex >= 0 {
		heapify(arr, 0, lastIndex)
		Swap(arr, 0, lastIndex) // 将最大的元素推到后面
		lastIndex--
	}
}

func heapInsert(arr []int, insertIndex int) {
	for arr[insertIndex] > arr[(insertIndex-1)/2] {
		Swap(arr, (insertIndex-1)/2, insertIndex)
		insertIndex = (insertIndex - 1) / 2
	}
}

func heapInsertRecursion(arr []int, insertIndex int) {
	rootIndex := (insertIndex - 1) / 2
	if rootIndex == insertIndex {
		return
	}
	if arr[rootIndex] < arr[insertIndex] {
		Swap(arr, rootIndex, insertIndex)
		heapInsertRecursion(arr, rootIndex)
	}
}

// =========================================================

func heapifyRecursion(arr []int, index, lastIndex int) {
	l := 2*index + 1
	r := l + 1
	// 合法范围检测
	if l <= lastIndex {
		largestIndex := l
		if r <= lastIndex && arr[r] > arr[largestIndex] {
			largestIndex = r
		}
		largestIndex, _ = If(arr[index] > arr[largestIndex], index, largestIndex).(int)
		if largestIndex != index {
			Swap(arr, largestIndex, index)
			heapifyRecursion(arr, largestIndex, lastIndex)
		}
	}
}
func heapify(arr []int, index, lastIndex int) {
	l := index*2 + 1
	// 左孩子合法检测
	for l <= lastIndex {
		// 找到最大值
		largest, _ := If(l+1 <= lastIndex && arr[l+1] > arr[l], l+1, l).(int)
		largest = If(arr[largest] > arr[index], largest, index).(int)

		if largest == index {
			break
		}
		Swap(arr, largest, index)
		index = largest // 状态转移
		l = index*2 + 1
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{9, 4, 6, 7, 2, 11, -2, 6, -2, 4, 0, -5, 3, -1, 20}
	heapSort(arr)
	fmt.Println(arr)
}
