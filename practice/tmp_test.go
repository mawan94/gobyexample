package practice

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"math"
	"testing"
)

func method1(root *datastructure.TreeNode) {
	if root == nil {
		return
	}
	curr := root
	stack := &datastructure.Stack{}
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			fmt.Printf("%d ", curr.Val)
			stack.Push(curr)
			curr = curr.Left
		}
		curr = stack.Pop().(*datastructure.TreeNode)
		curr = curr.Right
	}
}

func method2(root *datastructure.TreeNode) {
	if root == nil {
		return
	}
	curr := root
	stack := &datastructure.Stack{}
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		curr = stack.Pop().(*datastructure.TreeNode)
		fmt.Printf("%d ", curr.Val)
		curr = curr.Right
	}
}

func method3(root *datastructure.TreeNode) {
	if root == nil {
		return
	}
	curr := root
	var prev *datastructure.TreeNode = nil
	stack := &datastructure.Stack{}
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		curr = stack.Pop().(*datastructure.TreeNode)
		if curr.Right != nil && curr.Right != prev {
			stack.Push(curr)
			curr = curr.Right
			prev = curr
		} else {
			fmt.Printf("%d ", curr.Val)
			curr = nil
		}
	}
}

func method5(root *datastructure.TreeNode) {
	if root == nil {
		return
	}
	curr := root
	for curr != nil {
		if curr.Left != nil {
			mostRight := curr.Left
			for mostRight.Right != nil && mostRight.Right != curr {
				mostRight = mostRight.Right
			}
			// first in
			if mostRight.Right == nil {
				fmt.Printf("%d ", curr.Val)
				mostRight.Right = curr
				curr = curr.Left
			} else { // second in
				mostRight.Right = nil
				curr = curr.Right
			}
		} else {
			fmt.Printf("%d ", curr.Val)
			curr = curr.Right
		}
	}
}

func method6(root *datastructure.TreeNode) {
	if root == nil {
		return
	}
	curr := root
	for curr != nil {
		if curr.Left != nil {
			mostRight := curr.Left
			for mostRight.Right != nil && mostRight.Right != curr {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = curr
				curr = curr.Left
			} else {
				fmt.Printf("%d ", curr.Val)
				mostRight.Right = nil
				curr = curr.Right
			}
		} else {
			fmt.Printf("%d ", curr.Val)
			curr = curr.Right
		}

	}
}

func method7(root *datastructure.TreeNode) {
	if root == nil {
		return
	}

	curr := &datastructure.TreeNode{
		Left: root,
	}
	for curr != nil {
		if curr.Left != nil {
			mostRight := curr.Left
			for mostRight.Right != nil && mostRight.Right != curr {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = curr
				curr = curr.Left
			} else {
				mostRight.Right = nil
				printNode(curr.Left)
				curr = curr.Right
			}

		} else {
			curr = curr.Right
		}
	}
}

func printNode(root *datastructure.TreeNode) {
	ret := reverseNode(nil, root)
	for ret != nil {
		fmt.Printf("%d ", ret.Val)
		ret = ret.Right
	}
	reverseNode(nil, root)
}

func reverseNode(pre, curr *datastructure.TreeNode) *datastructure.TreeNode {
	if curr == nil {
		return pre
	}
	waitProcess := curr.Right
	curr.Right = pre
	return reverseNode(curr, waitProcess)
}

// =========================================================================
// =========================================================================
// =========================================================================
// =========================================================================

type AvlNode struct {
	Left, Right *AvlNode
	Val         int
	height      int
}

type AvlTree struct {
	Root *AvlNode
	Size int
}

func NewAvlTree() *AvlTree {
	return &AvlTree{nil, 0}
}

