package node

import (
	"fmt"
	"testing"
)

type NodeList struct {
	Val  int
	Next *NodeList
}

type MyNodeList struct {
	Head *NodeList
	Size       int
}

/** Initialize your data structure here. */
func Constructor() MyNodeList {
	return MyNodeList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyNodeList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	res := this.Head
	for i := 0; i < index; i++ {
		res = res.Next
	}
	return res.Val

	//if index < 0 || index >= this.Size {
	//	return -1
	//}
	//if index == 0 {
	//	return this.Head.Val
	//}
	//tmp :=  &MyLinkedList{
	//	this.Head,
	//	this.Size,
	//}
	//tmp.Head = tmp.Head.Next
	//return tmp.Get(index - 1)
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyNodeList) AddAtHead(val int) {
	newHead := &NodeList{val, this.Head}
	this.Head = newHead
	this.Size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyNodeList) AddAtTail(val int) {
	this.AddAtIndex(this.Size ,val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyNodeList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
	} else if index > this.Size {
		return
	} else {
		dh := &NodeList{val, this.Head}
		for i := 0; i < index; i++ {
			dh = dh.Next
		}
		dh.Next = &NodeList{val, dh.Next}
		this.Size++
	}

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyNodeList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	dh := &NodeList{-1, this.Head}

	for i := 0; i < index; i++ {
		dh = dh.Next
	}
	if dh.Next == this.Head {
		this.Head = dh.Next.Next
	}

	dh.Next = dh.Next.Next
	this.Size --
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func TestABC(t *testing.T)  {
	linkedList := MyNodeList{}

	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1,2)
	fmt.Println(linkedList.Get(1))
	linkedList.DeleteAtIndex(1)
	fmt.Println(linkedList.Get(1))


	fmt.Println(linkedList)
}
