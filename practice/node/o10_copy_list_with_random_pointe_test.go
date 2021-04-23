package node

var cache = make(map[*Node]*Node)

type Node struct {
	Val int
	Next,Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil{
		return head
	}
	v,ok := cache[head]
	if ok {
		return v
	}
	root := &Node{
		head.Val,
		nil,
			nil,
	}
	cache[head] = root
	root.Next = copyRandomList(head.Next)
	root.Random = copyRandomList(head.Random)
	return root
}