package recursion
//
//import (
//	"fmt"
//	"math"
//)
//
///*
//https://leetcode-cn.com/circle/article/koSrVI/
//1 把问题规模缩小，找到问题最后一步的情况
//2 通过步骤1 找到最优子结构
//3 通过子结构定义转移方程式
//4 找到初始情况和边界条件
//*/
//var (
//	array [8]int
//	total int
//	goods = []Goods{
//		{10, 1},
//		{35, 3},
//		{22, 2},
//		{55, 5},
//		{48, 4},
//	}
//)
//
//func main() {
//	// 1 爬楼梯
//	fmt.Println(climbTheStairs(3))
//
//	// 2 翻转字符串
//	b := []byte("hello")
//	reverseStr(b)
//	fmt.Println(string(b))
//
//	// 3 链表翻转
//	head := &Node{
//		1,
//		&Node{
//			2,
//			&Node{
//				3,
//				nil,
//			},
//		},
//	}
//	reverseHead := reverseNode(head)
//	for reverseHead != nil {
//		fmt.Printf("%v -> ", reverseHead.Val)
//		reverseHead = reverseHead.Next
//	}
//	fmt.Println()
//
//	// 4 Node SWAP
//	node := &Node{
//		1,
//		&Node{
//			3,
//			&Node{
//				2,
//				nil,
//			},
//		},
//	}
//	result := swapPairs(node)
//	for result != nil {
//		fmt.Printf("%v -> ", result.Val)
//		result = result.Next
//	}
//
//	fmt.Println()
//
//	// 5 合并硬币
//	fmt.Println(minMergeCoinsCount(27, []int{2, 5, 7}))
//
//	// 6 走格子
//	fmt.Println(maxWaysCount(3, 3))
//
//	// 7 迷宫路径
//	maze := make([][]int, 7)
//
//	for i := 0; i < 7; i++ {
//		maze[i] = make([]int, 7)
//	}
//	for i := 0; i < 7; i++ {
//		maze[0][i] = 1
//		maze[6][i] = 1
//	}
//	for i := 0; i < 7; i++ {
//		maze[i][0] = 1
//		maze[i][6] = 1
//	}
//	maze[1][2] = 1
//	maze[1][3] = 1
//	maze[4][5] = 1
//
//	for i := 0; i < len(maze); i++ {
//		fmt.Println(maze[i])
//	}
//	fmt.Println("================")
//	mazePath(maze, 5, 5, 1, 1)
//	for i := 0; i < len(maze); i++ {
//		fmt.Println(maze[i])
//	}
//
//	// 8 汉诺塔
//	towerOfHanoi(3, "a", "b", "c")
//
//	// 9 八皇后
//	eightQueens(0)
//	fmt.Println("total: ", total)
//
//	// 10 ola
//	fmt.Println("ola: ", ola(4))
//
//	// 11 背包问题
//	fmt.Println(MaxPrice(6, 0))
//
//	// 12 二叉树生成
//	res := generateTrees(3)
//	for _, item := range res {
//		printTree(item)
//		fmt.Println()
//	}
//
//	//13 字符串的所有子序列
//	subsequence([]byte("abc"), 0, "")
//}
//
//// 1 爬楼
//// N层楼梯(N>=0) 1次只能爬1层或者2层  一共有几种方法？
//func climbTheStairs(N int) int {
//	if N <= 2 {
//		return N
//	}
//	return climbTheStairs(N-1) + climbTheStairs(N-2)
//}
//
//// 2 翻转字符串 (不开辟额外空间)
//func reverseStr(str []byte) {
//	if str == nil || len(str) <= 1 {
//		return
//	}
//	str[0], str[len(str)-1] = str[len(str)-1], str[0]
//	reverseStr(str[1 : len(str)-1])
//
//}
//
//// 3 翻转链表
//type Node struct {
//	Val  int
//	Next *Node
//}
//
//func reverseNode(head *Node) *Node {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	ret := reverseNode(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return ret
//}
//
//// 4 两两交换链表中的节点
//// 两两交换其中相邻的节点，并返回交换后的链表。
//// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换
//func swapPairs(node *Node) *Node {
//	if node == nil || node.Next == nil {
//		return node
//	}
//	nextNode := node.Next //没交换前的NEXT
//	ret := swapPairs(node.Next.Next)
//	nextNode.Next = node //没交换前的NEXT.NEXT
//	node.Next = ret
//	return nextNode
//}
//
//// 5 拼凑硬币
//// target 27
//// coins: 2  5  7
//// f(n) = min{ f(n - 2) + 1, f(n - 5) + 1,f(n - 7) + 1}
//func minMergeCoinsCount(target int, coins []int) int {
//	if target == 0 {
//		return 0
//	}
//	min := -1
//	for i := 0; i < len(coins); i++ {
//		if target >= coins[i] {
//			ret := minMergeCoinsCount(target-coins[i], coins)
//			if ret == -1 {
//				continue
//			} else if min == -1 {
//				min = ret + 1
//			} else if min > ret+1 {
//				min = ret + 1
//			}
//		}
//	}
//	return min
//}
//
//// 6 走格子  （一共有多少种走法） PS 只能向下或者向右
////BEGIN		 0		 0
////	0		 0		x,y-1
////	0		x-1,y	 END
//
//func maxWaysCount(x, y int) int {
//	if x == 1 || y == 1 {
//		return 1
//	}
//	return maxWaysCount(x-1, y) + maxWaysCount(x, y-1)
//
//}
//
//// 7 走格子（迷宫路径）
//// 1：墙    0：可走  B: 起点   E：终点
//// 走过的用2标识 走失败的用3标识
//const (
//	INIT = iota
//	WALL
//	SUCCESS
//	FAIL
//)
//
//func mazePath(maze [][]int, end_x, end_y int, my_x, my_y int) bool {
//	if maze[end_x][end_y] == SUCCESS {
//		return true
//	} else {
//		// 没走过
//		if maze[my_x][my_y] == INIT {
//			maze[my_x][my_y] = SUCCESS
//			// 策略： 上右下左
//			if mazePath(maze, end_x, end_y, my_x-1, my_y) {
//				return true
//			} else if mazePath(maze, end_x, end_y, my_x, my_y+1) {
//				return true
//			} else if mazePath(maze, end_x, end_y, my_x+1, my_y) {
//				return true
//			} else if mazePath(maze, end_x, end_y, my_x, my_y-1) {
//				return true
//			} else {
//				maze[my_x][my_y] = FAIL
//				return false
//			}
//		} else {
//			return false
//		}
//	}
//}
//
//// 8 汉诺塔   把A柱子上的圆盘全部移动到C
//func towerOfHanoi(N int, a, b, c string) {
//	if N == 1 {
//		fmt.Printf("%s -> %s \n", a, c)
//		return
//	}
//	towerOfHanoi(N-1, a, c, b)
//	fmt.Printf("%s -> %s \n", a, c)
//	towerOfHanoi(N-1, b, a, c)
//}
//
//// 9 八皇后
//// 国际象棋 8 * 8 保证每个皇后不能相互攻击到
//func eightQueens(N int) {
//	if N == 8 {
//		//fmt.Println(array)
//		total++
//		return
//	}
//	for i := 0; i < 8; i++ {
//		if check(i, N) {
//			array[N] = i
//			eightQueens(N + 1)
//		}
//	}
//}
//
//// 不能再同一列  不能在同一斜线
//func check(col, N int) bool {
//	for i := 0; i < N; i++ {
//		if array[i] == col || (math.Abs(float64(col-array[i])) == math.Abs(float64(N-i))) {
//			return false
//		}
//	}
//	return true
//}
//
//// 10 欧拉投信
//// 一个村每个人只能寄一份信，也只能收一份信 但是不能寄送给自己。若村里有N人 一共有多少方式？
//
//// 2     1		A -> B   B -> A
//// 3     2		A -> B  B -> C  C -> A    || A -> C  C -> B  B -> A
//// 4     ?      A B C D
//// 假设 A寄给B  B寄给A 就是两个人的问题  f(n - 2)
//// 假设 A寄给B  B没有寄给A  那么A需要有人给他回一封  B要给别人寄一封 两人的状态相互抵消 看做 f(n - 1) 回到三个人的游戏
//// A可以寄给B 也可以给C , D   f(n) = (n - 1) （f(n - 1) + f(n - 2)）
//func ola(N int) int {
//	if N == 1 {
//		return 0
//	} else if N == 2 {
//		return 1
//	} else if N == 3 {
//		return 2
//	} else {
//		return (N - 1) * (ola(N-2) + ola(N-1))
//	}
//}
//
//// 11 背包问题
//type Goods struct {
//	Price  int
//	Weight int
//}
//
//func MaxPrice(packageMaxWeight, curIndex int) int {
//	if curIndex == len(goods)-1 && packageMaxWeight < goods[curIndex].Weight {
//		// 最后一个重量不够
//		return 0
//	} else if curIndex == len(goods)-1 && packageMaxWeight >= goods[curIndex].Weight {
//		// 最后一个重量够 必然拿
//		return goods[curIndex].Price
//	} else {
//		// 拿
//		case1 := 0
//		if packageMaxWeight >= goods[curIndex].Weight {
//			case1 = MaxPrice(packageMaxWeight-goods[curIndex].Weight, curIndex+1) + goods[curIndex].Price
//		}
//		// 不拿
//		case2 := MaxPrice(packageMaxWeight, curIndex+1)
//		return int(math.Max(float64(case1), float64(case2)))
//	}
//}
//
//// 不同的二叉搜索树
////12 给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
//// 1, 2, 3, 4.... n
//// 当 看做是ROOT节点时候 其左节点为  【1-2】的所有排列   右节点为【4-N】的所有排列
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func generate(start, end int) []*TreeNode {
//	res := make([]*TreeNode, 0)
//	if start > end {
//		res = append(res, nil)
//		return res
//	}
//
//	for i := start; i <= end; i++ {
//		l := generate(start, i-1)
//		r := generate(i+1, end)
//
//		for _, ln := range l {
//			for _, rn := range r {
//				root := &TreeNode{
//					i,
//					ln,
//					rn,
//				}
//				res = append(res, root)
//			}
//		}
//	}
//	return res
//}
//
//func printTree(treeNode *TreeNode) {
//	if treeNode != nil {
//		fmt.Print(treeNode.Val, " ")
//		printTree(treeNode.Left)
//		printTree(treeNode.Right)
//	} else {
//		fmt.Print("nil ")
//	}
//}
//func generateTrees(n int) []*TreeNode {
//	if n == 0 {
//		return make([]*TreeNode, 0)
//	}
//	return generate(1, n)
//}
//
//// 13 字符串的所有子序列
//func subsequence(str []byte, idx int, path string) {
//	if idx >= len(str) {
//		fmt.Println(path)
//		return
//	} else {
//		// y
//		subsequence(str, idx+1, path+string(str[idx]))
//		// N
//		subsequence(str, idx+1, path)
//	}
//}
//
//// 14 字符串所有的组合
//func totalAssemble(str string,idx int)  {
//	if idx < len(str) {
//
//	}
//
//}
//
//
//
//
//// 15 Input: 5
////Output: True
////Explanation: 1 * 1 + 2 * 2 = 5
////题目描述：判断一个非负整数是否为两个整数的平方和。
////func f(input int) bool {
////	for i := 1; i < input; i++ {
////		for j := i + 1; j <= input; j++ {
////			if i*i+j*j == input {
////				return true
////			}
////		}
////	}
////	return false
////}
