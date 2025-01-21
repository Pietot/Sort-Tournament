package main

import (
	"fmt"
	mySort "gosort/algorithms"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Benchmark struct {
	name                   string
	function               func([]int) []int
	timeRandomList         float64
	timeRoughlySortedList  float64
	timeNearlySortedList   float64
	timeReversedSortedList float64
	averageTime            float64
}

func main() {
	benchmarks := []*Benchmark{
		{"sort.Ints", func(arr []int) []int {
			sort.Ints(arr)
			return arr
		}, 0, 0, 0, 0, 0},
		{"InsertionSort", mySort.InsertionSort, 0, 0, 0, 0, 0},
		{"BinaryInsertionSort", mySort.BinaryInsertionSort, 0, 0, 0, 0, 0},
		{"BubbleSort optimized", mySort.BubbleSortOptimized, 0, 0, 0, 0, 0},
		{"SelectionSort", mySort.SelectionSortOptimized, 0, 0, 0, 0, 0},
		{"DoubleSelectionSort", mySort.DoubleSelectionSort, 0, 0, 0, 0, 0},
		{"ShakerSort", mySort.ShakerSort, 0, 0, 0, 0, 0},
		{"QuickSort", mySort.QuickSort, 0, 0, 0, 0, 0},
		{"MergeSort", mySort.MergeSort, 0, 0, 0, 0, 0},
		{"HeapSort", mySort.HeapSort, 0, 0, 0, 0, 0},
		{"CombSort", mySort.CombSort, 0, 0, 0, 0, 0},
		{"ShellSort", mySort.ShellSort, 0, 0, 0, 0, 0},
		{"Radix LSD sort base 64", mySort.RadixLSDSort, 0, 0, 0, 0, 0},
		{"CountingSort", mySort.CountingSort, 0, 0, 0, 0, 0},
		{"OddEvenSort", mySort.OddEvenSort, 0, 0, 0, 0, 0},
		{"CircleSort", mySort.CircleSort, 0, 0, 0, 0, 0},
		{"TournamentSort", mySort.TournamentSort, 0, 0, 0, 0, 0},
		{"TreeSort", mySort.TreeSort, 0, 0, 0, 0, 0},
		{"GnomeSort", mySort.GnomeSort, 0, 0, 0, 0, 0},
		{"IntroSort ", mySort.IntroSort, 0, 0, 0, 0, 0},
		{"TimSort", mySort.TimSort, 0, 0, 0, 0, 0},
		{"SmoothSort", mySort.SmoothSort, 0, 0, 0, 0, 0},
		{"BlockSort", mySort.BlockSort, 0, 0, 0, 0, 0},
		{"BucketSort", mySort.BucketSort, 0, 0, 0, 0, 0},
	}

	n := 1_000_000

	randomArray := generateRandomList(n)
	roughlySortedArray := generateRoughlySortedList(n)
	nearlySortedArray := generateNearlySortedList(n)
	reversedArray := generateReversedSortedList(n)

	for _, benchmark := range benchmarks {
		benchmark.timeRandomList = measureTime(benchmark.function, append([]int(nil), randomArray...))
		benchmark.timeRoughlySortedList = measureTime(benchmark.function, append([]int(nil), roughlySortedArray...))
		benchmark.timeNearlySortedList = measureTime(benchmark.function, append([]int(nil), nearlySortedArray...))
		benchmark.timeReversedSortedList = measureTime(benchmark.function, append([]int(nil), reversedArray...))
		benchmark.averageTime = (benchmark.timeRandomList + benchmark.timeRoughlySortedList + benchmark.timeNearlySortedList + benchmark.timeReversedSortedList) / 4
		// Recompute the times 60 times to get a more accurate average time
		if benchmark.averageTime < 1000 {
			for i := 1; i < 60; i++ {
				benchmark.timeRandomList += measureTime(benchmark.function, append([]int(nil), randomArray...))
				benchmark.timeRoughlySortedList += measureTime(benchmark.function, append([]int(nil), roughlySortedArray...))
				benchmark.timeNearlySortedList += measureTime(benchmark.function, append([]int(nil), nearlySortedArray...))
				benchmark.timeReversedSortedList += measureTime(benchmark.function, append([]int(nil), reversedArray...))
			}
			benchmark.timeRandomList /= 60
			benchmark.timeRoughlySortedList /= 60
			benchmark.timeNearlySortedList /= 60
			benchmark.timeReversedSortedList /= 60
			benchmark.averageTime = (benchmark.timeRandomList + benchmark.timeRoughlySortedList + benchmark.timeNearlySortedList + benchmark.timeReversedSortedList) / 4
		}

		fmt.Printf("\033[32m%s\033[0m\n", benchmark.name)
		fmt.Printf("Random list: %.0f ms\n", benchmark.timeRandomList)
		fmt.Printf("Roughly sorted list: %.0f ms\n", benchmark.timeRoughlySortedList)
		fmt.Printf("Nearly sorted list: %.0f ms\n", benchmark.timeNearlySortedList)
		fmt.Printf("Reversed list: %.0f ms\n", benchmark.timeReversedSortedList)
		fmt.Printf("Average time: %.0f ms\n", benchmark.averageTime)
		fmt.Println()
	}

	err := exportToCsv(benchmarks)
	if err != nil {
		fmt.Printf("\033[31m%s\033[0m\n", err)
	}
}

func generateRandomList(n int) []int {
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

func generateRoughlySortedList(n int) []int {
	array := make([]int, n)
	for i := range array {
		array[i] = i + 1
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i += 5 {
		end := i + 5
		if end > n {
			end = n
		}
		r.Shuffle(end-i, func(a, b int) {
			array[i+a], array[i+b] = array[i+b], array[i+a]
		})
	}

	return array
}

func generateNearlySortedList(n int) []int {
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

func generateReversedSortedList(n int) []int {
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

	_, err = file.WriteString("Algorithm,Random list,Roughly sorted list,Nearly sorted list,Reversed list,Average time\n")
	if err != nil {
		return fmt.Errorf("failed writing to file: %s", err)
	}

	for _, benchmark := range benchmarks {
		_, err = file.WriteString(fmt.Sprintf("%s,%.3f,%.3f,%.3f,%.3f,%.3f\n", benchmark.name, benchmark.timeRandomList, benchmark.timeRoughlySortedList, benchmark.timeNearlySortedList, benchmark.timeReversedSortedList, benchmark.averageTime))
		if err != nil {
			return fmt.Errorf("failed writing to file: %s", err)
		}
	}

	fmt.Printf("\033[36m%s\033[0m\n", "Benchmark results successfully saved to benchmark.csv")
	return nil
}
