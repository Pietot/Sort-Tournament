package sort

func QuickSortDualPivot(array []int) []int {
	if len(array) < 2 {
		return array
	}
	quickSortDualPivot(array, 0, len(array)-1)
	return array
}

func quickSortDualPivot(array []int, low, high int) {
	for low < high {
		lp, rp := qpartition(array, low, high)

		if (lp - low) < (high - rp) {
			quickSortDualPivot(array, low, lp-1)
			low = rp + 1
		} else {
			quickSortDualPivot(array, rp+1, high)
			high = lp - 1
		}
	}
}

func qpartition(array []int, low, high int) (int, int) {
	if array[low] > array[high] {
		array[low], array[high] = array[high], array[low]
	}
	pivot1 := array[low]
	pivot2 := array[high]

	i := low + 1
	lt := low + 1
	gt := high - 1

	for i <= gt {
		if array[i] < pivot1 {
			array[i], array[lt] = array[lt], array[i]
			lt++
		} else if array[i] > pivot2 {
			array[i], array[gt] = array[gt], array[i]
			gt--
		}
		i++
	}
	lt--
	gt++

	array[low], array[lt] = array[lt], array[low]
	array[high], array[gt] = array[gt], array[high]

	return lt, gt
}
