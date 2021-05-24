package tree

import (
	"fmt"
	"testing"
)

type Entry struct {
	key   int
	value interface{}
}

type BNode struct {
	entries  []Entry
	parent   *BNode
	children []*BNode
	isLeaf   bool
}

type BTree struct {
	Root *BNode
	m    int // 阶数
	size int
}

func NewBTree(m int) *BTree {
	if m <= 2 {
		panic("m 不合法")
	}
	return &BTree{m: m}
}

// 分裂
func (tree *BTree) split(curr *BNode) {
	if curr != nil && len(curr.entries) >= tree.m {
		entries := curr.entries

		midIndex := (tree.m) / 2
		midEntry := entries[midIndex]
		// 构建分裂后新的左右孩子
		left := &BNode{
			entries: entries[:midIndex],
			isLeaf:  curr.isLeaf,
		}
		right := &BNode{
			entries: entries[midIndex+1:],
			isLeaf:  curr.isLeaf,
		}
		if !curr.isLeaf {
			left.children = curr.children[:midIndex+1]
			right.children = curr.children[midIndex+1:]
		}
		// 更新下层parent - 上层分裂后的产物
		for _, e := range left.children {
			e.parent = left
		}
		for _, e := range right.children {
			e.parent = right
		}
		parent := curr.parent
		if parent == nil {
			root := &BNode{
				entries:  []Entry{midEntry},
				isLeaf:   false,
				children: []*BNode{left, right},
			}
			root.children[0].parent, root.children[1].parent = root, root
			tree.Root = root
			return
		} else {
			// 组装parent
			parentsEntries := parent.entries
			parentsChildren := parent.children

			// 找到合适的位置插入到parentsEntries里面
			targetInsertIndex := findTargetInsertIndex(midEntry, parentsEntries)

			// insert midEntry 2 parent entries
			parent.entries = append(parentsEntries[:targetInsertIndex], append([]Entry{midEntry}, parentsEntries[targetInsertIndex:]...)...)
			// replace new children
			left.parent, right.parent = parent, parent // 如果parent不需要分裂则parent就不需要改动  要是需要分裂会在上层for_range做处理
			parent.children = append(parentsChildren[:targetInsertIndex], append([]*BNode{left, right}, parentsChildren[targetInsertIndex+1:]...)...)
			tree.split(parent)
		}
	}
}

// 在有序数组entries中找到 比entry大最少的索引
func findTargetInsertIndex(entry Entry, entries []Entry) int {
	targetIndex := 0
	for ; len(entries) > targetIndex && entry.key > entries[targetIndex].key; targetIndex++ {
	}
	return targetIndex
}

func (tree *BTree) insert(node *BNode, entry Entry) {
	if node == nil {
		tree.Root = &BNode{
			entries: []Entry{entry},
			isLeaf:  true,
		}
		return
	}
	// 找到目标插入点
	targetIndex := findTargetInsertIndex(entry, node.entries)
	// 叶子节点 尝试插入
	if node.isLeaf {
		// 在第targetIndex个位置插入data
		//node.entries = append(node.entries[:targetIndex],append([]int{entry}, node.entries[targetIndex:]...)...)  FIXME  存在内存污染 不能这么使用
		newEntries := make([]Entry, targetIndex)
		copy(newEntries, node.entries[:targetIndex])
		node.entries = append(newEntries, append([]Entry{entry}, node.entries[targetIndex:]...)...)
		// 检测是否需要分裂
		if len(node.entries) >= tree.m {
			tree.split(node)
		}
	} else { // 非叶子节点
		tree.insert(node.children[targetIndex], entry)
	}
}

func (tree *BTree) Insert(entry Entry) {
	tree.insert(tree.Root, entry)
	tree.size++
}

func (tree *BTree) Delete(entry Entry) {
	tree.delete(tree.Root, entry)
}

func (tree *BTree) findByKey(curr *BNode, key int) *Entry {
	if curr == nil {
		return nil
	}

	for i := 0; i < len(curr.entries); i++ {
		if curr.entries[i].key == key {
			return &curr.entries[i]
		}

		if key < curr.entries[i].key {
			return tree.findByKey(curr.children[i], key)
		}

		if i == len(curr.entries)-1 && curr.children != nil {
			return tree.findByKey(curr.children[len(curr.children)-1], key)
		}
	}

	return nil
}

func (tree *BTree) delete(node *BNode, entry Entry) {
	if node == nil {
		return
	}
	for i := 0; i < len(node.entries); i++ {
		if entry.key <= node.entries[i].key {
			if entry == node.entries[i] { // 目标元素 进行删除操作
				if node == tree.Root { // 直接删除
					node.entries = append(node.entries[:i], append(node.entries[i+1:])...)
					if len(node.entries) == 0 {
						node.isLeaf = true
					}
					return
				} else if node.isLeaf {
					tree.deleteLeafNode(node, entry)
				} else {
					// 找到其可以代替的节点 （前躯或者后继）
					replacement := node.children[i]
					for !replacement.isLeaf {
						replacement = replacement.children[len(replacement.children)-1]
					}
					node.entries[i] = replacement.entries[0]
					tree.deleteLeafNode(replacement, replacement.entries[0])
				}
			} else {
				tree.delete(node.children[i], entry)
			}
			break
		}
		if i == len(node.entries)-1 && node.children != nil {
			tree.delete(node.children[len(node.children)-1], entry)
			break
		}
	}
}

