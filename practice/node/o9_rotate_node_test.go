package node

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
)

func rotateRight(head *datastructure.NodeList, k int) *datastructure.NodeList {
	if head == nil || k == 0 || head.Next == nil {
		return head
	}
	var len = 0
	tmp := head
	for tmp != nil {
		len++
		tmp = tmp.Next
	}
	k = k % len
	if k == 0 {
		return head
	}
	// s看是尾指针  s.next是new_head
	s, f := head, head
	for i := 0; i < k; i++ {
		f = f.Next
	}
	for f.Next != nil {
		s = s.Next
		f = f.Next
	}
	f.Next = head
	for s.Next != head {
		head = head.Next
	}
	s.Next = nil
	return head
}

func TestRotate(t *testing.T) {
	n := &datastructure.NodeList{
		1,
		&datastructure.NodeList{
			2,
			&datastructure.NodeList{
				3,
				nil,
			},
		},
	}
	ret := rotateRight(n, 1)
	for ret != nil {
		fmt.Printf("%d -> ",ret.Val)
		ret = ret.Next
	}

}
