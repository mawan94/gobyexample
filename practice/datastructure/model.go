package datastructure

import (
	"fmt"
)

/*
	Model  -  FOR TEST
*/

//单向链表
type NodeList struct {
	Val  interface{}
	Next *NodeList
}

//双向链表
type LinkedListNode struct {
	Val        interface{}
	Prev, Next *LinkedListNode
}

// 二叉树
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

//TODO ARRAYLIST  LINKEDLIST       QUEUE  STACK  UPGRAY
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
