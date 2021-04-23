package node

import "gobyexample/practice/datastructure"

// 回文链表	Version 1
// 1. 利用快慢指针找到中间位置
// 2. 反转后半部分
// 3. 迭代比较
func isPalindrome(head *datastructure.NodeList) bool {
	if head == nil || head.Next == nil {
		return true
	}
	s := head     // 慢指针
	f := head.Next.Next   // 快指针

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	var reverseHead *datastructure.NodeList = nil

	// 如果f ！= nil 说明链表个数为基数， 否则为偶数个
	if f != nil { // 基数长度
		// reverseHead := s.Next.Next // 回文头 (待反转)
		reverseHead = reverse(s.Next.Next)
	}else {
		// reverseHead := s.Next
		reverseHead = reverse(s.Next)
	}

	for head !=nil && reverseHead!= nil {
		if head.Val != reverseHead.Val {
			return false
		}
		head = head.Next
		reverseHead = reverseHead.Next
	}
	return true
}


func reverse(head  *datastructure.NodeList) *datastructure.NodeList {
	if head == nil || head.Next == nil {
		return head
	}
	rh := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rh
}


// 回文链表	Version 2
// 递归首尾比较
func isPalindromeV2(head *datastructure.NodeList) bool {
	var f func(cur *datastructure.NodeList) bool
	f = func (cur *datastructure.NodeList) bool {
		if cur != nil {
			if !f(cur.Next) {
				return false
			}
			if cur.Val != head.Val {
				return false
			}
			head = head.Next
		}
		return true
	}
	return f(head)
}