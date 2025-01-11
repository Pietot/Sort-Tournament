package sort

func ShakerSort(array []int) []int {
	for i := 0; i < len(array)/2; i++ {
		swap := false
		for j := i; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swap = true
			}
		}
		if !swap {
			break
		}
		swap = false
		for j := len(array) - i - 2; j > i; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return array
}
