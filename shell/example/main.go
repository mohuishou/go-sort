package main

import (
	"fmt"

	"../"
)

func main() {
	a := []int{34, 8, 64, 51, 32, 21, 5, 3, 6, 9, 8, 7, 4, 5, 2, 3, 4}
	fmt.Println("原始数组：", a)
	b, n := shell.SortHibbardInt(a)
	fmt.Println("排序数组：", b)
	fmt.Println("交换次数：", n)
}
