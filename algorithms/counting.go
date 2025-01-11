package sort

func CountingSort(array []int) []int {
	max := getMax(array)
	count := make([]int, max+1)
	output := make([]int, len(array))

	for _, val := range array {
		count[val]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	for i := len(array) - 1; i >= 0; i-- {
		output[count[array[i]]-1] = array[i]
		count[array[i]]--
	}

	return output
}
