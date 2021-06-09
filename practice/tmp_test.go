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
			curr = node.Right
		} else if node.Right == nil {
			curr = node.Left
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

// =========================================================================
// =========================================================================
// =========================================================================
// =========================================================================

const (
	BLACK = iota
	RED
)

type RBNode struct {
	Left, Right, Parent *RBNode
	Data, Color         int
}
type RBTree struct {
	LEAF, Root *RBNode
	Size       int
}

func NewRBTree() *RBTree {
	blankNode := &RBNode{Color: BLACK}
	return &RBTree{blankNode, blankNode, 0}
}

func (tree *RBTree) getBrother(node *RBNode) *RBNode {
	if node == nil || node.Parent == nil {
		return nil
	} else if node.Parent.Left == node {
		return node.Parent.Right
	} else {
		return node.Parent.Left
	}
}

func (tree *RBTree) getGrandParent(node *RBNode) *RBNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	return node.Parent.Parent
}

func (tree *RBTree) getUncle(node *RBNode) *RBNode {
	gp := tree.getGrandParent(node)
	if gp != nil {
		if node.Parent == gp.Left {
			return gp.Right
		} else {
			return gp.Left
		}
	}
	return nil
}

func (tree *RBTree) getMinNode(node *RBNode) *RBNode {
	if node == tree.LEAF || node.Left == tree.LEAF {
		return node
	}
	return tree.getMinNode(node.Left)
}

func (tree *RBTree) rightRotation(node *RBNode) {
	if node == tree.LEAF || node.Left == tree.LEAF {
		return
	}
	left := node.Left
	node.Left = left.Right
	left.Right = node

	if node.Left != tree.LEAF {
		node.Left.Parent = node
	}

	if node.Parent == nil {
		tree.Root = left
	} else if node.Parent.Left == node {
		node.Parent.Left = left
	} else {
		node.Parent.Right = left
	}
	left.Parent = node.Parent
	node.Parent = left
}

func (tree *RBTree) leftRotation(node *RBNode) {
	if node == tree.LEAF || node.Right == tree.LEAF {
		return
	}
	right := node.Right
	node.Right = right.Left
	right.Left = node

	if node.Right != tree.LEAF {
		node.Right.Parent = node
	}

	if node.Parent == nil {
		tree.Root = right
	} else if node.Parent.Left == node {
		node.Parent.Left = right
	} else {
		node.Parent.Right = right
	}
	right.Parent = node.Parent
	node.Parent = right
}

func (tree *RBTree) add(node *RBNode, data int) (root, added *RBNode) {
	if node == tree.LEAF {
		added := &RBNode{
			Left:  tree.LEAF,
			Right: tree.LEAF,
			Data:  data,
			Color: RED,
		}
		if tree.Root == node {
			tree.Root = added
		}
		return added, added
	}
	var addedNode *RBNode = nil
	if node.Data > data {
		node.Left, addedNode = tree.add(node.Left, data)
		node.Left.Parent = node
	} else if node.Data < data {
		node.Right, addedNode = tree.add(node.Right, data)
		node.Right.Parent = node
	}
	return node, addedNode
}

func (tree *RBTree) fixAfterInsertion(curr *RBNode) {
	if curr == tree.Root {
		curr.Color = BLACK
	} else if curr.Parent.Color == RED {
		uncle := tree.getUncle(curr)
		grandParent := tree.getGrandParent(curr)
		if uncle.Color == RED {
			curr.Parent.Color = BLACK
			uncle.Color = BLACK
			grandParent.Color = RED
			tree.fixAfterInsertion(grandParent)
		} else {
			leftLine := (curr == curr.Parent.Left) && (curr.Parent == grandParent.Left)
			rightLine := (curr == curr.Parent.Right) && (curr.Parent == grandParent.Right)
			if leftLine || rightLine {
				curr.Parent.Color = BLACK
				grandParent.Color = RED
				if leftLine {
					tree.rightRotation(grandParent)
				} else {
					tree.leftRotation(grandParent)
				}
			} else {
				if curr == curr.Parent.Left {
					tree.rightRotation(curr.Parent)
					tree.fixAfterInsertion(curr.Right)
				} else {
					tree.leftRotation(curr.Parent)
					tree.fixAfterInsertion(curr.Left)
				}
			}
		}
	}

}

func (tree *RBTree) Add(data int) bool {
	_, added := tree.add(tree.Root, data)
	if added != nil {
		tree.Size++
		tree.fixAfterInsertion(added)
		return true
	}
	return false
}

func (tree *RBTree) delete(node *RBNode, data int) (root, removed, replacement *RBNode) {
	var rm *RBNode = nil
	var rep *RBNode = nil
	if node == nil {
		return nil, nil, nil
	}
	if node.Data > data {
		node.Left, rm, rep = tree.delete(node.Left, data)
	} else if node.Data < data {
		node.Right, rm, rep = tree.delete(node.Right, data)
	} else if node.Data == data {
		if node.Left == tree.LEAF {
			if node == tree.Root {
				tree.Root = node.Right
				tree.Root.Parent = nil
			} else {
				node.Right.Parent = node.Parent
			}
			return node.Right, node, node.Right

		} else if node.Right == tree.LEAF {
			if node == tree.Root {
				tree.Root = node.Left
				tree.Root.Parent = nil
			} else {
				node.Left.Parent = node.Parent
			}
			return node.Left, node, node.Left

		} else {
			replace := tree.getMinNode(node)
			replace.Left = node.Left
			replace.Right, rm, rep = tree.delete(node.Right, data)
			replace.Left.Parent = replace
			replace.Right.Parent = replace

			if node == tree.Root {
				tree.Root = replace
			}
			return replace, node, replace
		}
	}
	return node, rm, rep
}

