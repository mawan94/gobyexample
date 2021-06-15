package tree

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"math"
	"testing"
)

type AVLNode struct {
	Left   *AVLNode
	Right  *AVLNode
	Val    int
	height int
}

type AVLTree struct {
	Root *AVLNode
	Size int
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		Root: nil,
		Size: 0,
	}
}

// 获取平衡因子
func (tree *AVLTree) GetBalanceFactor(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return tree.GetHeight(node.Left) - tree.GetHeight(node.Right)
}

// 获取高度
func (tree *AVLTree) GetHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

// 根据值获取节点
func (tree *AVLTree) GetNode(root *AVLNode, val int) *AVLNode {
	if root == nil {
		return nil
	} else if root.Val > val {
		return tree.GetNode(root.Left, val)
	} else if root.Val < val {
		return tree.GetNode(root.Right, val)
	}
	return root
}

// 获取最小节点
func (tree *AVLTree) GetMinNode(root *AVLNode) *AVLNode {
	if root.Left == nil {
		return root
	}
	return tree.GetMinNode(root.Left)
}

// 右旋转
func (tree *AVLTree) rightRotate(node *AVLNode) *AVLNode {
	if node == nil || node.Left == nil {
		return node
	}
	left := node.Left

	node.Left = left.Right
	left.Right = node

	// 更新高度
	node.height = int(math.Max(float64(tree.GetHeight(node.Left)), float64(tree.GetHeight(node.Right)))) + 1
	left.height = int(math.Max(float64(tree.GetHeight(left.Left)), float64(tree.GetHeight(left.Right)))) + 1

	if tree.Root == node {
		tree.Root = left
	}
	return left
}

// 左旋转
func (tree *AVLTree) leftRotate(node *AVLNode) *AVLNode {
	if node == nil || node.Right == nil {
		return node
	}
	right := node.Right

	node.Right = right.Left
	right.Left = node
	// 更新高度
	node.height = int(math.Max(float64(tree.GetHeight(node.Left)), float64(tree.GetHeight(node.Right)))) + 1
	right.height = int(math.Max(float64(tree.GetHeight(right.Left)), float64(tree.GetHeight(right.Right)))) + 1

	if tree.Root == node {
		tree.Root = right
	}
	return right
}

func (tree *AVLTree) Add(Val int) {
	if tree.Root == nil {
		tree.Root = &AVLNode{
			Left:   nil,
			Right:  nil,
			Val:    Val,
			height: 1,
		}
	} else {
		tree.addNode(tree.Root, Val)
	}
	tree.Size++
}

func (tree *AVLTree) addNode(root *AVLNode, val int) *AVLNode {
	if root == nil {
		return &AVLNode{
			Left:   nil,
			Right:  nil,
			Val:    val,
			height: 1,
		}
	} else if root.Val < val {
		root.Right = tree.addNode(root.Right, val)
	} else if root.Val > val {
		root.Left = tree.addNode(root.Left, val)
	}

	// 更新高度
	root.height = int(math.Max(float64(tree.GetHeight(root.Left)), float64(tree.GetHeight(root.Right)))) + 1

	// 获取平衡因子
	balanceFactor := tree.GetBalanceFactor(root)
	if balanceFactor > 1 && tree.GetBalanceFactor(root.Left) > 0 { // LL
		return tree.rightRotate(root)
	} else if balanceFactor > 1 && tree.GetBalanceFactor(root.Left) < 0 { //LR
		root.Left = tree.leftRotate(root.Left)
		return tree.rightRotate(root)
	} else if balanceFactor < -1 && tree.GetBalanceFactor(root.Right) < 0 { // RR
		return tree.leftRotate(root)
	} else if balanceFactor < -1 && tree.GetBalanceFactor(root.Right) > 0 { // RL
		root.Right = tree.rightRotate(root.Right)
		return tree.leftRotate(root)
	}
	return root
}

// 删除
// 叶子节点
// 只有一个孩子
// 有两个孩子
func (tree *AVLTree) Remove(val int) *AVLNode {
	node := tree.GetNode(tree.Root, val)
	if node == nil {
		return nil
	}
	tree.Size--
	return tree.removeNode(tree.Root, val)
}

func (tree *AVLTree) removeNode(root *AVLNode, val int) *AVLNode {
	var ret *AVLNode = nil
	if root.Val > val {
		root.Left = tree.removeNode(root.Left, val)
		ret = root
	} else if root.Val < val {
		root.Right = tree.removeNode(root.Right, val)
		ret = root
	} else {
		if root.Left == nil { // 左子树为空
			if root == tree.Root {
				tree.Root = root.Right
			}
			ret = root.Right
		} else if root.Right == nil { // 右子树为空
			if root == tree.Root {
				tree.Root = root.Left
			}
			ret = root.Left
		} else { // 左右子树均不为空 找到前驱或者后继节点替换
			replace := tree.GetMinNode(root.Right)
			replace.Right = tree.removeNode(root.Right, replace.Val)
			replace.Left = root.Left
			ret = replace
			if root == tree.Root {
				tree.Root = ret
			}
		}
	}
	if ret == nil {
		return ret
	}
	// 更新高度
	ret.height = int(math.Max(float64(tree.GetHeight(ret.Left)), float64(tree.GetHeight(ret.Right)))) + 1
	// ***** 平衡调整 *****
	balanceFactor := tree.GetBalanceFactor(ret)
	if balanceFactor > 1 && tree.GetBalanceFactor(ret.Left) > 0 {
		return tree.rightRotate(ret)
	} else if balanceFactor > 1 && tree.GetBalanceFactor(ret.Left) < 0 {
		ret.Left = tree.leftRotate(ret.Left)
		return tree.rightRotate(ret)
	} else if balanceFactor < -1 && tree.GetBalanceFactor(ret.Right) < 0 {
		return tree.leftRotate(ret)
	} else if balanceFactor < -1 && tree.GetBalanceFactor(ret.Right) > 0 {
		ret.Right = tree.rightRotate(ret.Right)
		return tree.leftRotate(ret)
	}
	return ret
}

func (tree *AVLTree) bfsTest() {
	root := tree.Root
	if root == nil {
		return
	}
	currLine := 1
	currLast := root
	NextLast := root
	q := datastructure.NewLinkedList()
	q.Enqueue(root)

	fmt.Printf("第%d行   ", currLine)
	for !q.IsEmpty() {
		node := q.Dequeue().(*AVLNode)
		fmt.Printf("%d ", node.Val)
		if node.Left != nil {
			q.Enqueue(node.Left)
			NextLast = node.Left
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
			NextLast = node.Right
		}
		if currLast == node && !q.IsEmpty() {
			currLine++
			fmt.Println()
			fmt.Printf("第%d行   ", currLine)
			currLast = NextLast
		}
	}
}

func TestAVLTree(t *testing.T) {
	tree := NewAVLTree()
	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(8)
	tree.bfsTest()

	fmt.Println()
	fmt.Println()
	fmt.Println()

	tree.Remove(8)
	tree.Remove(2)
	tree.Remove(4)
	tree.Remove(5)
	tree.Remove(7)
	tree.Remove(3)
	tree.Remove(1)
	tree.Remove(6)
	tree.bfsTest()

	fmt.Println()
	fmt.Println()
	fmt.Println()
}
