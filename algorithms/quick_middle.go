package sort

func QuickSortMiddlePivot(array []int) []int {
	if len(array) < 2 {
		return array
	}
	left, right := 0, len(array)-1
	pivot := len(array) / 2
	array[pivot], array[right] = array[right], array[pivot]
	for i := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}
	array[left], array[right] = array[right], array[left]
	QuickSortMiddlePivot(array[:left])
	QuickSortMiddlePivot(array[left+1:])
	return array
}
