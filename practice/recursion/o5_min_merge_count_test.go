package recursion

import (
	"fmt"
	"math"
	"testing"
)

//假设你开了间小店，不能电子支付，钱柜里的货币只有 25 分、10 分、5 分和 1 分四种硬币，
//如果你是售货员且要找给客户 N 分钱的硬币，如何安排才能找给客人的钱既正确且硬币的个数又最少？

// n 要大于等于 硬币金额才能继续调用  否则return (无法达成条件)
// 每种面值的硬币都可以选择使用或者不使用
// f(41) = min {  f(41 - 25) + 1,  f(41 - 10) + 1 , f(41 - 5) + 1 ,  f(41 - 1) + 1 }

func minMergeCount(n int, coins []int, takeRecord string) (int, string) {
	if n == 0 {
		return 0, takeRecord
	}
	min := 1 << 31 - 1 - 1
	path := ""

	for i := 0; i < len(coins); i++ {
		if n >= coins[i] {
			count, p := minMergeCount(
				n-coins[i],
				coins,
				fmt.Sprintf("%s  %d", takeRecord, coins[i]))
			min = int(math.Min(float64(min), float64(count+1)))
			if min == count+1 {
				path = p
			}
		}
	}
	return min, path
}

func TestMinMergeCount(t *testing.T) {
	coins := []int{25, 10, 5, 1}
	count, s := minMergeCount(41, coins, "")
	fmt.Printf("count:  %d,		 s: %s\n", count, s)
}
