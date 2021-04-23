package node

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s, f := head, head
	for i := 1; i < k; i++ {
		if f.Next == nil {
			return head
		} else {
			f = f.Next
		}
	}
	// 断开连接
	waitProcess := f.Next
	f.Next = nil
	// 翻转链表
	reverseNode(s)

	s.Next = reverseKGroup(waitProcess, k)
	return f
}
