package sort

func IntroSort(array []int) []int {
	n := len(array)
	depthLimit := 2 * intLog2(n)
	return introSort(array, 0, n-1, depthLimit)
}

func introSort(array []int, low, high, depthLimit int) []int {
	if high-low < 16 {
		return InsertionSort(array[low : high+1])
	}
	if depthLimit == 0 {
		return HeapSort(array[low : high+1])
	}
	pivot := medianOfThree(array[low], array[(low+high)/2], array[high])
	pivotIndex := partition(array, low, high, pivot)
	introSort(array, low, pivotIndex, depthLimit-1)
	introSort(array, pivotIndex+1, high, depthLimit-1)
	return array
}

func partition(array []int, low, high, pivot int) int {
	i, j := low, high
	for i < j {
		for array[i] < pivot {
			i++
		}
		for array[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		array[i], array[j] = array[j], array[i]
	}
	return j
}

func medianOfThree(a, b, c int) int {
	if a < b {
		if b < c {
			return b
		}
		if a < c {
			return c
		}
		return a
	}
	if a < c {
		return a
	}
	if b < c {
		return c
	}
	return b
}

func intLog2(n int) int {
	result := 0
	for n >>= 1; n > 0; n >>= 1 {
		result++
	}
	return result
}
