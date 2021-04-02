package recursion

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
)

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//示例 1：
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//示例 2：
//
//输入：l1 = [], l2 = []
//输出：[]
//示例 3：
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//提示：
//
//两个链表的节点数目范围是 [0, 50]
//-100 <= Node.val <= 100
//l1 和 l2 均按 非递减顺序 排列
func combineTwoOrderedLinkedList(l1, l2 *datastructure.NodeList) *datastructure.NodeList {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if (l1.Val).(int) < (l2.Val).(int) {
		l1.Next = combineTwoOrderedLinkedList(l1.Next, l2)
		return l1
	} else {
		l2.Next = combineTwoOrderedLinkedList(l1, l2.Next)
		return l2
	}
}

func TestCombineTwoOrderedLinkedList(t *testing.T) {
	l1 := &datastructure.NodeList{
		Val: 3,
		Next: &datastructure.NodeList{
			Val: 6,
			Next: &datastructure.NodeList{
				Val: 9,
			},
		},
	}

	l2 := &datastructure.NodeList{
		Val: 2,
		Next: &datastructure.NodeList{
			Val: 7,
		},
	}

	res := combineTwoOrderedLinkedList(l1,l2)
	for res!=nil {
		fmt.Printf("%d -> ",(res.Val).(int))
		res = res.Next
	}
	fmt.Printf("nil\n")

}
