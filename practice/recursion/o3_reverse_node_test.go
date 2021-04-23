package recursion

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
)

// 链表翻转 （不开辟额外空间）

// 1 -> 2 -> 3 -> nil
// 1 -> 2 <- 3
// 1为head； 让2的下一个指向1； 让1和后面断开关联；返回reverseHead（ps 当前是3打头的 即换头操作）

//如果是后续遍历，那么你可以想象后面的链表都处理好了，怎么处理的不用管。
func reverseNode(head *datastructure.NodeList) *datastructure.NodeList {
	if head == nil || head.Next == nil {
		return head
	}
	reverseHead := reverseNode(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reverseHead
}

//如果是前序遍历，那么你可以想象前面的链表都处理好了，怎么处理的不用管。
func reverseNodeV2(head,pre *datastructure.NodeList) *datastructure.NodeList {
	if head == nil{
		return pre
	}
	next := head.Next// 后一段待处理的链表
	head.Next = pre
	return reverseNodeV2(next,head)
}

func reverseNodeV3(head *datastructure.NodeList) *datastructure.NodeList {
	var pre *datastructure.NodeList
	cur := head
	tail := head

	for tail.Next != nil {
		tail = tail.Next
	}

	for cur != tail {
		tmp := cur.Next
		cur.Next = pre

		pre = cur
		cur = tmp
	}
	tail.Next = pre
	return tail
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
	//node = reverseNodeV2(node,nil)
	//node = reverseNodeV3(node)


	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println()

}
