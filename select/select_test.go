package selection

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{1, 100, 2, 99, 3, 88}
	Sort(arr)
	t.Log(arr)
}
