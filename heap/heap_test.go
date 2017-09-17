package HeapSort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{39, 100, 1, 99, 0, 34, 2, 35}
	HeapSort(arr)
	t.Log(arr)
}
