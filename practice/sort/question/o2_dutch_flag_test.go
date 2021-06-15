package question

import (
	"fmt"
	"gobyexample/practice/sort"
	"testing"
)

//给定一个数组arr，和一个数num，
//请把小于num的数放在数组的 左边，
//等于num的数放 在数组的中间，
//大于num的数放在数组的 右边。
//要求额外空间复杂度O(1)，时间复杂度 O(N)
func dutchFlag(arr []int) {
	//randPivotValue := arr[rand.Intn(len(arr)-1)]
	randPivotValue := arr[0]
	l := -1
	r := len(arr)
	i := 0

	for i != r {
		//左边括号向前移动
		//交换当前元素和移动后的元素
		//指针向后移动
		if arr[i] < randPivotValue {
			l++
			sort.Swap(arr, l, i)
			i++
			// 右边括号向左移动
			// 交换当前元素和移动后的元素
			// 指针不动（因为交换过来的元素还没被看过）
		} else if arr[i] > randPivotValue {
			r--
			sort.Swap(arr, r, i)
		} else {
			i++
		}
	}
}

func TestDutchFlag(t *testing.T) {
	arr := []int{6, 6, 8, 9, 0, 12, 4, 7, 5, 6}
	dutchFlag(arr)
	fmt.Println(arr)
}
