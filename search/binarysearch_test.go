package search

import "testing"

func TestBinarySearch(t *testing.T) {
	a := []interface{} { 2, 4, 5, 7, 15, 17 }
	comp := func(x interface{}) int {
		i := x.(int)
		if x == 15 {
			return 0
		} else if i < 15 {
			return -1
		} else {
			return 1
		}
	}
	index := BinarySearch(a, comp)
	if a[index] != 15 {
		t.Error("BinarySearch returned wrong index!")
	}
}