func (tree *RBTree) fixAfterDelete(curr *RBNode) {
	for curr != tree.Root && curr.Color == BLACK {
		brother := tree.getBrother(curr)
		if curr == curr.Parent.Left {
			if brother.Color == RED {
				brother.Color = BLACK
				curr.Parent.Color = RED
				tree.leftRotation(curr.Parent)
				brother = curr.Parent.Right
			}
			if brother.Left.Color == BLACK && brother.Right.Color == BLACK {
				brother.Color = RED
				curr = curr.Parent
			} else {
				if brother.Right.Color == BLACK {
					brother.Left.Color = BLACK
					brother.Color = RED
					tree.rightRotation(brother)
					brother = curr.Parent.Right
				}
				brother.Color = curr.Parent.Color
				curr.Parent.Color = BLACK
				brother.Right.Color = BLACK
				tree.leftRotation(curr.Parent)
				curr = tree.Root
			}
		} else {
			if brother.Color == RED {
				brother.Color = BLACK
				curr.Parent.Color = RED
				tree.rightRotation(curr.Parent)
				brother = curr.Parent.Left
			}
			if brother.Right.Color == BLACK && brother.Left.Color == BLACK {
				brother.Color = RED
				curr = curr.Parent
			} else {
				if brother.Left.Color == BLACK {
					brother.Right.Color = BLACK
					brother.Color = RED
					tree.leftRotation(brother)
					brother = curr.Parent.Left
				}
				brother.Color = curr.Parent.Color
				curr.Parent.Color = BLACK
				brother.Left.Color = BLACK
				tree.rightRotation(curr.Parent)
				curr = tree.Root
			}
		}
	}
	curr.Color = BLACK
}

func (tree *RBTree) Delete(data int) bool {
	_, removed, replacement := tree.delete(tree.Root, data)
	if removed != nil {
		tree.Size--
		if replacement == tree.LEAF {
			tree.fixAfterDelete(removed)
		} else {
			tree.fixAfterDelete(replacement)
		}
		return true
	}
	return false
}

func (tree *RBTree) bfsRBtree() {
	tmp := tree.Root
	if tmp == nil || tmp == tree.LEAF {
		return
	}

	currentLast := tmp
	nextLast := tmp
	current := 1

	q := &datastructure.Queue{}
	q.Add(tmp)
	fmt.Printf("第%d行 -->  ", current)
	for !q.IsEmpty() {
		node := q.Poll().(*RBNode)
		if node.Left != nil && node.Left != tree.LEAF {
			q.Add(node.Left)
			nextLast = node.Left
		}
		if node.Right != nil && node.Right != tree.LEAF {
			q.Add(node.Right)
			nextLast = node.Right
		}

		var color = ""
		if node.Color == BLACK {
			color = "B"
		} else {
			color = "R"
		}
		fmt.Printf("\t%d(%s) ", node.Data, color)

		if node.Left != nil && node.Left != tree.LEAF {
			fmt.Printf("/")
		}
		if node.Right != nil && node.Right != tree.LEAF {
			fmt.Printf("\\")
		}

		if node == currentLast && !q.IsEmpty() {
			current++
			currentLast = nextLast
			fmt.Println()
			fmt.Printf("第%d行 -->  ", current)
		}
	}
}

func TestFunc(t *testing.T) {

	brt := NewRBTree()
	//brt.Add(1)
	brt.Add(4)
	brt.Add(23)
	brt.Add(5)
	brt.Add(6)
	//brt.Add(13)
	brt.Add(24)
	//brt.Add(43)
	brt.Add(3)
	brt.Add(123)
	//brt.Add(123)
	brt.bfsRBtree()

	fmt.Println()
	fmt.Println()

	brt.Delete(24)
	brt.Delete(123)

	brt.bfsRBtree()

	fmt.Println()
	fmt.Println()

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

	divingBoard(1, 1, 3)

}

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return make([]int, 0)
	}

	ret := make([]int, 0)
	if k == 1 {
		// level1
		ret = append(ret, shorter)
		if shorter != longer {
			ret = append(ret, longer)
		}
		return ret
	}

	// level2
	ret = append(ret, 2*shorter)
	if shorter != longer {
		ret = append(ret, shorter+longer)
		ret = append(ret, 2*longer)
	}

	for currentLevel := 3; currentLevel <= k; currentLevel++ {
		for i := 0; i < len(ret); i++ {
			ret[i] = ret[i] + shorter
		}

		if shorter != longer {
			ret = append(ret, currentLevel*longer)
		}

	}
	return ret
}
