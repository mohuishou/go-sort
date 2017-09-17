package selection

//Sort 选择排序
func Sort(arr []int) {
	var LEN = len(arr)
	minIndex := 0
	for i := 0; i < LEN; i++ {
		minIndex = i
		for j := i; j < LEN; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}
}
