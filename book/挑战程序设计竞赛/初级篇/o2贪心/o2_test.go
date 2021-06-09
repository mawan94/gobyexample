package o2贪心

import (
	"fmt"
	"math"
)

/*
	2.2.1 硬币问题
	V 面值切片		C 对应面值的数量		T 目标金额
*/

func solve2_2_1(V, C []int, T int) {
	ans := 0
	for i := len(V); i >= 0; i-- {
		useCount := math.Min(float64(T/V[i]), float64(C[i])) // 使用硬币i的枚数
		T -= int(useCount) * V[i]
		ans += int(useCount)
	}
	println(ans)
}

/*
	2.2.2 区间问题
	N 工作数量
*/
type Job struct {
	B, E int
}

func solve2_2_2(N int, sortedJobs []Job) {
	ans := 0
	t := 0 // 工作的结束时间
	for i := 0; i < N; i++ {
		if t < sortedJobs[i].B {
			ans++
			t = sortedJobs[i].E
		}
	}
	println(ans)
}

/*
	2.2.3 字典最小序
*/
func solve2_2_3(first, last int, S string) {
	ret := ""
	for first <= last {
		left := false
		for i := 0; i+first <= last; i++ {
			if S[first+i] < S[last-i] {
				left = true
				break
			} else if S[first+i] > S[last-i] {
				left = false
				break
			}
		}
		if left {
			ret = fmt.Sprintf("%s%s", string(S[first]), ret)
			first++
		} else {
			ret = fmt.Sprintf("%s%s", string(S[last]), ret)
			last--
		}
	}
}

/*
	2.2.4 Saruman's Army (POJ 3069)
	X : 点的切片
	R : RANGE
	N : 点的数量
*/
func solve2_2_4(sortedX []int, R, N int) {
	i, ans := 0, 0
	for i < N {
		s := sortedX[i]
		i++
		for i < N && sortedX[i] <= s+R {
			i++
		}
		p := sortedX[i-1]
		ans++

		for i < N && sortedX[i] <= p+R {
			i++
		}
	}
	println(ans)
}

/*
	2.2.5 Fence Repair (POJ 3253)
	N:  切成N块
	L:	目标尺寸
*/
func solve2_2_5(N int, L []int) {
	ans := 0
	for 1 < N {
		// 求出最短木板和次短木板
		mii1, mii2 := 0, 1
		if L[mii1] > L[mii2] {
			L[mii1], L[mii2] = L[mii2], L[mii1]
		}

		for i := 2; i < N; i++ {
			if L[i] < L[mii1] {
				mii2 = mii1
				mii1 = i
			} else if L[i] < L[mii2] {
				mii2 = i
			}
		}

		// 合并
		t := L[mii1] + L[mii2]
		ans += t

		if mii1 == N-1 {
			L[mii1], L[mii2] = L[mii2], L[mii1]
		}
		L[mii1] = t
		L[mii2] = L[N-1] // ?
		N--
	}
}
