package tree

import (
	"fmt"
	"gobyexample/practice/datastructure"
	"testing"
	"unicode/utf8"
)

type TrieNode struct {
	isWord   bool
	value    interface{}
	children map[string]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{Root: &TrieNode{children: make(map[string]*TrieNode)}}
}

func (trie *Trie) Delete(str string) {
	curr := trie.Root
	stack := datastructure.NewLinkedList()
	for i := 0; i < utf8.RuneCountInString(str); i++ {
		if n, ok := curr.children[string([]rune(str)[i:i+1])]; ok {
			// 删除中间节点
			if i == utf8.RuneCountInString(str)-1 && len(n.children) > 0 {
				n.isWord = false
				return
			} else {
				stack.Push(n)
				curr = n
			}
		} else {
			return
		}
	}
	// 删除叶子节点
	var f func(string, int, *TrieNode)
	f = func(str string, index int, root *TrieNode) {
		if root.isWord {
			return
		}
		delete(root.children, string([]rune(str)[index:index+1]))
		if len(root.children) == 0 && !stack.IsEmpty() {
			f(str, index-1, stack.Pop().(*TrieNode))
		}
	}
	f(str, utf8.RuneCountInString(str)-1, stack.Pop().(*TrieNode))
}

func (trie *Trie) Insert(str string) {
	curr := trie.Root
	index := 0
	for _, c := range str {
		// 否来到了最后一个字符（单词的结尾）
		var isLast = false
		if index == utf8.RuneCountInString(str)-1 {
			isLast = true
		}
		// 如果已经存在了当前字符就继续往下找
		if node, ok := curr.children[string(c)]; ok {
			if isLast { //
				node.isWord = true
			}
			curr = node
		} else { // 不存在就创建一个 放到curr的children里面
			newNode := &TrieNode{
				isWord:   isLast,
				value:    string(c),
				children: make(map[string]*TrieNode),
			}
			curr.children[string(c)] = newNode
			curr = newNode
		}
		index++
	}
}

// 完整词语
func (trie *Trie) Search(str string) bool {
	curr := trie.Root
	index := 0
	for _, c := range str {
		node := curr.children[string(c)]
		if node == nil {
			return false
		} else {
			curr = node
		}
		// 是否为完整单词
		if index == utf8.RuneCountInString(str)-1 && node.isWord {
			return true
		}
		index++
	}
	return false
}

// 综合模糊查询
func (trie *Trie) Like(suffix string) []string {
	ret := make([]string, 0)
	var f func(*TrieNode, int, string)

	// currIndex  表示当前处理的字符
	f = func(node *TrieNode, currIndex int, path string) {
		// 成功在某一条路径上匹配完了 suffix的所有字符
		if currIndex == utf8.RuneCountInString(suffix) {
			// 成功找到一条数据
			ret = append(ret, trie.PreLike(path)...)
			return
		}
		if node == nil {
			return
		}
		children := node.children
		for _, v := range children {
			// 如果找到了就将目标定位到下一个字符（新的目标任务）
			if v.value.(string) == string([]rune(suffix)[currIndex:currIndex+1]) {
				f(v, currIndex+1, path+v.value.(string))
			} else { // 如果没有找到就继续在子路径上向下找
				f(v, currIndex, path+v.value.(string))
			}
		}
	}
	f(trie.Root, 0, "")
	return ret
}

// 前缀（abc%）
func (trie *Trie) PreLike(prefix string) []string {
	ret := make([]string, 0)
	curr := trie.Root
	index := 0
	for _, c := range prefix {
		node := curr.children[string(c)]
		// 没有任何下属子节点直接返回
		if node == nil {
			return ret
		} else {
			curr = node
		}
		// 如果最后一个字符正好是一个字符的结尾，则将其先加入返回值中
		if index == utf8.RuneCountInString(prefix)-1 && node.isWord {
			ret = append(ret, prefix)
		}
		index++
	}

	// 定义一个函数 ： 找到当前节点下所有的单词
	var f func(*TrieNode, string)
	f = func(node *TrieNode, path string) {
		children := node.children
		for k, n := range children {
			if n.isWord {
				// 加入结果集
				ret = append(ret, prefix+path+k)
				if len(n.children) > 0 { // 继续向下拼接
					f(n, path+k)
				}
			} else {
				f(n, path+k)
			}
		}
	}
	f(curr, "")
	return ret
}

func TestTrie(t *testing.T) {
	tire := NewTrie()
	tire.Insert("我啊哦s")
	tire.Insert("我啊s")
	tire.Insert("😸啊😬")
	tire.Insert("嘻嘻")
	tire.Insert("哈哈")
	tire.Insert("哈哈哈")

	tire.Delete("哈哈哈")
	fmt.Println("========")

	println(tire.Search("嘻"))  // f
	println(tire.Search("嘻嘻")) // t
	println(tire.Search("哈哈")) // t

	fmt.Println("========")
	ret := tire.PreLike("嘻嘻")
	for _, v := range ret {
		fmt.Println(v)
	}

	fmt.Println("========")
	ret = tire.Like("啊")
	for _, v := range ret {
		fmt.Println(v)
	}
}
