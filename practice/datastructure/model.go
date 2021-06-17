package datastructure

import (
	"fmt"
	"sort"
)

/*
	Model  -  FOR TEST
*/

// 单向链表节点
type NodeList struct {
	Val  interface{}
	Next *NodeList
}

// 双向链表节点
type LinkedListNode struct {
	Val        interface{}
	Prev, Next *LinkedListNode
}

// 二叉树节点
type TreeNode struct {
	Val         interface{}
	Left, Right *TreeNode
}

// 栈
type Stack interface {
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
}

// 队列
type Queue interface {
	Enqueue(e interface{})
	Dequeue() interface{}
	IsEmpty() bool
}

// 可比较接口
type Comparable interface {
	compareTo(other Comparable) int
}

// 双向链表实现
type LinkedList struct {
	dummyHead, dummyTail *LinkedListNode
	size                 int
}

func NewLinkedList() *LinkedList {
	dummyHead := &LinkedListNode{}
	dummyTail := &LinkedListNode{}
	dummyHead.Next, dummyTail.Prev = dummyTail, dummyHead
	return &LinkedList{
		dummyHead: dummyHead,
		dummyTail: dummyTail,
	}
}

func NewLinkedListBySlice(arr []interface{}) *LinkedList {
	list := NewLinkedList()
	for i := 0; i < len(arr); i++ {
		list.add(i, arr[i])
	}
	return list
}

func (list *LinkedList) print(printFunc func(interface{})) {
	head := list.dummyHead.Next
	for head != list.dummyTail {
		printFunc(head.Val)
		head = head.Next
	}
}

func (list *LinkedList) get(index int) *LinkedListNode {
	if index < 0 || index >= list.size {
		panic(fmt.Sprintf("不合法的索引: %d", index))
	}
	if index <= list.size/2 {
		curr := list.dummyHead
		for i := 0; i < index; i++ {
			curr = curr.Next
		}
		return curr.Next
	} else {
		curr := list.dummyTail
		for i := list.size - 1; i > index; i++ {
			curr = curr.Prev
		}
		return curr.Prev
	}
}

func (list *LinkedList) set(index int, e interface{}) {
	if index < 0 || index >= list.size {
		panic(fmt.Sprintf("不合法的索引: %d", index))
	}
	if index <= list.size/2 {
		curr := list.dummyHead.Next
		for i := 0; i < index; i++ {
			curr = curr.Next
		}
		curr.Val = e
	} else {
		curr := list.dummyTail.Prev
		for i := list.size - 1; i > index; i-- {
			curr = curr.Prev
		}
		curr.Val = e
	}
}

func (list *LinkedList) add(index int, e interface{}) {
	if index < 0 || index > list.size {
		panic(fmt.Sprintf("不合法的索引: %d", index))
	}
	if index <= list.size/2 {
		curr := list.dummyHead
		for i := 0; i < index; i++ {
			curr = curr.Next
		}
		node := &LinkedListNode{
			Prev: curr,
			Next: curr.Next,
			Val:  e,
		}
		curr.Next.Prev = node
		curr.Next = node

	} else {
		curr := list.dummyTail
		for i := list.size; i > index; i-- {
			curr = curr.Prev
		}
		node := &LinkedListNode{
			Prev: curr.Prev,
			Next: curr,
			Val:  e,
		}
		curr.Prev.Next = node
		curr.Prev = node
	}
	list.size++
}

func (list *LinkedList) addFirst(e interface{}) {
	list.add(0, e)
}

func (list *LinkedList) addLast(e interface{}) {
	list.add(list.size, e)
}

func (list *LinkedList) remove(index int) LinkedListNode {
	if index < 0 || index >= list.size {
		panic(fmt.Sprintf("不合法的索引: %d", index))
	}
	if index <= list.size/2 {
		prev := list.dummyHead
		for i := 0; i < index; i++ {
			prev = prev.Next
		}
		ret := prev.Next
		prev.Next = prev.Next.Next
		prev.Next.Prev = prev
		list.size--
		ret.Next, ret.Prev = nil, nil
		return *ret
	} else {
		prev := list.dummyTail
		for i := list.size - 1; i > index; i-- {
			prev = prev.Prev
		}
		ret := prev.Prev
		prev.Prev = prev.Prev.Prev
		prev.Prev.Next = prev
		list.size--
		ret.Next, ret.Prev = nil, nil
		return *ret
	}

}

func (list *LinkedList) removeFirst() LinkedListNode {
	return list.remove(0)
}

func (list *LinkedList) removeLast() LinkedListNode {
	return list.remove(list.size - 1)
}

func (list *LinkedList) Push(e interface{}) {
	list.addFirst(e)
}

