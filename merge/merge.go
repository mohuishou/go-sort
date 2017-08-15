package merge

//INFINITY 一个比较大的值用作哨兵
const INFINITY = 0xffff

func merge(arr []int, start, end int) {
	if start < end {
    //从中间划分，分别左右两边排序
		mid := (end + start) / 2
		merge(arr, start, mid)
		merge(arr, mid + 1, end)
    
    //下面进行归并操作，将两个长度较小但是已经排序完成的数组合并成一个较长长度的排序数组
    
    //新建一个数组用于存放左边的值
		arr1 := make([]int, mid - start + 2)
		copy(arr1, arr[start:mid + 1])
		arr1[mid - start + 1] = INFINITY

    //新建一个数组用于存放右边的值
		arr2 := make([]int, end - mid + 1)
		copy(arr2, arr[mid + 1:end+1])
		arr2[end - mid] = INFINITY

    //比较大小
		j, k := 0, 0
		for i := start; i <= end; i++ {
			if arr1[j] <= arr2[k] {
				arr[i] = arr1[j]
				j++
			} else {
				arr[i] = arr2[k]
				k++
			}
		}

	}
}
