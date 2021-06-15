package sort

import (
	"fmt"
	"testing"
)

/*
	归并排序：稳定性排序 （合并时候左边优先即可）
	核心思想：分而治之。将目标容器拆解成一段段小单位容器。保证每一次小容器有序后递归弹栈时自然整个容器就有序了。
*/
func mergeSort(arr []int) {
	sub(arr, 0, len(arr)-1, make([]int, len(arr)))
}

func sub(arr []int, b, e int, tmp []int) {
	if b >= e {
		return
	}
	m := b + (e-b)>>1
	sub(arr, b, m, tmp)
	sub(arr, m+1, e, tmp)
	merge(arr, b, m, e, tmp)
}

func merge(arr []int, b, m, e int, tmp []int) {
	left, right := b, m+1
	tmpIndex := 0

	// 相同位置进行对齐比价
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
	// 剩余的一边直接追加剩余元素到tmp中
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
	// 将有序片段容器tmp覆盖到原先集合对应的位置
	for i := 0; i < tmpIndex; i++ {
		arr[b+i] = tmp[i]
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4, -5, 3}
	mergeSort(arr)
	fmt.Println(arr)
}
