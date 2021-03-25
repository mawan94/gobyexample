package recursion

import (
	"fmt"
	"testing"
)

// 翻转字符串 (要求不用额外空间)
func reverseStr(str []byte, b, e int) {
	if str == nil || b == e {
		return
	}
	str[b], str[e] = str[e], str[b]
	reverseStr(str, b+1, e-1)
}

func TestReverseStr(t *testing.T) {
	str := []byte("hello")
	reverseStr(str, 0, len(str)-1)
	fmt.Println(string(str))
}
