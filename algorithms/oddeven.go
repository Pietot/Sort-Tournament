package sort

func OddEvenSort(array []int) []int {
	n := len(array)
	sorted := false
	for !sorted {
		sorted = true
		for i := 1; i < n-1; i += 2 {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				sorted = false
			}
		}
		for i := 0; i < n-1; i += 2 {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				sorted = false
			}
		}
	}
	return array
}
