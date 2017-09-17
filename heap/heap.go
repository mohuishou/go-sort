package HeapSort

import (
	"fmt"
)

//HeapSort 堆排序
func HeapSort(arr []int) {
	LEN := len(arr)
	for i := LEN/2 - 1; i >= 0; i-- {
		HeapAjust(arr, i, LEN)
	}
	fmt.Println(arr)

	for i := LEN - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		fmt.Println(arr, i)
		HeapAjust(arr, 0, i)
	}
}

//HeapAjust 堆调整
func HeapAjust(arr []int, start int, length int) {
	tmp := arr[start]
	for i := 2*start + 1; i < length; i = i * 2 {
		if i+1 < length && arr[i] < arr[i+1] {
			i++
		}
		if tmp > arr[i] {
			break
		}
		arr[start] = arr[i]
		start = i
	}
	arr[start] = tmp
}
