package sort

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	middle := len(array) / 2
	left := MergeSort(array[:middle])
	right := MergeSort(array[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}
