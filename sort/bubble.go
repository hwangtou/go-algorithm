package sort

func bubble(arr []int) {
	for i := 0; i != (len(arr) - 1); i++ {
		for j := 0; j != (len(arr) - i - 1); j++ {
			if arr[j] > arr[j + 1] {
				swap := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = swap
			}
		}
	}
}