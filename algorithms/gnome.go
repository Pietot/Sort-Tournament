package sort

func GnomeSort(array []int) []int {
	n := len(array)
	for i := 1; i < n; {
		if i == 0 || array[i-1] <= array[i] {
			i++
		} else {
			array[i], array[i-1] = array[i-1], array[i]
			i--
		}
	}
	return array
}
