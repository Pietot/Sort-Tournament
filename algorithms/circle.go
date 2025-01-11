package sort

func CircleSort(array []int) []int {
	n := len(array)
	for i := 0; i < n/2; i++ {
		left, right := 0, n-1
		for left <= right {
			if array[left] > array[right] {
				array[left], array[right] = array[right], array[left]
			}
			left++
			right--
		}
	}
	return array
}
