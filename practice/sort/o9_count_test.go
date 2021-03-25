package sort

import (
	"fmt"
	"math"
	"testing"
)

func countSort(arr []int) {
	min, max := minValue(arr), maxValue(arr)
	model := make([]uint, max-min+1)

	for i := 0; i < len(arr); i++ {
		model[int(math.Abs(float64(min)))+arr[i]]++
	}

	for j, index := 0, 0; j < len(model); j++ {
		for ; model[j] > 0; model[j] -- {
			arr[index] = min + j
			index++
		}
	}
}

func TestCountSort(t *testing.T) {
	arr := []int{-4, -2, -3, -20, 4}
	countSort(arr)
	fmt.Println(arr)
}