func GetHeight(node *AvlNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func MinNode(node *AvlNode) *AvlNode {
	if node.Left != nil {
		return MinNode(node.Left)
	}
	return node
}

func GetBalanceFactor(node *AvlNode) int {
	if node == nil {
		return 0
	}
	return GetHeight(node.Left) - GetHeight(node.Right)
}

func (tree *AvlTree) LeftRotate(node *AvlNode) *AvlNode {
	if node == nil || node.Right == nil {
		return node
	}
	right := node.Right
	node.Right = right.Left
	right.Left = node
	// 更新高度
	right.height = 1 + int(math.Max(float64(GetHeight(right.Left)), float64(GetHeight(right.Right))))
	node.height = 1 + int(math.Max(float64(GetHeight(node.Left)), float64(GetHeight(node.Right))))
	if node == tree.Root {
		tree.Root = right
	}
	return right
}

func (tree *AvlTree) RightRotate(node *AvlNode) *AvlNode {
	if node == nil || node.Left == nil {
		return node
	}
	left := node.Left
	node.Left = left.Right
	left.Right = node
	// 更新高度
	left.height = 1 + int(math.Max(float64(GetHeight(left.Left)), float64(GetHeight(left.Right))))
	node.height = 1 + int(math.Max(float64(GetHeight(node.Left)), float64(GetHeight(node.Right))))
	if node == tree.Root {
		tree.Root = left
	}
	return left
}

func (tree *AvlTree) Add(val int) {
	if tree.Root == nil {
		tree.Root = &AvlNode{height: 1, Val: val}
	} else {
		tree.add(tree.Root, val)
	}
	tree.Size++
}

func (tree *AvlTree) add(node *AvlNode, val int) *AvlNode {
	if node == nil {
		return &AvlNode{height: 1, Val: val}
	}
	if node.Val < val {
		node.Right = tree.add(node.Right, val)
	} else if node.Val > val {
		node.Left = tree.add(node.Left, val)
	} else {
		panic("不能插入重复值元素！")
	}
	// 更新高度
	node.height = 1 + int(math.Max(float64(GetHeight(node.Left)), float64(GetHeight(node.Right))))

	// 获取平衡因子
	balanceFactor := GetBalanceFactor(node)
	// 判断是否需要调整?
	// -------- 左高 --------
	if balanceFactor > 1 && GetBalanceFactor(node.Left) > 0 {
		return tree.RightRotate(node)
	} else if balanceFactor > 1 && GetBalanceFactor(node.Left) < 0 {
		node.Left = tree.LeftRotate(node.Left)
		return tree.RightRotate(node)
		// -------- 右高 --------
	} else if balanceFactor < -1 && GetBalanceFactor(node.Right) < 0 {
		return tree.LeftRotate(node)
	} else if balanceFactor < -1 && GetBalanceFactor(node.Right) > 0 {
		node.Right = tree.RightRotate(node.Right)
		return tree.LeftRotate(node)
	}

	return node
}

func (tree *AvlTree) Remove(val int) bool {
	remove := tree.remove(tree.Root, val)
	if remove == nil {
		return false
	}
	tree.Size--
	return true
}

func (tree *AvlTree) remove(node *AvlNode, val int) *AvlNode {
	if node == nil {
		return nil
	}
	var curr *AvlNode = nil

	if node.Val < val {
		node.Right = tree.remove(node.Right, val)
		curr = node
	} else if node.Val > val {
		node.Left = tree.remove(node.Left, val)
		curr = node
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			replace := MinNode(node.Right)
			replace.Left = node.Left
			replace.Right = tree.remove(node.Right, replace.Val)

			replace.height = 1 + int(math.Max(float64(GetHeight(replace.Left)), float64(GetHeight(replace.Right))))
			curr = replace
		}
	}

	balanceFactor := GetBalanceFactor(curr)
	if balanceFactor > 1 && GetBalanceFactor(curr.Left) > 0 {
		return tree.RightRotate(curr)
	} else if balanceFactor > 1 && GetBalanceFactor(curr.Left) < 0 {
		curr.Left = tree.LeftRotate(curr.Left)
		return tree.RightRotate(curr)
	} else if balanceFactor < -1 && GetBalanceFactor(curr.Right) < 0 {
		return tree.LeftRotate(curr)
	} else if balanceFactor < -1 && GetBalanceFactor(curr.Right) > 0 {
		curr.Right = tree.RightRotate(curr.Right)
		return tree.LeftRotate(curr)
	}
	return curr
}

func TestFunc(t *testing.T) {

	tree := NewAvlTree()
	tree.Add(1)
	tree.Add(3)
	tree.Add(5)
	tree.Add(6)
	tree.Add(2)
	tree.Add(4)

	tree.Remove(1)
	tree.Remove(4)
	tree.Remove(5)
	tree.Remove(2)

	root := &datastructure.TreeNode{
		Val: 1,
		Left: &datastructure.TreeNode{
			Val: 2,
		},
		Right: &datastructure.TreeNode{
			Val: 3,
		},
	}

	method1(root)
	fmt.Println()
	method2(root)
	fmt.Println()
	method3(root)
	fmt.Println()
	method5(root)
	fmt.Println()
	method6(root)
	fmt.Println()
	method7(root)
	fmt.Println()
	//fun7(root)
	//fmt.Println()
	//level(root)
	fmt.Println()

}
