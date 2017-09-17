package main

import (
	"fmt"

	"github.com/mohuishou/go-sort/insertion"
)

func main() {
	a := []int{34, 8, 64, 51, 32, 21}
	fmt.Println("原始数组：", a)
	b, n := insertion.SortInt(a)
	fmt.Println("排序数组：", b)
	fmt.Println("交换次数：", n)
}
