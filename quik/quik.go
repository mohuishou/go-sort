package quik

//Sort 快速排序
func Sort(arr []int) {
	Quik(arr, 0, len(arr)-1)
}

//Quik 快速排序递归实现
func Quik(arr []int, left int, right int) {
	if right-left < 2 {
		return
	}
	p := middle3(arr, left, right)

	i := left + 1
	j := right - 2
	for {
		for arr[i] < p {
			i++
		}
		for arr[j] > p {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			break
		}
	}
	arr[i], arr[right-1] = arr[right-1], arr[i]
	Quik(arr, left, i-1)
	Quik(arr, i+1, right)
}

func middle3(arr []int, left int, right int) int {
	center := (right + left) / 2
	if arr[left] > arr[center] {
		arr[left], arr[center] = arr[center], arr[left]
	}
	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}
	if arr[center] > arr[right] {
		arr[center], arr[right] = arr[right], arr[center]
	}
	arr[center], arr[right-1] = arr[right-1], arr[center]
	return arr[right-1]
}
