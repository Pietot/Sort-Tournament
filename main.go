package main

import (
	"fmt"
	mySort "gosort/algorithms"
	"math/rand"
	"sort"
	"time"
)

func main() {
	array := generateUniqueRandomList(1_000_000)

	start := time.Now()
	sort.Ints(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32msort.Ints\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.InsertionSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32minsertion sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.BinaryInsertionSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mbinary insertion sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.BubbleSortOptimized(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mbubble sort optimized\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.SelectionSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mselection sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.DoubleSelectionSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mdouble selection sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.ShakerSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mshaker sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.QuickSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mquick sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.MergeSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mmerge sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.HeapSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mheap sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.CombSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mcomb sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.ShellSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mshell sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.RadixLSDSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mradix LSD sort base 64\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.CountingSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mcounting sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.OddEvenSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32modd-even sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.CircleSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mcircle sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.TournamentSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mtournament sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.TreeSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mtree sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.GnomeSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mgnome sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.IntroSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mintro sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.TimSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mTim sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.SmoothSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32msmooth sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)

	start = time.Now()
	mySort.BlockSort(append([]int(nil), array...))
	fmt.Printf("Time taken for \033[32mblock sort\033[0m: %.3f ms\n", time.Since(start).Seconds()*1000)
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
