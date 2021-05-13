package tree

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
)

// ================== 简单递归版 ====================
/*
 核心:  一共会路过同一节点3次  要理清楚第几次来到节点时进行打印或处理
*/
func recursionFront(root *datastructure.TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val.(int))
	recursionFront(root.Left)
	recursionFront(root.Right)
}

func recursionMiddle(root *datastructure.TreeNode) {
	if root == nil {
		return
	}
	recursionMiddle(root.Left)
	fmt.Printf("%d ", root.Val.(int))
	recursionMiddle(root.Right)
}

func recursionBack(root *datastructure.TreeNode) {
	if root == nil {
		return
	}
	recursionBack(root.Left)
	recursionBack(root.Right)
	fmt.Printf("%d ", root.Val.(int))
}

// ================== 双色标记版 ====================
/*
 双色迭代法是一种可以用迭代模拟递归的写法，其写法和递归非常相似，要比普通迭代简单地多。
 实现上 WHITE 就表示的是递归中的第一次进入过程，Gray 则表示递归中的从叶子节点返回的过程。
*/
const (
	WHITE = iota // 第一次入栈标记为白色
	GRAY         // 第一次入栈标记为灰色
)

type Mark struct {
	Color int
	Node  *datastructure.TreeNode
}

func twoColorMarkFront(root *datastructure.TreeNode) {
	stack := &datastructure.Stack{}
	stack.Push(Mark{WHITE, root})
	for !stack.IsEmpty() {
		mark := stack.Pop().(Mark)
		if mark.Node == nil {
			continue
		}
		if mark.Color == WHITE {
			stack.Push(Mark{WHITE, mark.Node.Right})
			stack.Push(Mark{WHITE, mark.Node.Left})
			stack.Push(Mark{GRAY, mark.Node})
		} else {
			fmt.Printf("%d ", mark.Node.Val.(int))
		}
	}
}

func twoColorMarkMiddle(root *datastructure.TreeNode) {
	stack := &datastructure.Stack{}
	stack.Push(Mark{WHITE, root})
	for !stack.IsEmpty() {
		mark := stack.Pop().(Mark)
		if mark.Node == nil {
			continue
		}
		if mark.Color == WHITE {
			stack.Push(Mark{WHITE, mark.Node.Right})
			stack.Push(Mark{GRAY, mark.Node})
			stack.Push(Mark{WHITE, mark.Node.Left})
		} else {
			fmt.Printf("%d ", mark.Node.Val.(int))
		}
	}
}

func twoColorMarkBack(root *datastructure.TreeNode) {
	stack := &datastructure.Stack{}
	stack.Push(Mark{WHITE, root})
	for !stack.IsEmpty() {
		mark := stack.Pop().(Mark)
		if mark.Node == nil {
			continue
		}
		if mark.Color == WHITE {
			stack.Push(Mark{GRAY, mark.Node})
			stack.Push(Mark{WHITE, mark.Node.Right})
			stack.Push(Mark{WHITE, mark.Node.Left})
		} else {
			fmt.Printf("%d ", mark.Node.Val.(int))
		}
	}
}

// ======================================
func stackIterFront(root *datastructure.TreeNode) {
	stack := &datastructure.Stack{}
	curr := root
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			fmt.Printf("%d ", curr.Val.(int))
			stack.Push(curr)
			curr = curr.Left
		}
		if !stack.IsEmpty() {
			curr = stack.Pop().(*datastructure.TreeNode).Right
		}
	}
}

func stackIterMiddle(root *datastructure.TreeNode) {
	stack := &datastructure.Stack{}
	curr := root
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		if !stack.IsEmpty() {
			curr = stack.Pop().(*datastructure.TreeNode)
			fmt.Printf("%d ", curr.Val.(int))
			curr = curr.Right
		}
	}
}

func stackIterBack(root *datastructure.TreeNode) {
	stack := &datastructure.Stack{}
	complete := make(map[*datastructure.TreeNode]bool)
	curr := root
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		curr = stack.Pop().(*datastructure.TreeNode)
		if curr.Right != nil && !complete[curr]{
			stack.Push(curr)
			complete[curr] = true
			curr = curr.Right
		} else {
			fmt.Printf("%d ", curr.Val.(int))
			curr = nil
		}
	}
}

