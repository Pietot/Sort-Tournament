package sort

func DoubleSelectionSort(array []int) []int {
	for i := 0; i < len(array)/2; i++ {
		minIndex, maxIndex := i, len(array)-i-1
		if array[minIndex] > array[maxIndex] {
			minIndex, maxIndex = maxIndex, minIndex
		}
		for j := i + 1; j < len(array)-i; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			} else if array[j] > array[maxIndex] {
				maxIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
		if maxIndex == i {
			maxIndex = minIndex
		}
		array[len(array)-i-1], array[maxIndex] = array[maxIndex], array[len(array)-i-1]
	}
	return array
}
