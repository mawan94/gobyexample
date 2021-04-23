package node
//给你一个链表数组，每个链表都已经按升序排列。
//请你将所有链表合并到一个升序链表中，返回合并后的链表。



func mergeKLists(lists []*ListNode) *ListNode {
	return merge11(lists, 0, len(lists) - 1)
}


func merge11(lists []*ListNode,l,r int) *ListNode {
	if l > r {
		return nil
	}
	if l == r {
		return lists[l]
	}
	mid := (l + r) >> 1
	return process11(merge11(lists, l, mid), merge11(lists, mid + 1, r))
}

func process11(l1,l2 *ListNode) *ListNode{
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = process11(l1.Next,l2)
		return l1
	}else {
		l2.Next = process11(l1,l2.Next)
		return l2
	}
}