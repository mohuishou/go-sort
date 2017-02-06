package main

import (
	"fmt"
)

func main() {
	//冒泡排序
	a := []int{8, 6, 1, 4, 3, 9, 5, 65, 66, 0}
	fmt.Println("原始数组：", a)
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
	fmt.Println("排序数组：", a)
}
