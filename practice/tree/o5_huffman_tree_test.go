package tree

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
	"unicode/utf8"
)

type HuffmanNode struct {
	left, right *HuffmanNode
	data        byte
	weight      int
}

type HuffmanNodeList []HuffmanNode

func (list HuffmanNodeList) Len() int {
	return len(list)
}
func (list HuffmanNodeList) Less(i, j int) bool {
	return list[i].weight < list[j].weight
}
func (list HuffmanNodeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

type HuffmanTree struct {
	Root *HuffmanNode
}

// 构建哈夫曼树
func CreateHuffmanTree(bytes []byte) *HuffmanTree {
	// k: byte ,v: 字符重复出现的次数
	byteMap := make(map[byte]int)
	for _, b := range bytes {
		if count, ok := byteMap[b]; ok {
			byteMap[b] = count + 1
		} else {
			byteMap[b] = 1
		}
	}

	// 将map中的元素按照weight从小到大进行排序 放入list中
	list := make([]HuffmanNode, len(byteMap))

	currIndex := 0
	for k, v := range byteMap {
		node := HuffmanNode{
			data:   k,
			weight: v,
		}
		list[currIndex] = node
		currIndex++
	}
	sort.Sort(HuffmanNodeList(list))

	// 递归函数构建哈夫曼树
	var f func([]HuffmanNode) *HuffmanNode
	f = func(list []HuffmanNode) *HuffmanNode {
		if list == nil || len(list) == 0 {
			return nil
		}
		if len(list) == 1 {
			return &list[0]
		}
		l := list[0]
		r := list[1]
		node := HuffmanNode{
			left:   &l,
			right:  &r,
			weight: l.weight + r.weight,
		}
		list = list[2:]
		list = append(list, node)
		sort.Sort(HuffmanNodeList(list))
		return f(list)
	}
	return &HuffmanTree{
		Root: f(list),
	}
}

// 根据哈弗曼树获取哈夫曼编码
func HuffmanCode(tree *HuffmanTree) map[byte]string {
	hfmCode := make(map[byte]string)
	var f func(*HuffmanNode, string)
	f = func(node *HuffmanNode, path string) {
		if node != nil {
			if node.left != nil {
				f(node.left, path+"0")
			} else {
				hfmCode[node.data] = path
			}

			if node.right != nil {
				f(node.right, path+"1")
			} else {
				hfmCode[node.data] = path
			}
		}
	}
	f(tree.Root, "")
	return hfmCode
}

func Encode(table map[byte]string, src []byte) ([]byte, int) {
	str := ""
	for _, b := range src {
		str += table[b]
	}
	var retLen int
	if utf8.RuneCountInString(str)%8 == 0 {
		retLen = utf8.RuneCountInString(str) / 8
	} else {
		retLen = (utf8.RuneCountInString(str) / 8) + 1
	}

	ret := make([]byte, retLen)
	for i, j := 0, 0; i < len(ret); i, j = i+1, j+8 {
		if j+8 < len(str) {
			n, _ := strconv.ParseInt(subStr(str, j, j+8), 2, 32)
			ret[i] = byte(n)
		} else {
			n, _ := strconv.ParseInt(subStr(str, j, len(str)), 2, 32)
			ret[i] = byte(n)
		}
	}
	return ret, len(str)
}

func Decode(bytes []byte, table map[byte]string, length int) []byte {
	str := ""
	for i := 0; i < len(bytes); i++ {
		if i == len(bytes)-1 {
			for j := utf8.RuneCountInString(str) + utf8.RuneCountInString(fmt.Sprintf("%b", bytes[i])); j < length; j++ {
				str += "0"
			}
			str += fmt.Sprintf("%b", bytes[i])
		} else {
			for j := utf8.RuneCountInString(fmt.Sprintf("%b", bytes[i])); j < 8; j++ {
				str += "0"
			}
			str += fmt.Sprintf("%b", bytes[i])
		}
	}
	decodeMap := make(map[string]byte)
	for k, v := range table {
		decodeMap[v] = k
	}
	// 每个前缀不一样
	ret := make([]byte,0)

	var f func(int, int)
	f = func(begin, end int) {
		if begin > end {
			return
		}
		if b, ok := decodeMap[string([]rune(str)[begin:end])];ok {
			ret = append(ret, b)
			f(begin + (end - begin),utf8.RuneCountInString(str))
		}else {
			f(begin,end - 1)
		}
	}
	f(0,utf8.RuneCountInString(str))
	return ret
}

func subStr(str string, begin, end int) string {
	list := []rune(str)
	list = list[begin:end]
	return string(list)
}

func TestHuffman(t *testing.T) {
	s := "是撒撒😜😜😜😜"
	src := []byte(s)
	tree := CreateHuffmanTree(src)
	codeTable := HuffmanCode(tree)
	huffmanBytes, len := Encode(codeTable, []byte(s))
	ret := Decode(huffmanBytes, codeTable, len)
	fmt.Println(string(ret))
}
