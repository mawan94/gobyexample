package recursion

import (
	"math"
	"testing"
)

//给你一个整数数组 jobs ，其中 jobs[i] 是完成第 i 项工作要花费的时间。

//请你将这些工作分配给 k 位工人。所有工作都应该分配给工人，且每项工作只能分配给一位工人。
//工人的 工作时间 是完成分配给他们的所有工作花费时间的总和。
//请你设计一套最佳的工作分配方案，使工人的 最大工作时间 得以 最小化 。

//示例 1：
//
//输入：jobs = [3,2,3], k = 3
//输出：3
//解释：给每位工人分配一项工作，最大工作时间是 3 。
//示例 2：
//
//输入：jobs = [1,2,4,7,8], k = 2
//输出：11
//解释：按下述方式分配工作：
//1 号工人：1、2、8（工作时间 = 1 + 2 + 8 = 11）
//2 号工人：4、7（工作时间 = 4 + 7 = 11）
//最大工作时间是 11 。
//
//提示：
//
//1 <= k <= jobs.length <= 12
//1 <= jobs[i] <= 107

var res  = 1 << 31 - 1



func workTime(jobs []int, K int) int {
	personTime:=make([]int,K)
	backtrack(jobs,personTime,0)
	return res
}

func backtrack(jobs, personTime []int, index int) {
	if index == len(jobs) {
		maxTime := -1
		// 找到最长任务消耗的时间
		for i := 0; i < len(personTime); i++ {
			maxTime = int(math.Max(float64(personTime[i]), float64(maxTime)))
		}
		res = int(math.Min(float64(res), float64(maxTime)))
		return
	}

	for i := 0; i < len(personTime); i++ {
		//进行剪枝1 (不是最优解)
		if personTime[i] + jobs[index] >= res {
			continue
		}
		personTime[i] += jobs[index]
		backtrack(jobs,personTime,index+1)
		personTime[i] -= jobs[index]
		//进行剪枝2 因为当工人的工作时间都是0时，那么选择放在哪个工人效果都是一样的，他们的分配没有顺序之分的
		//比如三个工人：[1,2,4,7][8][]，放完8回溯会[1,2,4,7][][8],这样是不是就没意义了啊，所以就不用放了
		if personTime[i] == 0 {
			break
		}
	}
}



func TestWorkTime(t *testing.T) {
	//输入：jobs = [3,2,3], k = 3
	//输出：3
	jobs := []int{3,2,3}
	println(workTime(jobs, 2))
}
