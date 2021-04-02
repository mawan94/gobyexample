package recursion

import (
	"fmt"
	"testing"
)

// 有一个n阶的楼梯。
//有两种方式爬楼
//1.一次走一步 2.一次走两步
// 爬到n层一共有多少种方式？


// 假设 n = 3
// 只存在两种情况 1.当前在第一阶（再爬2层）   2.从当前在第二阶（再爬1层）
// 达到3层的情况可以看成是 来到第一层的所有情况和来到第二层所有与情况的总和
// f(3) = f(1) + f(2)
func climbTheStairs(n int) int {
	if n <= 0 {
		return 0
	}else if n <= 2 {
		return n
	}
	return climbTheStairs(n - 2) + climbTheStairs(n - 1)
}

func TestClimbTheStairs(t *testing.T)  {
	fmt.Println(climbTheStairs(10))
}