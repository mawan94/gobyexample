package bitoperation

import (
	"fmt"
	"testing"
)

//一个数组中有两种数出现了奇数次，其他数都出现了偶数次，怎么找到这两个数
func oddEvenTimes(arr []int) {
	tmp := 0 // a ^ b
	for i := 0; i < len(arr); i++ {
		tmp ^= arr[i]
	}

	// 找到a和b不同的最低位
	diff := tmp & (^tmp + 1)

	// 找出其中的一个数
	oneOfThem := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]&diff == 0 {
			oneOfThem ^= arr[i]
		}
	}
	fmt.Printf("【 %d , %d 】 \n", oneOfThem, tmp^oneOfThem)
}

func TestOddEvenTimes(t *testing.T) {
	oddEvenTimes([]int{2, 2, 2, 2, -5, -5, 1, 3, 3, -5})
}
