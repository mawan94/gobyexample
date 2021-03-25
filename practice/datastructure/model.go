package datastructure

/*
Model  -  FOR TEST
*/



//单向链表
type NodeList struct {
	Val  interface{}
	Next *NodeList
}

func GenerateRandomNodeList(len int) *NodeList {
	return nil
}




//双向链表
type LinkedList struct {
	Val      interface{}
	Previous *LinkedList
	Next     *LinkedList
}

func GenerateRandomLinkedList(len int) *LinkedList {
	return nil
}




// 二叉树
type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func GenerateRandomBinaryTree(len int) *TreeNode {
	return nil
}

// *****************************************************************
