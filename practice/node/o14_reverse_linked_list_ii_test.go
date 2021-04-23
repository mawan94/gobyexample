package node

import (
	"testing"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// a  :  left的前一个
	// b  :  left
	// c  :  right
	// d  :  right 的后一个
	dummyNode := &ListNode{-1, head}
	a := dummyNode
	for i := 0; i < left - 1; i++ {
		a = a.Next
	}
	c := a
	for i := 0; i < 1 + right - left; i++ {
		c = c.Next
	}

	b := a.Next
	d := c.Next

	a.Next = nil
	c.Next = nil

	reverseNode(b)
	b.Next = d
	a.Next = c
	return dummyNode.Next
}


func reverseBetweenV2(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{-1, head}
	pre := dummyNode
	for i := 0; i < left - 1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right - left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}


func reverseNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseNode(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

func TestReverseBetween(t *testing.T) {
	n := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	reverseBetween(n, 2, 4)
}