// https://blog.csdn.net/danmo_wuhen/article/details/104339630
// =========================================
/*
 建立一种机制，对于没有左子树的节点只到达一次，对于有左子树的节点会到达两次
 假定当前节点为cur 并且开始时赋值为根节点root
	1. 判断cur 是否为nil
	2. 如果不是nil
		1) 如果没有左孩子，pre向右更新 (pre = pre.Right)
		2) 如果有左孩子，则从左子树找到最右侧的节点 记作pre
			1.if pre.Right == nil  pre.Right = cur    cur = cur.Left
		    2.if pre.Right == cur  pre,Right = nil	  cur = cur.Right
	3. if curr == nil return
*/
func morrisFront(root *datastructure.TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil { // 第一次到达左子树的最右端
				fmt.Printf("%d ", cur.Val.(int))
				pre.Right = cur
				cur = cur.Left
			} else { // 第二次到达左子树的最右端
				pre.Right = nil
				cur = cur.Right
			}
		} else {
			fmt.Printf("%d ", cur.Val.(int))
			cur = cur.Right
		}
	}
}

func morrisMiddle(root *datastructure.TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else {
				fmt.Printf("%d ", cur.Val.(int))
				pre.Right = nil
				cur = cur.Right
			}
		} else {
			fmt.Printf("%d ", cur.Val.(int))
			cur = cur.Right
		}
	}
}

func morrisBack(root *datastructure.TreeNode) {
	curr := &datastructure.TreeNode{
		Left: root,
	}
	for curr != nil {
		if curr.Left != nil {
			mostRight := curr.Left
			for mostRight.Right != nil && mostRight.Right != curr {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil { // 第一次
				mostRight.Right = curr
				curr = curr.Left
			} else { // 第二次
				mostRight.Right = nil
				PrintNode(curr.Left)
				curr = curr.Right
			}
		} else {
			curr = curr.Right
		}
	}
}

//func reverseNode(pre,cur *datastructure.TreeNode) *datastructure.TreeNode {
//	if cur == nil   {
//		return pre
//	}
//	waitProcess := cur.Right
//	cur.Right = pre
//	return reverseNode(cur,waitProcess)
//}

func PrintNode(root *datastructure.TreeNode) {
	stack := &datastructure.Stack{}
	tmp := root
	for tmp != nil {
		stack.Push(tmp)
		tmp = tmp.Right
	}
	for !stack.IsEmpty() {
		fmt.Printf("%d ", stack.Pop().(*datastructure.TreeNode).Val)
	}
}

// ********** BFS **********
func levelOrder(root *datastructure.TreeNode) {
	currLast, nextLast := root, root // 当前层的最后一个元素 ， 下一层的最后一个元素
	currLevel := 1                   // 当前层数

	queue := datastructure.Queue{}
	queue.Add(root)
	fmt.Printf("第%d层   ", currLevel)

	for !queue.IsEmpty() {
		curr := queue.Poll().(*datastructure.TreeNode)
		if curr.Left != nil {
			nextLast = curr.Left
			queue.Add(curr.Left)
		}
		if curr.Right != nil {
			nextLast = curr.Right
			queue.Add(curr.Right)
		}

		fmt.Printf("%d ", curr.Val.(int))
		if curr == currLast && !queue.IsEmpty() {
			currLast = nextLast
			currLevel++
			fmt.Println()
			fmt.Printf("第%d层   ", currLevel)
		}
	}
}

func TestIterationTree(t *testing.T) {
	root := &datastructure.TreeNode{
		Val: 1,
		Left: &datastructure.TreeNode{
			Val: 2,
			Left: &datastructure.TreeNode{
				Val: 5,
			},
			Right: &datastructure.TreeNode{
				Val: 4,
			},
		},
		Right: &datastructure.TreeNode{
			Val: 3,
			Right: &datastructure.TreeNode{
				Val: 8,
			},
		},
	}

	fmt.Println("=======递归版=======")
	recursionFront(root)
	fmt.Println()
	recursionMiddle(root)
	fmt.Println()
	recursionBack(root)
	fmt.Println()

	fmt.Println("=======双色标记=======")
	twoColorMarkFront(root)
	fmt.Println()
	twoColorMarkMiddle(root)
	fmt.Println()
	twoColorMarkBack(root)
	fmt.Println()

	fmt.Println("=======栈迭代=======")
	stackIterFront(root)
	fmt.Println()
	stackIterMiddle(root)
	fmt.Println()
	stackIterBack(root)
	fmt.Println()

	fmt.Println("=======Morris=======")
	morrisFront(root)
	fmt.Println()
	morrisMiddle(root)
	fmt.Println()
	morrisBack(root)
	fmt.Println()

	fmt.Println("=======层序遍历=======")
	levelOrder(root)
	fmt.Println()
}
