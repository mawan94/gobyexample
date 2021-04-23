package node

import "gobyexample/practice/datastructure"

// 删除链表倒数第N个节点
func removeNthFromEnd(head *datastructure.NodeList, n int) *datastructure.NodeList {
	src := head
	var step = head
	for i := 0 ;i < n; i++ {
		step = step.Next
	}
	if step == nil {
		return head.Next
	}
	for step.Next != nil {
		src = src.Next
		step = step.Next
	}
	src.Next = src.Next.Next
	return head
}