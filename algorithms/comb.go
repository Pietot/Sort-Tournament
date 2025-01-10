package sort

func CombSort(array []int) []int {
	gap := len(array)
	arrLength := len(array)
	shrink := 1.3
	sorted := false

	for !sorted {
		gap = int(float64(gap) / shrink)
		if gap <= 1 {
			gap = 1
			sorted = true
		}

		for offset := 0; offset+gap < arrLength; offset++ {
			if array[offset] > array[offset+gap] {
				array[offset], array[offset+gap] = array[offset+gap], array[offset]
				sorted = false
			}
		}
	}

	return array
}
