package recursion

import (
	"fmt"
	"testing"
)

// 汉诺塔
//相传在古印度圣庙中，有一种被称为汉诺塔(Hanoi)的游戏。
//该游戏是在一块铜板装置上，有三根杆(编号A、B、C)，在A杆自下而上、由大到小按顺序放置n个金盘。
//游戏的目标：把A杆上的金盘全部移到C杆上，并仍保持原有顺序叠好。
//操作规则：每次只能移动一个盘子，并且在移动过程中三根杆上都始终保持大盘在下，小盘在上，操作过程中盘子可以置于A、B、C任一杆上

// a 为当前柱子
// b 为中转柱子
// c 为目标柱子
func towerOfHanoi(n int, a, b, c string) int {
	if n == 1 {
		fmt.Printf("%s -> %s \n", a, c)
		return 1
	}
	step1 := towerOfHanoi(n-1, a, c, b) // 将n-1块圆盘 从a移动到b 借助c
	fmt.Printf("%s -> %s \n", a, c)     // 将a最大的一块移动到c
	step2 := towerOfHanoi(n-1, b, a, c) // 将n-1块圆盘 从b移动到c 借助a
	return step1 + step2 + 1
}

func TestTowerOfHanoi(t *testing.T) {
	fmt.Println("一共移动了【", towerOfHanoi(3, "A", "B", "C"), "】次")
}
