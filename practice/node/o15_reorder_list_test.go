package node

//给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//示例 1:
//
//给定链表 1->2->3->4, 重新排列为 1->4->2->3.
//示例 2:
//
//给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

func reorderList(head *ListNode) {
	// 1.找到原链表的中点
	// 2.反转后半部分链表
	// 3. 合并链表
	if head == nil || head.Next == nil {
		return
	}
	s, f := head, head.Next.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	tmp := s.Next
	s.Next = nil
	// 后半部分
	l2 := reserve(tmp)
	// 前半部分
	l1 := head

	// 合并两个链表
	head = mergeList(l1, l2, true)
}

func mergeList(l1, l2 *ListNode, flag bool) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if flag {
		l1.Next = mergeList(l1.Next, l2, !flag)
		return l1
	} else {
		l2.Next = mergeList(l1, l2.Next, !flag)
		return l2
	}
}

func reserve(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reserve(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}
