package sort

import (
	"fmt"
	"math"
	"testing"
)

func elementLength(i int) int {
	if i < 0 {
		i = i * -1
	}
	if i < 10 {
		return 1
	}
	return 1 + elementLength(i/10)
}

func minValue(arr []int) int {
	minVal := arr[0]
	for i := 0; i < len(arr); i++ {
		if minVal > arr[i] {
			minVal = arr[i]
		}
	}
	return minVal
}

func maxValue(arr []int) int {
	maxVal := arr[0]
	for i := 0; i < len(arr); i++ {
		if maxVal < arr[i] {
			maxVal = arr[i]
		}
	}
	return maxVal
}

func subArr(arr []int) (negativeNumberArr []int, positiveNumberArr []int) {
	negativeNumberArr = make([]int, 0)
	positiveNumberArr = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 0 {
			positiveNumberArr = append(positiveNumberArr, arr[i])
		} else {
			negativeNumberArr = append(negativeNumberArr, arr[i])
		}
	}
	return negativeNumberArr, positiveNumberArr
}

func initBuckets(buckets [][]int) {
	for i := 0; i < 10; i++ {
		buckets[i] = make([]int, 0)
	}
}

func cover(arr []int, buckets [][]int, indexManager []int) {
	for i, curIdx := 0, 0; i < 10; i++ {
		for j := 0; j < indexManager[i]; j++ {
			if buckets[i][j] < 0 {
				arr[len(arr)-1-curIdx] = buckets[i][j]
			} else {
				arr[curIdx] = buckets[i][j]
			}
			curIdx++
		}
	}
}

func sort(arr []int, maxLength int, buckets [][]int) {
	for k := 1; maxLength > 0; maxLength-- {
		initBuckets(buckets)
		var indexManager = make([]int, 10)
		for i := 0; i < len(arr); i++ {
			curDigit := int(math.Abs(float64(arr[i]/(k)))) % 10
			buckets[curDigit] = append(buckets[curDigit], arr[i])
			indexManager[curDigit]++
		}
		cover(arr, buckets, indexManager)
		k *= 10
	}
}

func mergeArr(arr1, arr2, src []int) {
	copy(src, append(arr1, arr2...))
}

func radixSort(arr []int) {
	if arr == nil {
		return
	}
	buckets := make([][]int, 10)
	negativeNumberArr, positiveNumberArr := subArr(arr)
	negativeNumberArrMaxLen, positiveNumberArrMaxLen := elementLength(minValue(negativeNumberArr)), elementLength(maxValue(positiveNumberArr))

	sort(negativeNumberArr, negativeNumberArrMaxLen, buckets)
	sort(positiveNumberArr, positiveNumberArrMaxLen, buckets)
	mergeArr(negativeNumberArr, positiveNumberArr, arr)
}

func TestRadixSort(t *testing.T) {
	arr := []int{4, 6, 7, 2, 11, -2, 6, -2, 4,0, -5, 3}
	radixSort(arr)
	fmt.Println(arr)
}
