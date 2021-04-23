package practice

import (
	"fmt"
	"gobyexample/practice/datastructure"
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
			}else {
				mostRight.Right = nil
				printNode(curr.Left)
				curr = curr.Right
			}

		}else {
			curr = curr.Right
		}
	}
}

func printNode(root *datastructure.TreeNode)  {
	ret := reverseNode(nil,root)
	for ret != nil {
		fmt.Printf("%d ",ret.Val)
		ret = ret.Right
	}
	reverseNode(nil,root)
}

func reverseNode(pre,curr *datastructure.TreeNode)*datastructure.TreeNode {
	if curr == nil {
		return pre
	}
	waitProcess := curr.Right
	curr.Right = pre
	return reverseNode(curr,waitProcess)
}


func TestFunc(t *testing.T) {
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
