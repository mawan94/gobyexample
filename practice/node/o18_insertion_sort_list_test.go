package node

func insertionSortList(head *ListNode) *ListNode {
	// 只有一个元素就无须排序
	if head == nil {
		return head
	}

	// 维护三个指针 leftHead   pre  cur
	leftHead, pre, cur := &ListNode{Next: head}, head, head.Next

	for cur != nil {
		if pre.Val <= cur.Val {
			pre = pre.Next
		} else {
			tmp := leftHead
			for tmp.Next.Val <= cur.Val {
				tmp = tmp.Next
			}
			pre.Next = cur.Next
			cur.Next = tmp.Next
			tmp.Next = cur
		}
		cur = pre.Next
	}
	return leftHead.Next
}
