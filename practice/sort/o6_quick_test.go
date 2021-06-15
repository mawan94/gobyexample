package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

/*
	快速排序：不具备稳定性
	核心思想：在容器中先将一个元素保证有序 然后再对其余两头拆分进行递归调用
*/
func quickSort(arr []int) {
	s(arr, 0, len(arr)-1)
}

func s(arr []int, b, e int) {
	if e > b {
		//prev, next := partitionV1(arr, b, e)
		prev, next := partitionV2(arr, b, e)
		s(arr, b, prev)
		s(arr, next, e)
	}
}

//挖坑填数
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

// 挤压挪动
func partitionV2(arr []int, b, e int) (prev, next int) {
	l, r := b-1, e+1                  // 左右边界   =>   ）E1 E2 E3 E4 ...（
	pivotVal := arr[b+rand.Intn(e-b)] // 随机挑选一个中轴值

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
