package tree

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
)

// https://juejin.cn/post/6844903510463479821

const (
	BLACK = iota
	RED
)

type RBNode struct {
	Parent *RBNode
	Left   *RBNode
	Right  *RBNode
	color  int
	Val    int
}

type RBTree struct {
	LEAF *RBNode
	Root *RBNode
	Size uint64
}

func New() *RBTree {
	blankLeaf := &RBNode{color: BLACK}
	return &RBTree{
		LEAF: blankLeaf,
		Root: blankLeaf,
		Size: 0,
	}
}

//节点左旋必须有右孩子，右旋必须有左孩子。
//如果N经过旋转变成了根节点，一定要记得将RBTree结构体中的根节点指针root指向N，这是容易出错的地方
func (tree *RBTree) leftRotate(node *RBNode) *RBNode {
	if node == nil || node.Right == nil {
		return node
	}
	right := node.Right
	node.Right = right.Left
	right.Left = node

	if node.Right != nil {
		node.Right.Parent = node
	}

	if node.Parent == nil {
		tree.Root = right
	} else if node.Parent.Left == node {
		node.Parent.Left = right
	} else if node.Parent.Right == node {
		node.Parent.Right = right
	}
	right.Parent = node.Parent
	node.Parent = right
	return right
}

func (tree *RBTree) rightRotate(node *RBNode) *RBNode {
	if node == nil || node.Left == nil {
		return node
	}
	left := node.Left
	node.Left = left.Right
	left.Right = node

	if node.Left != nil {
		node.Left.Parent = node
	}

	if node.Parent == nil {
		tree.Root = left
	} else if node.Parent.Left == node {
		node.Parent.Left = left
	} else if node.Parent.Right == node {
		node.Parent.Right = left
	}

	left.Parent = node.Parent
	node.Parent = left
	return left
}

func (tree *RBTree) getUncle(node *RBNode) *RBNode {
	gp := tree.getGrandparents(node)
	if gp == nil {
		return nil
	}
	if gp.Left == node.Parent {
		return gp.Right
	} else {
		return gp.Left
	}
}

func (tree *RBTree) getGrandparents(node *RBNode) *RBNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	return node.Parent.Parent
}

