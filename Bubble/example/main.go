package main

import (
	"fmt"

	"../"
)

func main() {
	//冒泡排序
	a := []int{8, 6, 1, 4, 3, 9, 5, 65, 66, 0}
	fmt.Println("原始数组：", a)
	b := bubble.SortInt(a)
	fmt.Println("排序数组：", b)
}
