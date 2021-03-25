package recursion

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
)

// 链表翻转 （不开辟额外空间）

// 1 -> 2 -> 3 -> nil
// 1 -> 2 <- 3
// 1为head； 让2的下一个指向1； 让1和后面断开关联；返回reverseHead（ps 当前是3打头的 换头操作）
func reverseNode(head *datastructure.NodeList) *datastructure.NodeList {
	if head == nil || head.Next == nil {
		return head
	}
	reverseHead := reverseNode(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reverseHead
}

func TestReverseNode(t *testing.T) {
	node := &datastructure.NodeList{
		1,
		&datastructure.NodeList{
			2,
			&datastructure.NodeList{
				3,
				nil,
			},
		},
	}
	node = reverseNode(node)

	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println()

}
