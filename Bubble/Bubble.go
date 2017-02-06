//Package bubble 冒泡排序
package bubble

// SortInt int数组排序
func SortInt(a []int) []int {

	aLen := len(a)

	for i := aLen - 1; i >= 0; i-- {
		//标记，如果flag一次冒泡之后没有改变，那么证明排序已完成，不需要再次排序，直接退出
		flag := 0

		//一次冒泡
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				//交换两个变量的值，无需引入临时变量
				a[j], a[j+1] = a[j+1], a[j]

				//有交换，flag=1
				flag = 1
			}
		}

		//判断一次冒泡，是否存在交换
		if flag == 0 {
			break
		}

	}

	return a
}
