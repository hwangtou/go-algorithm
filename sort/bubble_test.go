package sort

import "testing"

func TestBubble(t *testing.T) {
	randArray := []int{6, 8, 9, 0, 4, 3, 2, 7, 1, 5}
	bubble(randArray)
	t.Log(randArray)
}
