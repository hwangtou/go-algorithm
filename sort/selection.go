package sort

import "log"

func selection(arr []int) {
	for i := 0; i != (len(arr) - 1); i++ {
		min := i
		for j := i; j != len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			swap := arr[min]
			arr[min] = arr[i]
			arr[i] = swap
		}
		log.Println(arr, min, i)
	}
}
