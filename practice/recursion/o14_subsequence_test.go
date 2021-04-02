package recursion

import (
	"fmt"
	"testing"
)

// 打印字符串的所有子序列
func subsequence(str []byte, cur int, path string) {
	if cur == len(str) {
		fmt.Println(path)
		return
	}
	subsequence(str, cur+1, fmt.Sprintf("%s  %s", path, string(str[cur])))
	subsequence(str, cur+1, path)
}

func TestSubsequence(t *testing.T) {
	subsequence([]byte("abccc"),0,"")
}
