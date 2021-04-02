package recursion

import (
	"fmt"
	"math"
	"testing"
)

// 小偷问题
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
//如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

func smartThief(room []int, cur int) int {
	if cur < 0 {
		return 0
	}
	if cur == 0 {
		return room[cur]
	}
	if cur == 1 {
		return int(math.Max(float64(room[cur]), float64(room[cur-1])))
	} else {
		way1 := room[cur] + smartThief(room, cur-2)
		way2 := smartThief(room, cur-1)
		return int(math.Max(float64(way2), float64(way1)))
	}

}

func TestSmartThief(t *testing.T) {
	room := []int{5, 9, 4, 1, 2}
	fmt.Println(smartThief(room, len(room)-1))
}