func (tree *BTree) deleteLeafNode(node *BNode, entry Entry) {
	tree.size--
	//case1: ceil(m/2) - 1 > entries size    :直接删除
	if len(node.entries) > tree.m%2+tree.m/2-1 {
		//TODO  这里可以用二分法快速定位
		targetIndex := 0
		for ; node.entries[targetIndex] != entry; targetIndex++ {
		}
		node.entries = append(node.entries[:targetIndex], node.entries[targetIndex+1:]...)
	} else {
		parent := node.parent
		// 找到node的兄弟
		nodeIndex := 0 // currNode在当前parent里面的索引位置
		for ; parent.children[nodeIndex] != node; nodeIndex++ {
		}
		var leftBrother, rightBrother *BNode
		if nodeIndex > 0 {
			leftBrother = parent.children[nodeIndex-1]
		}
		if len(parent.children)-1 > nodeIndex {
			rightBrother = parent.children[nodeIndex+1]
		}

		caseLeft := leftBrother != nil && len(leftBrother.entries) > tree.m%2+tree.m/2-1
		caseRight := rightBrother != nil && len(rightBrother.entries) > tree.m%2+tree.m/2-1
		// case2 : 找兄弟节点借一个
		if caseLeft || caseRight {
			targetIndex := 0
			for ; node.entries[targetIndex] != entry; targetIndex++ {
			}
			node.entries[targetIndex] = parent.entries[nodeIndex]
			if caseLeft {
				parent.entries[nodeIndex] = leftBrother.entries[len(leftBrother.entries)-1]
				leftBrother.entries = append(leftBrother.entries[:len(leftBrother.entries)-1])
			} else {
				parent.entries[nodeIndex] = rightBrother.entries[0]
				rightBrother.entries = append(rightBrother.entries[1:])
			}
		} else {
			// case3 节点向上合并  递归处理
			var f = func(*BNode) {}
			f = func(curr *BNode) {

				if curr.parent == nil {
					tree.Root = curr.children[0]
					tree.Root.parent = nil
					tree.Root.isLeaf = false
					return
				}

				parent := curr.parent
				nodeIndex := 0 // currNode在当前parent（child）里面的索引位置
				for ; parent.children[nodeIndex] != curr; nodeIndex++ {
				}

				fmt.Println(len(parent.children)>>1, nodeIndex)
				if (len(parent.children)>>1)-1 <= nodeIndex { // curr在parent.child右半部分   parent entry放后面
					entry := parent.entries[nodeIndex-1]
					brother := parent.children[nodeIndex-1]
					mergeNode := &BNode{
						entries:  append(brother.entries, entry),
						parent:   parent,
						isLeaf:   node.isLeaf,
						children: append(brother.children, curr.children...), //brother 的孩子和 curr的孩子
					}
					parent.entries = append(parent.entries[:nodeIndex-1], append(parent.entries[nodeIndex:])...)
					parent.children = append(parent.children[:nodeIndex-1], append([]*BNode{mergeNode}, parent.children[nodeIndex+1:]...)...)
					// 修改同步parent
					for _, child := range mergeNode.children {
						child.parent = mergeNode
					}

					if len(parent.entries) < tree.m%2+tree.m/2-1 {
						f(parent)
					}
				} else { // curr在parent.child左半部分 parent entry放前面
					entry := parent.entries[nodeIndex]
					brother := parent.children[nodeIndex+1]
					mergeNode := &BNode{
						entries:  append([]Entry{entry}, brother.entries...),
						parent:   parent,
						isLeaf:   node.isLeaf,
						children: append(curr.children, brother.children...), //brother 的孩子和 curr的孩子
					}
					// 删除nodeIndex位置元素
					parent.entries = append(parent.entries[:nodeIndex], append(parent.entries[nodeIndex+1:])...)
					// 替代孩子
					parent.children = append(parent.children[:nodeIndex], append([]*BNode{mergeNode}, parent.children[nodeIndex+2:]...)...)
					// 修改同步parent
					for _, child := range mergeNode.children {
						child.parent = mergeNode
					}
					if len(parent.entries) < tree.m%2+tree.m/2-1 {
						f(parent)
					}
				}
			}
			f(node)
		}
	}
}

func TestBTree(t *testing.T) {
	bTree := NewBTree(8)

	for i := 1; i <= 1000000; i++ {
		bTree.Insert(Entry{key: i, value: i})
	}
	fmt.Println(bTree)
	entry := bTree.findByKey(bTree.Root,123456)
	fmt.Println(entry)

}
