package node

import "gobyexample/practice/datastructure"

// 返回链表相交的点
func getIntersectionNode(headA, headB *datastructure.NodeList) *datastructure.NodeList {
	a := headA
	b := headB
	flag := 0
	for flag <= 2 {
		if a == b {
			return a
		}
		if a != nil && b!= nil {
			a = a.Next
			b = b.Next
		}else {
			if a == nil {
				flag ++
				a = headB
				b = b.Next
			}
			if b == nil {
				flag ++
				b = headA
				a = a.Next
			}
		}
	}
	return nil
}