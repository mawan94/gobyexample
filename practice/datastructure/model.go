package datastructure

/*
Model  -  FOR TEST
*/

//单向链表
type NodeList struct {
	Val  interface{}
	Next *NodeList
}

//双向链表
type LinkedListNode struct {
	Val        interface{}
	Prev, Next *LinkedListNode
}

// 二叉树
type TreeNode struct {
	Val         interface{}
	Left, Right *TreeNode
}

// 栈
type Stack struct {
	top interface{}
	arr []interface{}
}

// 队列
type Queue struct {
	arr []interface{}
}

// ***************************STACK相关实现**************************************
func (s *Stack) Push(val interface{}) {
	s.top = val
	s.arr = append(s.arr, val)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	ret := s.top
	s.arr = s.arr[:len(s.arr)-1]
	if len(s.arr) == 0 {
		s.top = nil
	} else {
		s.top = s.arr[len(s.arr)-1]
	}
	return ret
}

func (s *Stack) IsEmpty() bool {
	return s.arr == nil || len(s.arr) == 0
}

// *****************************QUEUE相关实现************************************
func (q *Queue) IsEmpty() bool {
	return q.arr == nil || len(q.arr) == 0
}

func (q *Queue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := q.arr[0]
	q.arr = q.arr[1:]
	return ret
}

func (q *Queue) Add(val interface{}) {
	if q.arr == nil {
		q.arr = make([]interface{}, 0)
	}
	q.arr = append(q.arr, val)
}
