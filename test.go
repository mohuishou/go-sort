package main

import (
	"fmt"
)

func main() {
	fmt.Println(insert([]int{7, 99, 0, 45, 2, 45, 67, 12}))
}

func insert(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for ; j > 0 && tmp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}
