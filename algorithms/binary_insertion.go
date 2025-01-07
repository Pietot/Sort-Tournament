package sort

func BinaryInsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		key := array[i]
		left, right := 0, i
		for left < right {
			mid := left + (right-left)/2
			if key < array[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		copy(array[left+1:i+1], array[left:i])
		array[left] = key
	}
	return array
}
