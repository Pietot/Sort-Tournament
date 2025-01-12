package sort

func SmoothSort(array []int) []int {
	heapify(array)
	for i := len(array) - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		siftDown(array, 0, i)
	}
	return array
}

func heapify(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		siftDown(array, i, len(array))
	}
}

func siftDown(array []int, root, end int) {
	for {
		child := 2*root + 1
		if child >= end {
			break
		}
		if child+1 < end && array[child] < array[child+1] {
			child++
		}
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			root = child
		} else {
			break
		}
	}
}