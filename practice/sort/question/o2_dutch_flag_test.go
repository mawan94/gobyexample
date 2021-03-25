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
		if arr[i] < randPivotValue {
			l++
			sort.Swap(arr, l, i)
			i++
		} else if arr[i] > randPivotValue {
			r--
			sort.Swap(arr, r, i)
		} else {
			i++
		}
	}
}

func TestDutchFlag(t *testing.T) {
	arr := []int{ 6, 6, 8, 9, 0, 12, 4, 7, 5, 6}
	dutchFlag(arr)
	fmt.Println(arr)
}
