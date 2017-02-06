//Package bubble 冒泡排序
package bubble

// SortInt int数组排序
// a[]int
func SortInt(a []int) []int {
	aLen := len(a)
	for i := 0; i < aLen-1; i++ {
		for j := 0; j < aLen-i-1; j++ {
			if a[j] > a[j+1] {
				tmp := a[j+1]
				a[j+1] = a[j]
				a[j] = tmp
			}
		}
	}
	return a
}

//SortInt32 int32数组排序
func SortInt32(a []int) []int {
	aLen := len(a)
	for i := 0; i < aLen-1; i++ {
		for j := 0; j < aLen-i-1; j++ {
			if a[j] > a[j+1] {
				tmp := a[j+1]
				a[j+1] = a[j]
				a[j] = tmp
			}
		}
	}
	return a
}