func (list *LinkedList) Pop() interface{} {
	return list.removeFirst().Val
}

func (list *LinkedList) Peek() interface{} {
	return list.get(0).Val
}

func (list *LinkedList) Enqueue(e interface{}) {
	list.addFirst(e)
}

func (list *LinkedList) Dequeue() interface{} {
	return list.removeLast().Val
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList) Len() int {
	return list.size
}

func (list *LinkedList) Less(i, j int) bool {
	iData, jData := list.get(i).Val, list.get(j).Val
	ret := iData.(Comparable).compareTo(jData.(Comparable))
	if ret > 0 {
		return false
	}
	return true
}

func (list *LinkedList) Swap(i, j int) {
	if i == j {
		return
	}
	iData, jData := list.get(i).Val, list.get(j).Val
	list.set(i, jData)
	list.set(j, iData)
}

//func (list *LinkedList) Swap1(i, j int) {
//	// 节点交换
//	var node1 *LinkedListNode = list.dummyHead.Next
//	var node2 *LinkedListNode = list.dummyHead.Next
//	if i < 0 || i >= list.size {
//		panic(fmt.Sprintf("不合法的索引: %d", i))
//	}
//	if j < 0 || j >= list.size {
//		panic(fmt.Sprintf("不合法的索引: %d", j))
//	}
//	if i == j {
//		return
//	}
//
//	index1, index2 := i, j
//	for index1 != 0 || index2 != 0 {
//		if index1 != 0 {
//			node1 = node1.Next
//			index1--
//		}
//		if index2 != 0 {
//			node2 = node2.Next
//			index2--
//		}
//	}
//	if i > j {
//		node1, node2 = node2, node1
//	}
//
//	node1Prev := node1.Prev
//	node2Prev := node2.Prev
//	// 断掉关联
//	node1Prev.Next = nil
//	node2Prev.Next = nil
//
//	node2.Next, node1.Next = node1.Next, node2.Next
//
//}

func (list *LinkedList) sort() {
	sort.Sort(list)
}

func (list *LinkedList) reverse() {
	sort.Reverse(list)
}

// 二分搜索树
// ***************** Binary Search Tree *****************
type BinarySearchTree struct {
	root *TreeNode
	size int
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) add(e Comparable) {
	var addFunc func(*TreeNode, Comparable) *TreeNode
	addFunc = func(node *TreeNode, e Comparable) *TreeNode {
		if node == nil {
			bst.size++
			if bst.root == nil {
				bst.root = &TreeNode{Val: e}
				return bst.root
			}
			return &TreeNode{Val: e}
		}
		if node.Val.(Comparable).compareTo(e) <= -1 {
			node.Right = addFunc(node.Right, e)
		} else if node.Val.(Comparable).compareTo(e) >= 1 {
			node.Left = addFunc(node.Left, e)
		}
		return node
	}
	addFunc(bst.root, e)
}

func (bst *BinarySearchTree) remove(e Comparable) {
	var removeFunc func(*TreeNode, Comparable) *TreeNode
	removeFunc = func(node *TreeNode, comparable Comparable) *TreeNode {
		if node == nil {
			return nil
		} else if node.Val.(Comparable).compareTo(e) <= -1 {
			node.Right = removeFunc(node.Right, e)
			return node
		} else if node.Val.(Comparable).compareTo(e) >= 1 {
			node.Left = removeFunc(node.Left, e)
			return node
		} else if node.Val.(Comparable).compareTo(e) == 0 {
			bst.size--
			if node.Left == nil {
				if bst.root == node {
					bst.root = node.Right
				}
				return node.Right
			} else if node.Right == nil {
				if bst.root == node {
					bst.root = node.Left
				}
				return node.Left
			}
			replacement := node.Left
			for replacement.Right != nil {
				replacement = replacement.Right
			}
			replacement.Left = removeFunc(node.Left, replacement.Val.(Comparable))
			replacement.Right = node.Right
			if bst.root == node {
				bst.root = replacement
			}
			return replacement
		}
		return node
	}
	removeFunc(bst.root, e)
}

func (bst *BinarySearchTree) bfs(printFunc func(e interface{})) {
	println()
	if bst.root == nil {
		return
	}
	var queue Queue = NewLinkedList()
	queue.Enqueue(bst.root)

	for !queue.IsEmpty() {
		node := queue.Dequeue().(*TreeNode)
		printFunc(node.Val)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
}

func (bst *BinarySearchTree) sortOrder(printFunc func(e interface{})) {
	println()
	var midOrder func(node *TreeNode)
	midOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		midOrder(node.Left)
		printFunc(node.Val)
		midOrder(node.Right)
	}
	midOrder(bst.root)
}
