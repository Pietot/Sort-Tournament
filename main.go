package main

import (
	"fmt"
	mySort "gosort/algorithms"
	"math/rand"
	"sort"
	"time"
)

func main() {
	array := generateUniqueRandomList(30000)

	start := time.Now()
	sort.Ints(append([]int(nil), array...))
	fmt.Printf("Time taken for sort.Ints: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.InsertionSort(append([]int(nil), array...))
	fmt.Printf("Time taken for InsertionSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.BinaryInsertionSort(append([]int(nil), array...))
	fmt.Printf("Time taken for BinaryInsertionSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.BubbleSortOptimized(append([]int(nil), array...))
	fmt.Printf("Time taken for BubbleSort optimized: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.SelectionSort(append([]int(nil), array...))
	fmt.Printf("Time taken for SelectionSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.DoubleSelectionSort(append([]int(nil), array...))
	fmt.Printf("Time taken for DoubleSelectionSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.ShakerSort(append([]int(nil), array...))
	fmt.Printf("Time taken for ShakerSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.QuickSort(append([]int(nil), array...))
	fmt.Printf("Time taken for QuickSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.MergeSort(append([]int(nil), array...))
	fmt.Printf("Time taken for MergeSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.HeapSort(append([]int(nil), array...))
	fmt.Printf("Time taken for HeapSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.CombSort(append([]int(nil), array...))
	fmt.Printf("Time taken for CombSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.ShellSort(append([]int(nil), array...))
	fmt.Printf("Time taken for ShellSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.RadixLSDSort(append([]int(nil), array...))
	fmt.Printf("Time taken for Radix LSD sort base 64: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.CountingSort(append([]int(nil), array...))
	fmt.Printf("Time taken for CountingSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.OddEvenSort(append([]int(nil), array...))
	fmt.Printf("Time taken for OddEvenSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.CircleSort(append([]int(nil), array...))
	fmt.Printf("Time taken for CircleSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.TournamentSort(append([]int(nil), array...))
	fmt.Printf("Time taken for TournamentSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.TreeSort(append([]int(nil), array...))
	fmt.Printf("Time taken for TreeSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.GnomeSort(append([]int(nil), array...))
	fmt.Printf("Time taken for GnomeSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.IntroSort(append([]int(nil), array...))
	fmt.Printf("Time taken for IntroSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.TimSort(append([]int(nil), array...))
	fmt.Printf("Time taken for TimSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.SmoothSort(append([]int(nil), array...))
	fmt.Printf("Time taken for SmoothSort: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.BlockSort(append([]int(nil), array...))
	fmt.Printf("Time taken for BlockSort: %.3f ms\n", time.Since(start).Seconds()*1000)
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
