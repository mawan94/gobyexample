package sort

func IsEquals(arr1, arr2 []int) bool {
	if (arr1 != nil && arr2 == nil) || (arr1 == nil && arr2 != nil) {
		return false
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	if arr1 == nil && arr2 == nil {
		return true
	}
	return true
}

//func rightSort(arr []int) {
//	sort.Ints(arr)
//}

func Copy(source []int) []int {
	result := make([]int, len(source))
	copy(result, source)
	return result
}

func Swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func If(condition bool, arg1, arg2 interface{}) interface{} {
	if condition {
		return arg1
	}else {
		return arg2
	}
}
