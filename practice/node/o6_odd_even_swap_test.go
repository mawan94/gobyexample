package node

import "gobyexample/practice/datastructure"

func oddEvenList(head *datastructure.NodeList) *datastructure.NodeList {
	// 当没有节点或只有 1 个节点时直接返回
	if head == nil || head.Next == nil {
		return head
	}
	// o1: 奇链表头, e1: 偶链表头, o: 当前奇链表节点, e: 当前偶链表节点
	o1, e1, o, e := head, head.Next, head, head.Next
	for o.Next != nil && e.Next != nil {
		en := e.Next
		enn := en.Next

		o.Next = en  // 当前奇链表节点的下一个节点是当前偶链表节点的下一个节点
		e.Next = enn // 当前偶链表的下一个节点是当前偶链表节点的下下一个节点
		// 前进
		o = o.Next
		e = e.Next
	}
	o.Next = e1 // 把偶链表头连接到奇链表尾
	return o1
}
