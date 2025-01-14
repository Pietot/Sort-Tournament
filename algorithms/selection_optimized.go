package sort

func SelectionSort(array []int) []int {
	arrLength := len(array)

	for i := 0; i < arrLength; i++ {
		minIndex := i

		for j := i + 1; j < arrLength; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		array[i], array[minIndex] = array[minIndex], array[i]
	}

	return array
}
