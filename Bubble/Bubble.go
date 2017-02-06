package Bubble //冒泡排序

// SortInt int数组排序
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
