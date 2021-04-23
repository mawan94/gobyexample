package node

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	return process(head)

}
func process(node *ListNode) *ListNode{
	if node == nil || node.Next == nil {
		return node
	}
	if node.Val != node.Next.Val {
		node.Next = process(node.Next)
		return node
	}else {
		for node != nil && node.Next != nil && node.Val == node.Next.Val {
			node = node.Next
		}
		return process(node.Next)
	}
}

func TestDeleteDuplicates(t *testing.T)  {
	n := &ListNode{
		1,
		&ListNode{
			1,
				nil,
		},
	}
	deleteDuplicates(n)
}
