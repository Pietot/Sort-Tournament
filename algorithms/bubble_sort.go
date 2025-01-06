package gosort

func BubbleSortOptimized(array []int) []int {
	unsorted_len := len(array)
	for i := 0; i < len(array); i++ {
		swap := false
		for j := 0; j < unsorted_len-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swap = true
			}
		}
		if !swap {
			break
		}
		unsorted_len--
	}
	return array
}
