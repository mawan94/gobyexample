package recursion

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
)

//给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
// 实例:
//输入：3
//输出：
//[
//  [1,null,3,2],
//  [3,2,null,1],
//  [3,1,null,null,2],
//  [2,1,3],
//  [1,null,2,null,3]
//]
//解释：
//以上的输出对应以下 5 种不同结构的二叉搜索树：
//
//	1         3     3      2      1
//	\       /     /      / \      \
//	3     2     1      1   3      2
//	/     /       \                 \
//	2     1         2                 3

func generate(s, e int) []*datastructure.TreeNode {
	all := make([]*datastructure.TreeNode, 0)
	if s > e {
		all = append(all, nil)
		return all
	}
	//i : currentRoot
	for currentRoot := s; currentRoot <= e; currentRoot++ {
		leftAll := generate(s, currentRoot - 1)
		rightAll := generate(currentRoot + 1, e)
		for _, leftNode := range leftAll {
			for _, rightNode := range rightAll {
				root := &datastructure.TreeNode{
					Val:   currentRoot,
					Left:  leftNode,
					Right: rightNode,
				}
				all = append(all, root)
			}
		}
	}
	return all
}

func generateTrees(n int) []*datastructure.TreeNode {
	if n < 1 {
		return make([]*datastructure.TreeNode, 0)
	}
	return generate(1, n)
}

func TestGenerateTrees(t *testing.T) {
	res := generateTrees(3)
	fmt.Println(res)
}
