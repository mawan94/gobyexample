package node

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
)

func sortList(head *datastructure.NodeList) *datastructure.NodeList {
	if head == nil || head.Next == nil {
		return head
	}
	s, f := head, head.Next.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	mid := s.Next
	s.Next = nil

	left := sortList(head)
	right := sortList(mid)
	return merge(left, right)
}

// 合并两个有序链表
func merge(l1, l2 *datastructure.NodeList) *datastructure.NodeList {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val.(int) < l2.Val.(int) {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
}

func TestSortNode(t *testing.T) {
	n := &datastructure.NodeList{
		4,
		&datastructure.NodeList{
			2,
			&datastructure.NodeList{
				1,
				&datastructure.NodeList{
					3,
					nil,
				},
			},
		},
	}

	ret := sortList(n)
	fmt.Println(ret)
}
