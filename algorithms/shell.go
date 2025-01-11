package sort

func ShellSort(array []int) []int {
	arrLength := len(array)
	gap := arrLength / 2
	shrink := 2.3

	for gap > 0 {
		for i := gap; i < arrLength; i++ {
			for j := i; j >= gap && array[j] < array[j-gap]; j -= gap {
				array[j], array[j-gap] = array[j-gap], array[j]
			}
		}
		gap = int(float64(gap) / shrink)
	}

	return array
}
