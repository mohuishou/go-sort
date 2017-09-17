//Package shell 排序
package shell

//SortInt 传统shell排序,间隔为N/2
//相邻间隔可能不互质，可能会出现前置排序无用的情况
func SortInt(a []int) ([]int, int) {
	n := 0
	aLen := len(a)

	//定义间隔
	for i := aLen / 2; i > 0; i = i / 2 {

		//插入排序
		for j := i; j < aLen; j++ {
			tmp := a[j]
			k := j
			for ; k >= i && tmp < a[k-i]; k = k - i {
				a[k] = a[k-i]
				n++
			}
			a[k] = tmp
		}
	}
	return a, n
}

//SortHibbardInt Hibbard算法，间隔为2^k-1
func SortHibbardInt(a []int) ([]int, int) {
	n, i := 0, 0
	aLen := len(a)
	for i = 1; i <= aLen-1; i = i*2 + 1 {
	}
	//定义间隔
	for ; i > 0; i = (i - 1) / 2 {
		// println(i)
		//插入排序
		for j := i; j < aLen; j++ {
			tmp := a[j]
			k := j
			for ; k >= i && tmp < a[k-i]; k = k - i {
				a[k] = a[k-i]
				n++
			}
			a[k] = tmp
		}
	}
	return a, n
}
