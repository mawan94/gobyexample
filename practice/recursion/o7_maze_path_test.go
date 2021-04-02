package recursion

import (
	"fmt"
	"testing"
)

// 走迷宫
// 有一个 n * m 的迷宫 （迷宫中存在障碍）
// 每次只能上下左右移动一格
// 从x点出发 到达y点 是否能走成功？ 画出行进路线。

// 迷宫模型
//	起点	2	0	0
//	0	0	2	0
//	0	0	0	0
//	0	0	2	0
//	0	0	2	终点

// 0 代表未尝试走过的格子
// 1 代表走过的格子
// 2 代表墙
// 3 代表走过并且不通的格子

const (
	INIT = iota
	SUCCESS
	WALL
	FAIL
)

func mazePath(maze [][]int, x, y int) bool {
	// 达到目标点
	if maze[4][3] == SUCCESS {
		return true
		// 不合法的移动
	} else if x < 0 || y < 0 || x > 4 || y > 3 {
		return false
		// 可以移动
	} else if maze[x][y] == INIT {
		maze[x][y] = SUCCESS // 移动到本格
		// 进行探路
		if mazePath(maze, x-1, y) { // 上
			return true
		} else if mazePath(maze, x, y+1) { // 右
			return true
		} else if mazePath(maze, x+1, y) { // 下
			return true
		} else if mazePath(maze, x, y-1) { // 左
			return true
		} else {
			maze[x][y] = FAIL // 本格"上下左右"都不通
			return false
		}
		// 不要重复移动
	} else {
		return false
	}
}

func TestMazePath(t *testing.T) {
	maze := make([][]int, 5)
	for i := 0; i < len(maze); i++ {
		maze[i] = make([]int, 4)
	}
	maze[0][1] = WALL
	maze[1][2] = WALL
	maze[3][2] = WALL
	maze[4][2] = WALL

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			fmt.Printf("%d  ", maze[i][j])
		}
		fmt.Println()
	}
	mazePath(maze, 0, 0)
	fmt.Println()
	fmt.Println("==========")
	fmt.Println()
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			fmt.Printf("%d  ", maze[i][j])
		}
		fmt.Println()
	}
}
