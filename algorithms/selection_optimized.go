package sort

func SelectionSortOptimized(array []int) []int {
	arrLength := len(array)
	sortedArray := make([]int, arrLength)

	for i := 0; i < arrLength; i++ {
		minIndex := i

		for j := i + 1; j < arrLength; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		sortedArray[i] = array[minIndex]
	}

	return sortedArray
}
