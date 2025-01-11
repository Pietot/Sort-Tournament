package sort

func TimSort(array []int) []int {
	n := len(array)
	minrun := 32
	for i := 0; i < n; i += minrun {
		insertionSort(array, i, min(i+minrun, n))
	}
	size := minrun
	for size < n {
		for left := 0; left < n; left += 2 * size {
			mid := min(left+size, n)
			right := min(left+2*size, n)
			if mid < right {
				tim_merge(array, left, mid, right)
			}
		}
		size *= 2
	}
	return array
}

func insertionSort(array []int, left, right int) {
	for i := left + 1; i < right; i++ {
		for j := i; j > left && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}

func tim_merge(array []int, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid
	if n1 < 0 || n2 < 0 {
		return
	}
	leftArray := make([]int, n1)
	rightArray := make([]int, n2)
	for i := 0; i < n1; i++ {
		leftArray[i] = array[left+i]
	}
	for i := 0; i < n2; i++ {
		rightArray[i] = array[mid+i]
	}
	i, j := 0, 0
	for k := left; k < right; k++ {
		if i < n1 && (j >= n2 || leftArray[i] <= rightArray[j]) {
			array[k] = leftArray[i]
			i++
		} else {
			array[k] = rightArray[j]
			j++
		}
	}
}
