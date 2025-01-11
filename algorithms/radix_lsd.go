package sort

const base = 64

func RadixLSDSort(array []int) []int {
	max := getMax(array)
	for exp := 1; max/exp > 0; exp *= base {
		countingSort(array, exp)
	}
	return array
}

func getMax(array []int) int {
	max := array[0]
	for _, val := range array {
		if val > max {
			max = val
		}
	}
	return max
}

func countingSort(array []int, exp int) {
	n := len(array)
	output := make([]int, n)
	count := make([]int, base)

	for i := 0; i < n; i++ {
		count[(array[i]/exp)%base]++
	}

	for i := 1; i < base; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[(array[i]/exp)%base]-1] = array[i]
		count[(array[i]/exp)%base]--
	}

	for i := 0; i < n; i++ {
		array[i] = output[i]
	}
}
