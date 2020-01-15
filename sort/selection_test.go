package sort

import "testing"

func TestSelection(t *testing.T) {
	randArray := []int{6, 8, 9, 0, 4, 3, 2, 7, 1, 5}
	selection(randArray)
	t.Log(randArray)
}
