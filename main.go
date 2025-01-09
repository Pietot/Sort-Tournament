package main

import (
	"fmt"
	mySort "gosort/algorithms"
	"math/rand"
	"time"
	"sort"
)

func main() {
	array := generateUniqueRandomList(30000)

	start := time.Now()
	sort.Ints(append([]int(nil), array...))
	fmt.Println("Time taken for sort.Ints:", time.Since(start))

	start = time.Now()
	mySort.BinaryInsertionSort(append([]int(nil), array...))
	fmt.Println("Time taken for BinaryInsertionSort:", time.Since(start))
	
	start = time.Now()
	mySort.BubbleSortOptimized(append([]int(nil), array...))
	fmt.Println("Time taken for BubbleSortOptimized:", time.Since(start))
	
	start = time.Now()
	mySort.DoubleSelectionSort(append([]int(nil), array...))
	fmt.Println("Time taken for DoubleSelectionSort:", time.Since(start))

	start = time.Now()
	mySort.ShakerSort(append([]int(nil), array...))
	fmt.Println("Time taken for ShakerSort:", time.Since(start))

	start = time.Now()
	mySort.QuickSort(append([]int(nil), array...))
	fmt.Println("Time taken for QuickSort:", time.Since(start))

	start = time.Now()
	mySort.MergeSort(append([]int(nil), array...))
	fmt.Println("Time taken for MergeSort:", time.Since(start))

	start = time.Now()
	mySort.HeapSort(append([]int(nil), array...))
	fmt.Println("Time taken for HeapSort:", time.Since(start))
	
}

func generateUniqueRandomList(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}