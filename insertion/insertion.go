// Package insertion 插入排序
package insertion

//SortInt 插入排序
func SortInt(a []int) ([]int, int) {
	n := 0
	aLen := len(a)

	//从i=1开始，第一个数不用排序，一个数相当于已经有序了
	for i := 1; i < aLen; i++ {
		//取出新数（类似摸牌）
		tmp := a[i]

		//将新数和之前的数从后往前（从大到小）一次比较，如果新数更小,就将以前的数往后移一位
		j := i
		for ; (j > 0) && (tmp < a[j-1]); j-- {
			n++
			a[j] = a[j-1] //后移一位
		}
		//新数插入
		a[j] = tmp

	}
	return a, n
}
