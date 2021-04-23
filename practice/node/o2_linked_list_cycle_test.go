package node

import (
	"gobyexample/practice/datastructure"
	"testing"
)
// 快慢指针判断是否存在环
func hasCycle(head *datastructure.NodeList) bool {
	if head == nil || head.Next == nil {
		return false
	}
	s := head.Next
	f := head.Next.Next
	for s != nil && f != nil && f.Next != nil {
		if s == f {
			return true
		}
		s = s.Next
		f = f.Next.Next
	}
	return false
}

// 返回环链表的入环节点  当相交的时候快指针返回head 再同时一次走一步 再次相遇的时候就是入环点
func detectCycle(head *datastructure.NodeList) *datastructure.NodeList {
	if head == nil || head.Next == nil {
		return nil
	}
	s := head.Next
	f := head.Next.Next

	flag := true

	for f != nil && f.Next != nil {
		if s == f {
			if flag {
				f = head
				flag = !flag
			}else {
				return s
			}
		}else {
			s = s.Next
			if flag {
				f = f.Next.Next
			}else {
				f = f.Next
			}
		}
	}
	return nil
}

func TestHasCycle(t *testing.T) {

}
