package sort

import (
	mySort "gosort/algorithms"
	"testing"
)

func TestSorts(test *testing.T) {
	functions := []struct {
		name     string
		function func([]int) []int
	}{
		{"InsertionSort", mySort.InsertionSort},
		{"BinaryInsertionSort", mySort.BinaryInsertionSort},
		{"BubbleSortOptimized", mySort.BubbleSortOptimized},
		{"SelectionSortOptimized", mySort.SelectionSortOptimized},
		{"DoubleSelectionSort", mySort.DoubleSelectionSort},
		{"ShakerSort", mySort.ShakerSort},
		{"QuickSort", mySort.QuickSort},
		{"MergeSort", mySort.MergeSort},
		{"HeapSort", mySort.HeapSort},
		{"CombSort", mySort.CombSort},
		{"ShellSort", mySort.ShellSort},
		{"RadixLSDSort", mySort.RadixLSDSort},
		{"CountingSortStable", mySort.CountingSortStable},
		{"OddEvenSort", mySort.OddEvenSort},
		{"CircleSort", mySort.CircleSort},
		{"TournamentSort", mySort.TournamentSort},
		{"TreeSort", mySort.TreeSort},
		{"GnomeSort", mySort.GnomeSort},
		{"IntroSort", mySort.IntroSort},
		{"TimSort", mySort.TimSort},
		{"SmoothSort", mySort.SmoothSort},
		{"BlockSort", mySort.BlockSort},
		{"BucketSort", mySort.BucketSort},
		{"CountingSortUnstable", mySort.CountingSortUnstable},
	}

	unsortedArray := [8]int{5, 3, 8, 6, 2, 7, 1, 4}
	sortedArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, function := range functions {
		test.Run(function.name, func(test *testing.T) {
			result := function.function(unsortedArray[:])
			if !compareArrays(result, sortedArray) {
				test.Errorf("Expected %v but got %v", sortedArray, result)
			}
		})
	}
}

func compareArrays(first_array []int, second_array [8]int) bool {
	if len(first_array) != len(second_array) {
		return false
	}

	for index, value := range first_array {
		if value != second_array[index] {
			return false
		}
	}

	return true
}
