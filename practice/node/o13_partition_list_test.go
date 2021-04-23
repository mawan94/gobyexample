package node

import "testing"

func partition(head *ListNode, x int) *ListNode {
	var l1 = &ListNode{} // l1  <= x
	ans := l1
	var l2  = &ListNode{} // l2  > x
	l2Head := l2

	for head != nil {
		if head.Val < x {
			l1.Next = head
			l1 = l1.Next
		}else {
			l2.Next = head
			l2 = l2.Next
		}
		head = head.Next
	}
	l2.Next = nil
	l1.Next = l2Head.Next
	return ans.Next
}

func TestPartition(t *testing.T) {
	n := &ListNode{
		1,
		&ListNode{
			4,
			&ListNode{
				3,
				&ListNode{
					2,
					&ListNode{
						2,
						&ListNode{
							5,
							&ListNode{
								1,
								nil,
							},
						},
					},
				},
			},
		},
	}
	partition(n,3)

}
