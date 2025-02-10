package main

import (
	"fmt"
	mySort "gosort/algorithms"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Benchmark struct {
	name                    string
	function                func([]int) []int
	timeSortedArray         float64
	timeRandomArray         float64
	timeRoughlySortedArray  float64
	timeNearlySortedArray   float64
	timeReversedSortedArray float64
	averageTime             float64
}

func main() {
	benchmarks := []*Benchmark{
		{"sort.Ints", func(arr []int) []int {
			sort.Ints(arr)
			return arr
		}, 0, 0, 0, 0, 0, 0},
		{"InsertionSort", mySort.InsertionSort, 0, 0, 0, 0, 0, 0},
		{"BinaryInsertionSort", mySort.BinaryInsertionSort, 0, 0, 0, 0, 0, 0},
		{"BubbleSort optimized", mySort.BubbleSortOptimized, 0, 0, 0, 0, 0, 0},
		{"SelectionSort optimized", mySort.SelectionSortOptimized, 0, 0, 0, 0, 0, 0},
		{"DoubleSelectionSort", mySort.DoubleSelectionSort, 0, 0, 0, 0, 0, 0},
		{"ShakerSort", mySort.ShakerSort, 0, 0, 0, 0, 0, 0},
		{"QuickSort middle pivot", mySort.QuickSort, 0, 0, 0, 0, 0, 0},
		{"MergeSort", mySort.MergeSort, 0, 0, 0, 0, 0, 0},
		{"HeapSort", mySort.HeapSort, 0, 0, 0, 0, 0, 0},
		{"CombSort", mySort.CombSort, 0, 0, 0, 0, 0, 0},
		{"ShellSort", mySort.ShellSort, 0, 0, 0, 0, 0, 0},
		{"Radix LSD sort base 64", mySort.RadixLSDSort, 0, 0, 0, 0, 0, 0},
		{"CountingSort stable", mySort.CountingSortStable, 0, 0, 0, 0, 0, 0},
		{"OddEvenSort", mySort.OddEvenSort, 0, 0, 0, 0, 0, 0},
		{"CircleSort", mySort.CircleSort, 0, 0, 0, 0, 0, 0},
		{"TournamentSort", mySort.TournamentSort, 0, 0, 0, 0, 0, 0},
		{"TreeSort", mySort.TreeSort, 0, 0, 0, 0, 0, 0},
		{"GnomeSort", mySort.GnomeSort, 0, 0, 0, 0, 0, 0},
		{"IntroSort ", mySort.IntroSort, 0, 0, 0, 0, 0, 0},
		{"TimSort", mySort.TimSort, 0, 0, 0, 0, 0, 0},
		{"SmoothSort", mySort.SmoothSort, 0, 0, 0, 0, 0, 0},
		{"BlockSort", mySort.BlockSort, 0, 0, 0, 0, 0, 0},
		{"BucketSort", mySort.BucketSort, 0, 0, 0, 0, 0, 0},
		{"CountingSort unstable", mySort.CountingSortUnstable, 0, 0, 0, 0, 0, 0},
		{"RedBlackTreeSort", mySort.RedBlackTreeSort, 0, 0, 0, 0, 0, 0},
	}

	n := 1_000_000

	randomArray := generateRandomArray(n)
	roughlySortedArray := generateRoughlySortedArray(n)
	nearlySortedArray := generateNearlySortedArray(n)
	sortedArray := generateSortedArray(n)
	reversedArray := generateReversedSortedArray(n)

	computeTimeStart := time.Now()

	for _, benchmark := range benchmarks {
		benchmark.timeRandomArray = measureTime(benchmark.function, append([]int(nil), randomArray...))
		benchmark.timeRoughlySortedArray = measureTime(benchmark.function, append([]int(nil), roughlySortedArray...))
		benchmark.timeNearlySortedArray = measureTime(benchmark.function, append([]int(nil), nearlySortedArray...))
		benchmark.timeSortedArray = measureTime(benchmark.function, append([]int(nil), sortedArray...))
		benchmark.timeReversedSortedArray = measureTime(benchmark.function, append([]int(nil), reversedArray...))
		benchmark.averageTime = (benchmark.timeRandomArray + benchmark.timeRoughlySortedArray + benchmark.timeNearlySortedArray + benchmark.timeSortedArray + benchmark.timeReversedSortedArray) / 5
		// Recompute the times 60 times to get a more accurate average time
		if benchmark.averageTime < 1000 {
			for i := 1; i < 60; i++ {
				benchmark.timeRandomArray += measureTime(benchmark.function, append([]int(nil), randomArray...))
				benchmark.timeRoughlySortedArray += measureTime(benchmark.function, append([]int(nil), roughlySortedArray...))
				benchmark.timeNearlySortedArray += measureTime(benchmark.function, append([]int(nil), nearlySortedArray...))
				benchmark.timeSortedArray += measureTime(benchmark.function, append([]int(nil), sortedArray...))
				benchmark.timeReversedSortedArray += measureTime(benchmark.function, append([]int(nil), reversedArray...))
			}
			benchmark.timeRandomArray /= 60
			benchmark.timeRoughlySortedArray /= 60
			benchmark.timeNearlySortedArray /= 60
			benchmark.timeSortedArray /= 60
			benchmark.timeReversedSortedArray /= 60
			benchmark.averageTime = (benchmark.timeRandomArray + benchmark.timeRoughlySortedArray + benchmark.timeNearlySortedArray + benchmark.timeSortedArray + benchmark.timeReversedSortedArray) / 5
		}

		fmt.Printf("\033[32m%s\033[0m\n", benchmark.name)
		fmt.Printf("Average time: %.0f ms\n", benchmark.averageTime)
		fmt.Printf("Random array: %.0f ms\n", benchmark.timeRandomArray)
		fmt.Printf("Roughly sorted array: %.0f ms\n", benchmark.timeRoughlySortedArray)
		fmt.Printf("Nearly sorted array: %.0f ms\n", benchmark.timeNearlySortedArray)
		fmt.Printf("Sorted array: %.0f ms\n", benchmark.timeSortedArray)
		fmt.Printf("Reversed array: %.0f ms\n", benchmark.timeReversedSortedArray)
		fmt.Println()
	}

	fmt.Printf("\033[32m%s %.0f minutes\033[0m\n", "Total time:", time.Since(computeTimeStart).Minutes())

	err := exportToCsv(benchmarks)
	if err != nil {
		fmt.Printf("\033[31m%s\033[0m\n", err)
	}
}

func generateSortedArray(n int) []int {
	array := make([]int, n)
	for i := range array {
		array[i] = i + 1
	}

	return array
}

func generateRandomArray(n int) []int {
	array := make([]int, n)
	for i := range array {
		array[i] = i + 1
	}

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}

	return array
}

