package recursion

import (
	"fmt"
	"testing"
)


// 字符串的全排列  - 回溯
func fullArrangement(str []byte, curPosition int) {
	if curPosition == len(str)-1 {
		fmt.Println(string(str))
		return
	}
	for i := curPosition; i < len(str); i++ {
		// cut branches
		if i != curPosition && str[curPosition] == str[i] {
			continue
		}
		str[curPosition], str[i] = str[i], str[curPosition]
		fullArrangement(str, curPosition+1)
		str[curPosition], str[i] = str[i], str[curPosition]
	}
}

func TestFullArrangement(t *testing.T) {
	fullArrangement([]byte("abc"), 0)
}

