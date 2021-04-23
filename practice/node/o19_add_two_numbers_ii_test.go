package node

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 找到两个链表的长度
	len1,len2 := 0,0
	l1Tmp,l2Tmp := l1,l2

	for l1Tmp != nil {
		len1 ++
		l1Tmp = l1Tmp.Next
	}
	for l2Tmp != nil {
		len2 ++
		l2Tmp = l2Tmp.Next
	}

	// 边界讨论
	if len1 == 0 {
		return l2
	}else if len2 == 0 {
		return l1
	}else if len1 == 0 && len2 == 0 {
		return nil
	}

	// ==============
	if len1 > len2 {
		res,step := processII(l1,l2,len1,len2)
		if step != 0 {
			return &ListNode{1,res}
		}else {
			return res
		}
	}else {
		res,step := processII(l2,l1,len2,len1)
		if step != 0 {
			return &ListNode{1,res}
		}else {
			return res
		}
	}
}


func processII (long, short *ListNode, len1, len2 int) (*ListNode, int) {
	// 最小边界解
	if len1 == 1 && len2 == 1 {
		total := long.Val + short.Val
		if total >= 10 {
			long.Val = total % 10
			return long, 1
		}else {
			long.Val = total
			return long, 0
		}
	}
	// long 先走差值步
	if len1 > len2 {
		next,step := processII(long.Next, short, len1 - 1, len2)
		long.Next = next
		if long.Val + step >= 10 {
			long.Val = (long.Val + step) % 10
			return long, 1
		}else {
			long.Val = long.Val + step
			return long, 0
		}
	}else {
		// 长度相等的情况
		next,step := processII(long.Next, short.Next, len1 - 1, len2 - 1)
		long.Next = next
		if long.Val + short.Val +  step >= 10 {
			long.Val = (long.Val + short.Val + step) % 10
			return long, 1
		}else {
			long.Val = long.Val + short.Val + step
			return long, 0
		}
	}
}