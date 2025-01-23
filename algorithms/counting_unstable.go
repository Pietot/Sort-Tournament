package sort

func CountingSortUnstable(array []int) []int {
	max := getMax(array)
	count := make([]int, max+1)

	for _, val := range array {
		count[val]++
	}

	index := 0

	for i, c := range count {
		for c > 0 {
			array[index] = i
			index++
			c--
		}
	}
    
	return array
}
