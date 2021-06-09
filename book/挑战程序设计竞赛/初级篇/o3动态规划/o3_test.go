package o3动态规划

import "math"

/*
  2.3.1 背包问题
*/
var w, v []int
var n int
var MAX_W int

func Solve2_3_1_V1(i, j int) int {
	res := 0
	if i == n { // 没有剩余物品
		res = 0
	} else if j < w[i] { // 拿不动(尝试看一下个物品)
		res = Solve2_3_1_V1(i+1, j)
	} else {
		// 拿和不拿都尝试一下
		res = int(math.Max(float64(Solve2_3_1_V1(i+1, j)), float64(Solve2_3_1_V1(i+1, j-w[i]))+float64(v[i])))
	}
	return res
}

func Solve2_3_1_V2() {
	// 用一个二维数组保存一个已经处理计算过的分支  避免重复计算
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, MAX_W+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1 // 初始化
		}
	}
	var f func(i, j int) int
	f = func(i, j int) int {
		if dp[i][j] >= 0 {
			return dp[i][j]
		}
		res := 0
		if i == n {
			res = 0
		} else if j < w[i] {
			res = f(i+1, j)
		} else {
			res = int(math.Max(float64(f(i+1, j)), float64(f(i+1, j-w[i]))+float64(v[i])))
		}
		dp[i][j] = res
		return res
	}
	f(0, MAX_W)
}

//dp[n][j] = 0  PS : 第n个返回的价值是0  应为越界了（这里也是初始值）
//MAX{ dp[i + 1][j] , dp[i + 1][j - w[i]] + v[i]}
func Solve2_3_1_V3() {
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, MAX_W+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= MAX_W; j++ {
			if j < w[i] {
				dp[i][j] = dp[i+1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i+1][j-w[i]]+v[i])))
			}
		}
	}
}

// 换一种思想
// dp[i + 1][j] := 从前i个物品中选出总重量不超过j的物品时总价最大的值
// dp[0][j] = 0

// dp[i + 1][j] = dp[i][j] (j < w[i])         MAX{dp[i][j],dp[i][j - w[i]] + v[i]}
func Solve2_3_1_V4() {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j <= MAX_W; j++ {
			if j < w[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = int(math.Max(float64(dp[i][j]), float64(dp[i][j-w[i]]+v[i])))
			}
		}
	}
}
