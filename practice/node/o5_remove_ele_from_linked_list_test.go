package node

import "gobyexample/practice/datastructure"

// 移除链表值为val的节点

func removeElements(head *datastructure.NodeList, val int) *datastructure.NodeList {
	dummy := &datastructure.NodeList{Val: -1, Next: head}
	cur := dummy
	for cur.Next != nil {
		if  cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return dummy.Next
}