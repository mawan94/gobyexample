package recursion

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
)

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]

//[1,2,3,4]
// [1,2,4,3]
// 2指向1
// 1指向4
func swapPairs(head *datastructure.NodeList) *datastructure.NodeList {
	if head == nil || head.Next == nil {
		return head
	}
	// 需要预先保存一下head 和 head.next （递归结束合并结果时候会使用到）
	nextNode := head.Next

	result := swapPairs(head.Next.Next)
	nextNode.Next = head
	head.Next = result
	return nextNode
}

func TestSwapPairs(t *testing.T)  {
	node := &datastructure.NodeList{
		1,
		&datastructure.NodeList{
			2,
			&datastructure.NodeList{
				3,
				&datastructure.NodeList{
					4,
					nil,
				},
			},
		},
	}
	node = swapPairs(node)
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println()
}
