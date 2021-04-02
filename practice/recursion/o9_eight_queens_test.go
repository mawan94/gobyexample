package recursion

import (
	"fmt"
	"math"
	"testing"
)

var total = 0

// 八皇后问题
//在8×8格的国际象棋上摆放8个皇后，使其不能互相攻击，即任意两个皇后都不能处于同一行、同一列或同一斜线上，问有多少种摆法？
func eightQueens(position [8]int, row int) {
	if row == 8 {
		total++
		fmt.Println(position)
		return
	}
	for col := 0; col < 8; col++ {
		if checkFeasibility(position, row, col) {
			position[row] = col
			eightQueens(position, row+1)
		}
	}
}

// 检测摆放于 row col 位置上的可行性
func checkFeasibility(position [8]int, row, col int) bool {
	for i := 0; i < row; i++ {
		if col == position[i] || (float64(row-i) == math.Abs(float64(col-position[i]))) {
			return false
		}
	}
	return true
}

func TestEightQueens(t *testing.T) {
	eightQueens([8]int{}, 0)
	fmt.Println(total)
}

//关于回溯的三种问题，模板略有不同，
//第一种，返回值是true/false。
//第二种，求个数，设全局counter，返回值是void；求所有解信息，设result，返回值void。
//第三种，设个全局变量best，返回值是void。

//第一种：
//
//boolean solve(Node n) {
//	if n is a leaf node {
//		if the leaf is a goal node, return true
//	else return false
//	} else {
//		for each child c of n {
//			if solve(c) succeeds, return true
//		}
//		return false
//	}
//}

//第二种：
//
//void solve(Node n) {
//	if n is a leaf node {
//		if the leaf is a goal node, count++, return;
//		else return
//	} else {
//		for each child c of n {
//			solve(c)
//		}
//	}
//}


//第三种：
//
//void solve(Node n) {
//	if n is a leaf node {
//		if the leaf is a goal node, update best result, return;
//		else return
//	} else {
//		for each child c of n {
//			solve(c)
//		}
//	}
//}

