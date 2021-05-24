package tree

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
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
	// todo 优化
	for i, j := 0, 0; i < len(ret); i, j = i+1, j+8 {
		if j+8 < len(str) {
			ret[i] = binaryString2Byte(subString(str, j, j+8))
		} else {
			ret[i] = binaryString2Byte(subString(str, j, len(str)))
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
			str += byte2BinaryString(bytes[i])
		}
	}
	//最长编码的长度
	var hfmCodeMaxLen int
	decodeMap := make(map[string]byte)
	for k, v := range table {
		if len(v) > hfmCodeMaxLen{
			hfmCodeMaxLen = len(v)
		}
		decodeMap[v] = k
	}
	// 每个前缀不一样
	ret := make([]byte, 0)
	// todo 性能差
	var f func(int, int)
	f = func(begin, end int) {
		if begin > end {
			return
		}
		if b, ok := decodeMap[string([]rune(str)[begin:end])]; ok {
			ret = append(ret, b)
			f(begin+(end-begin), begin+(end-begin) + hfmCodeMaxLen)
		} else {
			f(begin, end - 1)
		}
	}
	f(0, hfmCodeMaxLen)
	return ret
}

func subString(str string, begin, end int) string {
	if begin > end {
		panic(fmt.Sprintf("illegal index！begin: %d, end: %d", begin, end))
	}
	list := []rune(str)
	list = list[begin:end]
	return string(list)
}

func binaryString2Byte(binaryString string) byte {
	var ret byte
	var j byte = 1
	for i := 0; i < len(binaryString); i++ {
		char := binaryString[len(binaryString)-1-i]
		if char == '0' || char == '1' {
			if char == '1' {
				ret += j
			}
			j *= 2
		} else {
			panic("illegal binaryString ！")
		}
	}
	return ret
}

func byte2BinaryString(x byte) string {
	s := ""
	var len int
	for x > 0 {
		a := x % 2
		s = fmt.Sprintf("%d", a) + s
		x >>= 1
		len++
	}
	ret := ""
	for len < 8 {
		ret += "0"
		len++
	}
	return ret + s
}

func TestHuffman(t *testing.T) {
	resp, _ := http.Get("https://www.bilibili.com")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(len(body)/1024,"KB")
	src := body
	tree := CreateHuffmanTree(src)
	codeTable := HuffmanCode(tree)
	huffmanBytes, length := Encode(codeTable, src)
	fmt.Println(len(huffmanBytes)/1024,"KB")

	ret := Decode(huffmanBytes, codeTable, length)
	fmt.Println(string(ret))
}