func (tree *RBTree) getBrother(node *RBNode) *RBNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node.Parent.Left == node {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (tree *RBTree) getMin(node *RBNode) *RBNode {
	if node.Left != tree.LEAF {
		return tree.getMin(node.Left)
	} else {
		return node
	}
}

func (tree *RBTree) Add(val int) bool {
	_, curr := tree.add(tree.Root, nil, val)
	if curr == nil {
		return false
	}
	tree.Size++
	if tree.Root == tree.LEAF {
		tree.Root = curr
	}
	// 平衡修复
	tree.fixAfterInsertion(curr)
	return true
}

// 二叉树的插入 返回的insertRet为插入的元素
func (tree *RBTree) add(node *RBNode, insert *RBNode, val int) (root, insertRet *RBNode) {
	if node == tree.LEAF {
		insert = &RBNode{
			Parent: nil,
			Left:   tree.LEAF,
			Right:  tree.LEAF,
			color:  RED,
			Val:    val,
		}
		return insert, insert
	}
	if node.Val > val {
		node.Left, insert = tree.add(node.Left, insert, val)
		node.Left.Parent = node
	} else if node.Val < val {
		node.Right, insert = tree.add(node.Right, insert, val)
		node.Right.Parent = node
	}
	return node, insert
}

// 插入后的平衡修复
func (tree *RBTree) fixAfterInsertion(curr *RBNode) {
	if curr.Parent == nil {
		tree.Root.color = BLACK
	} else if curr.Parent.color == RED { // 父节点为红色
		uncle := tree.getUncle(curr)
		grandparents := tree.getGrandparents(curr)
		if uncle.color == RED { // PARENT UNCLE 双红
			curr.Parent.color = BLACK
			uncle.color = BLACK
			grandparents.color = RED
			tree.fixAfterInsertion(grandparents)
		} else {
			case1 := curr == curr.Parent.Left && curr.Parent == grandparents.Left
			case2 := curr == curr.Parent.Right && curr.Parent == grandparents.Right
			if case1 || case2 { //父红同侧
				curr.Parent.color = BLACK
				grandparents.color = RED
				if case1 {
					tree.rightRotate(grandparents)
				} else {
					tree.leftRotate(grandparents)
				}
			} else { // 父红不同侧
				if curr.Parent.Left == curr { //RL
					tree.rightRotate(curr.Parent)
					tree.fixAfterInsertion(curr.Right)
				} else { //LR
					tree.leftRotate(curr.Parent)
					tree.fixAfterInsertion(curr.Left)
				}
			}
		}
	}
}

func (tree *RBTree) Remove(val int) bool {
	root, removed, replace := tree.remove(tree.Root, nil, nil, val)
	if root == nil {
		return false
	}
	tree.Size--
	if removed.color == BLACK && tree.Size > 0 {
		if replace != tree.LEAF {
			tree.fixAfterDelete(replace)
		} else if removed != tree.LEAF {
			tree.fixAfterDelete(removed)
		}
	}
	return true
}

// 两个思想（PS 若删除的为黑色节点）
// 1 找兄弟借一个黑色节点 （颜色互换+旋转）
// 2 让兄弟少一个黑节点 (染红兄弟节点)

// 两种情况
// 1 兄弟节点的子节点全黑 （思想2）
// 2 兄弟节点的子节点最多有一个黑色子节点（思想1）

// 一个技巧
// 1 保证兄弟节点自身为黑色
func (tree *RBTree) fixAfterDelete(curr *RBNode) {
	for tree.Root != curr && curr.color == BLACK {
		brother := tree.getBrother(curr)

		if curr == curr.Parent.Left { // 被删除的是左子节点
			// 保证兄弟节点一定是黑色的
			if brother.color == RED {
				brother.color = BLACK
				curr.Parent.color = RED
				tree.leftRotate(curr.Parent)
				brother = curr.Parent.Right
			}
			// 兄弟节点的两个子子节点都是黑色的
			if brother.Left.color == BLACK && brother.Right.color == BLACK {
				brother.color = RED
				curr = curr.Parent // 向上递归处理
			} else { // 兄弟节点的子节点最多有一个为黑色节点
				if brother.Right.color == BLACK { // 把情况转换成兄弟节点的右侧子节点是红色
					brother.Left.color = BLACK
					brother.color = RED
					tree.rightRotate(brother)
					brother = curr.Parent.Right
				}
				brother.color = curr.Parent.color                     // 复制父节点颜色
				curr.Parent.color, brother.Right.color = BLACK, BLACK // 将父节点和自身的右子节点染黑
				tree.leftRotate(curr.Parent)
				curr = tree.Root
			}

		} else { // 被删除的是右子节点 (镜像操作)
			// 保证兄弟节点一定是黑色的
			if brother.color == RED {
				brother.color = BLACK
				curr.Parent.color = RED
				tree.rightRotate(curr.Parent)
				brother = curr.Parent.Left
			}
			// 兄弟节点的两个子子节点都是黑色的
			if brother.Left.color == BLACK && brother.Right.color == BLACK {
				brother.color = RED
				curr = curr.Parent // 向上递归处理
			} else {
				if brother.Left.color == BLACK {
					brother.Right.color = BLACK
					brother.color = RED
					tree.leftRotate(brother)
					brother = curr.Parent.Left
				}
				brother.color = curr.Parent.color                    // 复制父节点颜色
				curr.Parent.color, brother.Left.color = BLACK, BLACK // 将父节点和自身的左子节点染黑
				tree.rightRotate(curr.Parent)
				curr = tree.Root
			}
		}
	}
	curr.color = BLACK
}

// 二叉树的删除
func (tree *RBTree) remove(node, removed, replace *RBNode, val int) (ret, removedRet, replaceRet *RBNode) {
	if node == tree.LEAF {
		return nil, tree.LEAF, tree.LEAF
	}
	if node.Val > val {
		node.Left, removed, replace = tree.remove(node.Left, removed, replace, val)
	} else if node.Val < val {
		node.Right, removed, replace = tree.remove(node.Right, removed, replace, val)
	} else if node.Val == val {
		if node.Left == tree.LEAF {
			if tree.Root == node {
				tree.Root = node.Right
				tree.Root.Parent = nil
			} else {
				node.Right.Parent = node.Parent
			}
			return node.Right, node, node.Right
		} else if node.Right == tree.LEAF {
			if tree.Root == node {
				tree.Root = node.Left
				tree.Root.Parent = nil
			} else {
				node.Left.Parent = node.Parent
			}
			return node.Left, node, node.Left
		} else {
			// 找到后继节点
			replace := tree.getMin(node.Right)
			replace.Left = node.Left
			replace.Left.Parent = replace
			replace.Right, removed, replace = tree.remove(node.Right, removed, replace, replace.Val)
			if replace.Right != tree.LEAF {
				replace.Right.Parent = replace
			}
			if tree.Root == node {
				tree.Root = replace
			}
			return replace, node, replace
		}
	}
	return node, removed, replace
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
		if node.color == BLACK {
			color = "B"
		} else {
			color = "R"
		}
		fmt.Printf("\t%d(%s) ", node.Val, color)

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

func TestRBtree(t *testing.T) {
	brt := New()
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

	brt.Remove(24)
	brt.Remove(123)

	brt.bfsRBtree()
}
