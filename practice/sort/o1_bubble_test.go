package sort

import (
	"fmt"
	"testing"
)

/*
	冒泡排序 （稳定性）
	核心思想： 将当前数与相邻元素（后卫）比较，如果比相邻元素大就往后推。
    这样每轮下来都能确定一个最大的数字。最多 len(arr) - 1轮后就能完成排序
*/
func bubbleSort(array []int) {
	ok := true
	for i := 1; i < len(array); i++ {
		for j := 0; j < len(array)-i; j++ {
			if array[j] > array[j+1] {
				Swap(array, j, j+1)
				ok = false
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
