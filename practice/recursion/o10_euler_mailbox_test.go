package recursion

import "testing"

// 欧拉信箱
// 每个人只能寄一份信，也只能收一份信 但是不能自己寄送给自己。
//若有N人，在保证每个人都有寄信和收信的情况下，一共有多少方式？

// a	b	 c	  d 四个人
// 如果a -> b   b -> a 那么就是   f(n - 2)的问题 (两两抵消)
// 如果a -> b   b \> a 那么就是   f(n - 1)的问题 (a需要收一封，b需要寄一封。状态抵消)
// a 可以选择寄给不同的人（不能寄送给自己） => 系数为 （n - 1）
func eulerMailbox(n int) int{
	if n == 1 {
		return 0
	}else if n == 2 {
		return 1
	}else if n == 3 {
		return 2
	}
	return (n - 1) * (eulerMailbox(n - 2) + eulerMailbox(n - 1))
}



func TestEulerMailbox(t *testing.T)  {
	println(eulerMailbox(4))
}