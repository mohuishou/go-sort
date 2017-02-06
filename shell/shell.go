//Package shell 排序
package shell

//SortInt 传统shell排序
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
