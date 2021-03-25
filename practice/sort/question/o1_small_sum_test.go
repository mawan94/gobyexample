package question

import (
	"fmt"
	"testing"
)

//在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。
//求一个数组 的小和。 例子:[1,3,4,2,5]
//1左边比1小的数，没有;
//3左边比3小的数，1;
//4左 边比4小的数，1、3;
//2左边比2小的数，1;
//5左边比5小的数，1、3、4、2; 所以小和为 1+1+3+1+1+3+4+2=16

func smallSum(arr []int) int {
	return getSmallSum(arr, 0, len(arr)-1)
}

func getSmallSum(arr []int, l, r int) int {
	if l < r {
		m := int(uint(l+r) / 2)
		return getSmallSum(arr, l, m) + getSmallSum(arr, m+1, r) + calculation(arr, l, m, r)
	}
	return 0
}

func calculation(arr []int, l, m, r int) int {
	help := make([]int, r-l+1)
	p1, p2 := l, m+1
	total, i := 0, 0
	for p1 <= m && p2 <= r {
		if arr[p1] < arr[p2] {
			total += (r - p2 + 1) * arr[p1]
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	for p1 <= m || p2 <= r {
		if p1 <= m {
			help[i] = arr[p1]
			p1++
		} else if p2 <= r {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	for i := 0; i < len(help); i++ {
		arr[l+i] = help[i]
	}
	return total
}

func TestSmallSum(t *testing.T) {
	fmt.Println(smallSum([]int{1, 3, 4, 2, 5}))
}
