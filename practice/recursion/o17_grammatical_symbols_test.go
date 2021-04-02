package recursion

import "testing"

//在第一行我们写上一个 0。接下来的每一行，将前一行中的0替换为01，1替换为10。
//
//给定行数 N 和序数 K，返回第 N 行中第 K个字符。（K从1开始）
//例子:
//
//输入: N = 1, K = 1
//输出: 0
//
//输入: N = 2, K = 1
//输出: 0
//
//输入: N = 2, K = 2
//输出: 1
//
//输入: N = 4, K = 5
//输出: 1
//
//解释:
//第一行: 0
//第二行: 01
//第三行: 0110
//第四行: 01101001
//
//注意：
//
//N的范围 [1, 30].
//K的范围 [1, 2^(N-1)].
func grammaticalSymbols(n, k, width int) int {
	if n == 1 {
		return 0
	}
	half := width >> 1
	if k > half {
		// 找到映射
		mappingK := k - half
		return 1 ^ grammaticalSymbols(n-1, mappingK, half)
	} else {
		// 找上一层答案
		return grammaticalSymbols(n-1, k, half)
	}
}

func getWidth(n int) int {
	if n == 1 {
		return 1
	}
	return getWidth(n-1) << 1
}

func TestGrammaticalSymbols(t *testing.T) {
	n := 4
	k := 4
	println(grammaticalSymbols(n, k, getWidth(n)))
}
