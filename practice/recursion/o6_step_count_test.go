package recursion

import "testing"

// 走格子
// 有一个 n * m 的格子
// 初始化时你站在左上角(0,0)
// 只能向右走或者向下走(每次走一格)
// 移动到右下角(n,m)一共有多少种方式？


func stepCount(n, m int) int {
	if n - 1 == 0 || m - 1  == 0 {
		return 1
	}
	return stepCount(n - 1, m) + stepCount(n, m - 1)
}

func TestStepCount(t *testing.T)  {
	println(stepCount(3, 3))
}

