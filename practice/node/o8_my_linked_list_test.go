package node

import (
	"gobyexample/practice/datastructure"
	"testing"
)

type MyLinkedList struct {
	Head,Tail *datastructure.LinkedListNode
	Size int
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this.GetNode(index)
	if node != nil {
		return node.Val.(int)
	}
	return -1
}

func (this *MyLinkedList) GetNode(index int) *datastructure.LinkedListNode {
	// mid = size / 2
	//   if index < mid  从前往后  else 从后往前 （降低复杂度）

	curSize := this.Size
	if index < 0 || index >= curSize {
	    return nil
	}
	mid := curSize / 2
	var node *datastructure.LinkedListNode = nil
	if index <= mid {
	    node = this.Head
	    for i := 0;i < index; i++ {
	        node = node.Next
	    }
	}else {
	    node = this.Tail
	    for i:= (curSize - 1) - index; i > 0; i-- {
	        node = node.Prev
	    }
	}
	return node
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	this.AddAtIndex(0,val)
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	this.AddAtIndex(this.Size,val)
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index == 0 {
		node := &datastructure.LinkedListNode{
			Prev: nil,
			Next: this.Head,
			Val: val,
		}
		if this.Head != nil {
			this.Head.Prev = node
		}else {
			this.Tail = node
		}
		this.Head = node
		this.Size++
	}else if index == this.Size {
		node := &datastructure.LinkedListNode{
			Prev: this.Tail,
			Next: nil,
			Val: val,
		}
		if this.Tail != nil {
			this.Tail.Next = node
		}else {
			this.Head = node
		}
		this.Tail = node
		this.Size++
	}else if index > 0 && index < this.Size {
		tmp := this.GetNode(index)
		prevNode := tmp.Prev
		node := &datastructure.LinkedListNode{
			Prev: prevNode,
			Next: tmp,
			Val: val,
		}
		prevNode.Next = node
		tmp.Prev = node
		this.Size++
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if this.Size == 0 {
		return
	}
	if index == 0 {
		this.Head = this.Head.Next
		this.Size--

	}else if index == this.Size - 1 {
		this.Tail = this.Tail.Prev
		this.Size--
	}else if index > 0 && index < this.Size {
		tmp := this.GetNode(index)
		prevNode := tmp.Prev
		nextNode := tmp.Next

		prevNode.Next = nextNode
		nextNode.Prev = prevNode
		tmp.Prev = nil
		tmp.Next = nil
		this.Size--
	}
}

func TestMyLinkedList(t *testing.T)  {
	list := MyLinkedList{
		nil,
		nil,
		0,
	}
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1,2)
	println(list.Get(1))
	list.DeleteAtIndex(1)
	println(list.Get(1))
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