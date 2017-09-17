package quik

import (
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	arr := randArr(3000)
	Sort(arr)
	// t.Log(arr)
}

func randArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n * 10)
	}
	return arr
}