func generateRoughlySortedArray(n int) []int {
	array := make([]int, n)
	for i := range array {
		array[i] = i + 1
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i += 5 {
		end := i + 5
		if end > n {
			end = n
		}
		random.Shuffle(end-i, func(a, b int) {
			array[i+a], array[i+b] = array[i+b], array[i+a]
		})
	}

	return array
}

func generateNearlySortedArray(n int) []int {
	array := make([]int, n)
	for i := range array {
		array[i] = i + 1
	}

	// Exchange 10% of the elements
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n/10; i++ {
		j := r.Intn(n)
		k := r.Intn(n)
		array[j], array[k] = array[k], array[j]
	}

	return array
}

func generateReversedSortedArray(n int) []int {
	array := make([]int, n)
	for i := range array {
		array[i] = n - i
	}

	return array
}

func measureTime(sortFunction func([]int) []int, array []int) float64 {
	start := time.Now()
	sortFunction(array)
	elapsed := time.Since(start).Seconds() * 1000
	return elapsed
}

func exportToCsv(benchmarks []*Benchmark) error {
	file, err := os.Create("benchmark.csv")
	if err != nil {
		return fmt.Errorf("failed creating file: %s", err)
	}
	defer file.Close()

	_, err = file.WriteString("Algorithm,Average time,Sorted Array,Random array,Roughly sorted array,Nearly sorted array,Reversed array\n")
	if err != nil {
		return fmt.Errorf("failed writing to file: %s", err)
	}

	for _, benchmark := range benchmarks {
		_, err = file.WriteString(fmt.Sprintf("%s,%.0f, %.0f,%.0f,%.0f,%.0f,%.0f\n", benchmark.name, math.Round(benchmark.averageTime), math.Round(benchmark.timeRandomArray), math.Round(benchmark.timeRoughlySortedArray), math.Round(benchmark.timeNearlySortedArray), math.Round(benchmark.timeSortedArray), math.Round(benchmark.timeReversedSortedArray)))
		if err != nil {
			return fmt.Errorf("failed writing to file: %s", err)
		}
	}

	fmt.Printf("\033[36m%s\033[0m\n", "Benchmark results successfully saved to benchmark.csv")
	return nil
}
